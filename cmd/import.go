package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var dryRunCommands []string

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import XFCE settings from a JSON file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		dryRun, _ := cmd.Flags().GetBool("dry-run")
		data, err := readImportFile(args[0])
		if err != nil {
			return err
		}
		err = validateImport(data, exportSchema)
		if err != nil {
			return err
		}

		for channel, props := range data {
			err := processChannel(channel, props, dryRun)
			if err != nil {
				return err
			}
		}

		if dryRun && len(dryRunCommands) > 0 {
			fmt.Printf("\n%s\n", strings.Join(dryRunCommands, "\n"))
		}
		return nil
	},
}

func readImportFile(path string) (ConfigMap, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var result ConfigMap
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func validateImport(data ConfigMap, schema SchemaSpec) error {
	for channel, props := range data {
		channelSpec, inSchema := schema[channel]
		if !inSchema {
			return fmt.Errorf("channel %q is not in schema", channel)
		}
		for prop, value := range props {
			_, inSchema := channelSpec[prop]
			if !inSchema {
				return fmt.Errorf(
					"property %q in channel %q is not in schema", prop, channel,
				)
			}
			err := validateImportValue(prop, value, channelSpec[prop].Type)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func validateImportValue(prop string, value any, t PropertyType) error {
	switch t {
	case TypeInt:
		v, ok := value.(float64)
		if !ok {
			return propertyTypeError(prop, jsonTypeName(value), t)
		}
		if v != float64(int64(v)) {
			return propertyTypeError(prop, fmt.Sprintf("float (%v)", v), t)
		}
	case TypeFloat:
		_, ok := value.(float64)
		if !ok {
			return propertyTypeError(prop, jsonTypeName(value), t)
		}
	case TypeBool:
		_, ok := value.(bool)
		if !ok {
			return propertyTypeError(prop, jsonTypeName(value), t)
		}
	case TypeString:
		_, ok := value.(string)
		if !ok {
			return propertyTypeError(prop, jsonTypeName(value), t)
		}
	case TypeArrayInt, TypeArrayFloat, TypeArrayBool, TypeArrayString:
		slice, ok := value.([]any)
		if !ok {
			return propertyTypeError(prop, jsonTypeName(value), t)
		}
		var elementType PropertyType
		switch t {
		case TypeArrayInt:
			elementType = TypeInt
		case TypeArrayFloat:
			elementType = TypeFloat
		case TypeArrayBool:
			elementType = TypeBool
		case TypeArrayString:
			elementType = TypeString
		}
		for _, elem := range slice {
			err := validateImportValue(prop, elem, elementType)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func jsonTypeName(value any) string {
	switch v := value.(type) {
	case float64:
		if v == float64(int64(v)) {
			return fmt.Sprintf("int (%v)", int64(v))
		}
		return fmt.Sprintf("float (%v)", v)
	case string:
		return fmt.Sprintf("string (%q)", v)
	case []any:
		return fmt.Sprintf("array (%v)", value)
	default:
		return fmt.Sprintf("%T (%v)", value, value)
	}
}

func propertyTypeError(prop string, value any, t PropertyType) error {
	return fmt.Errorf("property %q: expected %s, got %s", prop, t, value)
}

func processChannel(channel string, props map[string]any, dryRun bool) error {
	current, err := getProperties(realFunc, channel)
	if err != nil {
		return err
	}

	showChannel := true
	for prop, value := range props {
		spec := exportSchema[channel][prop]
		err := processProperty(
			channel,
			prop,
			value,
			current[prop],
			spec.Type,
			dryRun,
			&showChannel,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func processProperty(
	channel,
	prop string,
	value,
	currentValue any,
	propType PropertyType,
	dryRun bool,
	showChannel *bool,
) error {
	next := anyToXfconf(value)
	if currentValue == next {
		return nil
	}

	if dryRun {
		if *showChannel {
			fmt.Printf("%s:\n", channel)
			*showChannel = false
		}
		fmt.Printf("  %q: %q -> %q\n", prop, currentValue, next)
		return applyProperty(dryRunQuery, channel, prop, value, propType)
	} else {
		return applyProperty(realFunc, channel, prop, value, propType)
	}
}

func applyProperty(
	queryFunc queryFunc,
	channel, prop string,
	value any,
	t PropertyType,
) error {
	args := []string{"-c", channel, "-n", "-p", prop}
	xfType := xfconfType(t)
	switch v := value.(type) {
	case []any:
		for _, elem := range v {
			args = append(args, "-s", anyToXfconf(elem))
		}
		for range v {
			args = append(args, "-t", xfType)
		}
	default:
		args = append(args, "-s", anyToXfconf(value), "-t", xfType)
	}
	_, err := queryFunc(args...)
	return err
}

func anyToXfconf(value any) string {
	switch v := value.(type) {
	case bool:
		if v {
			return "true"
		}
		return "false"
	case float64:
		s := fmt.Sprintf("%g", v)
		return s
	case []any:
		parts := make([]string, len(v))
		for i, elem := range v {
			parts[i] = anyToXfconf(elem)
		}
		return "[" + strings.Join(parts, ",") + "]"
	default:
		return fmt.Sprintf("%v", v)
	}
}

func xfconfType(t PropertyType) string {
	switch t {
	case TypeInt, TypeArrayInt:
		return "int"
	case TypeFloat, TypeArrayFloat:
		return "double"
	case TypeBool, TypeArrayBool:
		return "bool"
	default:
		return "string"
	}
}

func dryRunQuery(args ...string) ([]byte, error) {
	parts := make([]string, len(args)+1)
	parts[0] = "xfconf-query"
	for i, arg := range args {
		if arg != "-" && strings.HasPrefix(arg, "-") {
			parts[i+1] = arg
		} else {
			parts[i+1] = fmt.Sprintf("%q", arg)
		}
	}
	dryRunCommands = append(dryRunCommands, strings.Join(parts, " "))
	return nil, nil
}

func init() {
	rootCmd.AddCommand(importCmd)
	importCmd.Flags().
		Bool("dry-run", false, "Show pending changes without applying them")
}
