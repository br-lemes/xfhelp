package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var outputsCmd = &cobra.Command{
	Use:   "outputs",
	Short: "List available XFCE display outputs (monitors, screens)",
	RunE: func(cmd *cobra.Command, args []string) error {
		profile, err := getActiveProfile(realFunc)
		if err != nil {
			return err
		}

		outputs, err := getOutputs(realFunc, profile)
		if err != nil {
			return err
		}
		fmt.Println(strings.Join(outputs, "\n"))
		return nil
	},
}

func getOutputs(queryFunc queryFunc, profile string) ([]string, error) {
	out, err := queryFunc("-c", "displays", "-l", "-v", "-p", "/"+profile)
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`^/` + profile + `/([^/]+)/Active\s+true$`)
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	outputs := []string{"Automatic", "Primary"}
	for _, line := range lines {
		matches := re.FindStringSubmatch(strings.TrimSpace(line))
		if matches == nil {
			continue
		}
		outputs = append(outputs, matches[1])
	}
	return outputs, nil
}

func getActiveProfile(queryFunc queryFunc) (string, error) {
	out, err := queryFunc("-c", "displays", "-p", "/ActiveProfile")
	if err != nil {
		return "", err
	}
	activeProfile := strings.TrimSpace(string(out))
	if activeProfile == "" {
		activeProfile = "Default"
	}
	return activeProfile, nil
}

func init() {
	rootCmd.AddCommand(outputsCmd)
}
