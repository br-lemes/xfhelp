package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var panelHideCmd = &cobra.Command{
	Use:   "hide [panel-id]",
	Short: "Hide XFCE panel by setting fictitious output name",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		panelID, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid panel-id: %s", args[0])
		}

		err = setPanelOutput(realFunc, panelID, "Hidden")
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	panelCmd.AddCommand(panelHideCmd)
}
