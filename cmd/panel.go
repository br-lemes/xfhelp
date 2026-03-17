package cmd

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var panelCmd = &cobra.Command{
	Use:   "panel",
	Short: "Manage XFCE panels",
}

func getPanels(queryFunc queryFunc) (map[int]string, error) {
	out, err := queryFunc("-c", "xfce4-panel", "-l", "-v")
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`/panels/panel-(\d+)/output-name\s+(.+)`)
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	props := make(map[int]string)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		matches := re.FindStringSubmatch(line)
		if len(matches) != 3 {
			continue
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			continue
		}
		props[id] = matches[2]
	}
	return props, nil
}

func setPanelOutput(queryFunc queryFunc, panelID int, outputName string) error {
	panels, err := getPanels(queryFunc)
	if err != nil {
		return err
	}
	if _, exists := panels[panelID]; !exists {
		return fmt.Errorf("panel %d does not exist", panelID)
	}
	panelPath := fmt.Sprintf("/panels/panel-%d/output-name", panelID)
	_, err = queryFunc("-c", "xfce4-panel", "-p", panelPath, "-s", outputName)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(panelCmd)
}
