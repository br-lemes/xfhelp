package cmd

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/spf13/cobra"
)

var panelShowCmd = &cobra.Command{
	Use:   "show [panel-id] [output-name]",
	Short: "Show XFCE panel on specified output",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		panelID, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid panel-id: %s", args[0])
		}

		profile, err := getActiveProfile(realFunc)
		if err != nil {
			return err
		}
		outputs, err := getOutputs(realFunc, profile)
		if err != nil {
			return err
		}
		outputName := args[1]
		if !slices.Contains(outputs, outputName) {
			return fmt.Errorf("output-name %q is not valid", outputName)
		}

		err = setPanelOutput(realFunc, panelID, outputName)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	panelCmd.AddCommand(panelShowCmd)
}
