package cmd

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type PanelSpec map[int]map[string]any

var panelCmd = &cobra.Command{
	Use:   "panel",
	Short: "Manage XFCE panels",
}

var panelSchema map[string]PropertySpec = map[string]PropertySpec{
	"autohide-behavior": {TypeInt},
	"background-rgba":   {TypeArrayFloat},
	"background-style":  {TypeInt},
	"enable-struts":     {TypeBool},
	"icon-size":         {TypeInt},
	"length":            {TypeFloat},
	"length-adjust":     {TypeBool},
	"mode":              {TypeInt},
	"output-name":       {TypeString},
	"plugin-ids":        {TypeArrayInt},
	"position":          {TypeString},
	"position-locked":   {TypeBool},
	"size":              {TypeInt},
}

func getPanels(queryFunc queryFunc) (PanelSpec, error) {
	out, err := queryFunc("-c", "xfce4-panel", "-l", "-v")
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`/panels/panel-(\d+)/([^\s]+)\s+(.+)`)
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	props := make(PanelSpec)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		matches := re.FindStringSubmatch(line)
		if len(matches) != 4 {
			continue
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			continue
		}
		if props[id] == nil {
			props[id] = make(map[string]any)
		}
		key := matches[2]
		value := matches[3]
		props[id][key], err = convertValue(key, value, panelSchema[key].Type)
		if err != nil {
			return nil, err
		}
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
