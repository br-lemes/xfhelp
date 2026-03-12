package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var panelListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all XFCE panels with their IDs and output names",
	RunE: func(cmd *cobra.Command, args []string) error {
		panels, err := getPanels(realFunc)
		if err != nil {
			return err
		}
		for id, panel := range panels {
			fmt.Printf("Panel %d: %s\n", id, panel["output-name"])
		}
		return nil
	},
}

func init() {
	panelCmd.AddCommand(panelListCmd)
}
