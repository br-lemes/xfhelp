package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export XFCE settings to JSON",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		untracked, _ := cmd.Flags().GetBool("untracked")
		var config ConfigMap
		var err error
		if untracked {
			config, err = getUntracked(realFunc, exportSchema)
		} else {
			config, err = getTracked(realFunc, exportSchema)
		}
		if err != nil {
			return err
		}
		jsonData, err := json.MarshalIndent(config, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(jsonData))
		return nil
	},
}

func getTracked(queryFunc queryFunc, schema SchemaSpec) (ConfigMap, error) {
	result := make(ConfigMap)
	for channel, channelSpec := range schema {
		props, err := getProperties(queryFunc, channel)
		if err != nil {
			return nil, err
		}
		filtered := make(map[string]any)
		for prop, value := range props {
			propSpec, propInSchema := channelSpec[prop]
			if propInSchema {
				v, err := convertValue(prop, value, propSpec.Type)
				if err != nil {
					return nil, err
				}
				filtered[prop] = v
			}
		}
		if len(filtered) > 0 {
			result[channel] = filtered
		}
	}
	return result, nil
}

func getUntracked(queryFunc queryFunc, schema SchemaSpec) (ConfigMap, error) {
	result := make(ConfigMap)
	channels, err := getChannels(queryFunc)
	if err != nil {
		return nil, err
	}
	for _, channel := range channels {
		props, err := getProperties(queryFunc, channel)
		if err != nil {
			return nil, err
		}
		channelSpec, inSchema := schema[channel]
		filtered := make(map[string]any)
		for prop, value := range props {
			_, propInSchema := channelSpec[prop]
			if !inSchema || !propInSchema {
				filtered[prop] = value
			}
		}
		if len(filtered) > 0 {
			result[channel] = filtered
		}
	}
	return result, nil
}

func getChannels(queryFunc queryFunc) ([]string, error) {
	out, err := queryFunc("-l")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	if len(lines) > 0 && strings.Contains(lines[0], ":") {
		return lines[1:], nil
	}
	return lines, nil
}

func getProperties(
	queryFunc queryFunc,
	channel string,
) (map[string]string, error) {
	out, err := queryFunc("-c", channel, "-l", "-v")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	props := make(map[string]string)
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			key := parts[0]
			val := strings.Join(parts[1:], " ")
			props[key] = val
		}
	}
	return props, nil
}

func convertValue(prop string, value string, t PropertyType) (any, error) {
	isArray := strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]")
	isArrayType := t == TypeArrayInt || t == TypeArrayBool ||
		t == TypeArrayFloat || t == TypeArrayString
	if isArray && !isArrayType {
		return nil, fmt.Errorf(
			"property %q: expected scalar, got array %q", prop, value,
		)
	}
	if !isArray && isArrayType {
		return nil, fmt.Errorf(
			"property %q: expected array, got scalar %q", prop, value,
		)
	}
	if isArray {
		elements := strings.Split(value[1:len(value)-1], ",")
		var elementType PropertyType
		switch t {
		case TypeArrayInt:
			elementType = TypeInt
		case TypeArrayBool:
			elementType = TypeBool
		case TypeArrayFloat:
			elementType = TypeFloat
		case TypeArrayString:
			elementType = TypeString
		}
		result := make([]any, len(elements))
		for i, element := range elements {
			v, err :=
				convertValue(prop, strings.TrimSpace(element), elementType)
			if err != nil {
				return nil, err
			}
			result[i] = v
		}
		return result, nil
	}
	switch t {
	case TypeInt:
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil,
				fmt.Errorf("property %q: expected int, got %q", prop, value)
		}
		return v, nil
	case TypeBool:
		v, err := strconv.ParseBool(value)
		if err != nil {
			return nil,
				fmt.Errorf("property %q: expected bool, got %q", prop, value)
		}
		return v, nil
	case TypeFloat:
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil,
				fmt.Errorf("property %q: expected float, got %q", prop, value)
		}
		return v, nil
	}
	return value, nil
}

func init() {
	rootCmd.AddCommand(exportCmd)
	exportCmd.Flags().Bool(
		"untracked",
		false,
		"Export properties not covered by the schema instead",
	)
}
