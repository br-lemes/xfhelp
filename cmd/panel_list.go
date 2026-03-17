package cmd

import (
	"fmt"
	"maps"
	"slices"

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

		ids := slices.Collect(maps.Keys(panels))
		slices.Sort(ids)
		for _, id := range ids {
			fmt.Printf("Panel %d: %s\n", id, panels[id])
		}
		return nil
	},
}

func init() {
	panelCmd.AddCommand(panelListCmd)
}
