package cmd

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/cobra"
)

type (
	Panel struct {
		AutoHideBehavior *int       `json:"autohide-behavior"`
		BackgroundRGBA   *[]float64 `json:"background-rgba"`
		BackgroundStyle  *int       `json:"background-style"`
		EnableStruts     *bool      `json:"enable-struts"`
		IconSize         *int       `json:"icon-size"`
		Length           *int       `json:"length"`
		LengthAdjust     *bool      `json:"length-adjust"`
		Mode             *int       `json:"mode"`
		Plugins          []Plugin   `json:"plugins"`
		Position         *string    `json:"position"`
		PositionLocked   *bool      `json:"position-locked"`
		Size             *int       `json:"size"`
		SpanMonitors     *bool      `json:"span-monitors"`
	}

	Plugin struct {
		Type   string
		Config any
	}

	NoConfig struct{}

	WhiskermenuConfig struct {
		ButtonIcon                   *string   `json:"button-icon"`
		ButtonSingleRow              *bool     `json:"button-single-row"`
		ButtonTitle                  *string   `json:"button-title"`
		CategoryIconSize             *int      `json:"category-icon-size"`
		CategoryShowName             *bool     `json:"category-show-name"`
		CommandHibernate             *string   `json:"command-hibernate"`
		CommandLockscreen            *string   `json:"command-lockscreen"`
		CommandLogout                *string   `json:"command-logout"`
		CommandLogoutuser            *string   `json:"command-logoutuser"`
		CommandMenueditor            *string   `json:"command-menueditor"`
		CommandProfile               *string   `json:"command-profile"`
		CommandRestart               *string   `json:"command-restart"`
		CommandSettings              *string   `json:"command-settings"`
		CommandShutdown              *string   `json:"command-shutdown"`
		CommandSuspend               *string   `json:"command-suspend"`
		CommandSwitchuser            *string   `json:"command-switchuser"`
		ConfirmSessionCommand        *bool     `json:"confirm-session-command"`
		DefaultCategory              *int      `json:"default-category"`
		Favorites                    *[]string `json:"favorites"`
		FavoritesInRecent            *bool     `json:"favorites-in-recent"`
		HoverSwitchCategory          *bool     `json:"hover-switch-category"`
		LauncherIconSize             *int      `json:"launcher-icon-size"`
		LauncherShowDescription      *bool     `json:"launcher-show-description"`
		LauncherShowName             *bool     `json:"launcher-show-name"`
		LauncherShowTooltip          *bool     `json:"launcher-show-tooltip"`
		MenuHeight                   *int      `json:"menu-height"`
		MenuOpacity                  *int      `json:"menu-opacity"`
		MenuWidth                    *int      `json:"menu-width"`
		PositionCategoriesAlternate  *bool     `json:"position-categories-alternate"`
		PositionCategoriesHorizontal *bool     `json:"position-categories-horizontal"`
		PositionCommandsAlternate    *bool     `json:"position-commands-alternate"`
		PositionProfileAlternate     *bool     `json:"position-profile-alternate"`
		PositionSearchAlternate      *bool     `json:"position-search-alternate"`
		ProfileShape                 *int      `json:"profile-shape"`
		RecentItemsMax               *int      `json:"recent-items-max"`
		ShowButtonIcon               *bool     `json:"show-button-icon"`
		ShowButtonTitle              *bool     `json:"show-button-title"`
		ShowCommandHibernate         *bool     `json:"show-command-hibernate"`
		ShowCommandLockscreen        *bool     `json:"show-command-lockscreen"`
		ShowCommandLogout            *bool     `json:"show-command-logout"`
		ShowCommandLogoutuser        *bool     `json:"show-command-logoutuser"`
		ShowCommandMenueditor        *bool     `json:"show-command-menueditor"`
		ShowCommandProfile           *bool     `json:"show-command-profile"`
		ShowCommandRestart           *bool     `json:"show-command-restart"`
		ShowCommandSettings          *bool     `json:"show-command-settings"`
		ShowCommandShutdown          *bool     `json:"show-command-shutdown"`
		ShowCommandSuspend           *bool     `json:"show-command-suspend"`
		ShowCommandSwitchuser        *bool     `json:"show-command-switchuser"`
		SortCategories               *bool     `json:"sort-categories"`
		StayOnFocusOut               *bool     `json:"stay-on-focus-out"`
		ViewMode                     *int      `json:"view-mode"`
	}

	LauncherConfig struct {
		Items *[]string `json:"items"`
	}

	TasklistConfig struct {
		FlatButtons                 *bool `json:"flat-buttons"`
		Grouping                    *bool `json:"grouping"`
		IncludeAllMonitors          *bool `json:"include-all-monitors"`
		IncludeAllWorkspaces        *bool `json:"include-all-workspaces"`
		MiddleClick                 *int  `json:"middle-click"`
		ShowHandle                  *bool `json:"show-handle"`
		ShowLabels                  *bool `json:"show-labels"`
		ShowOnlyMinimized           *bool `json:"show-only-minimized"`
		ShowTooltips                *bool `json:"show-tooltips"`
		ShowWireframes              *bool `json:"show-wireframes"`
		SortOrder                   *int  `json:"sort-order"`
		SwitchWorkspaceOnUnminimize *bool `json:"switch-workspace-on-unminimize"`
		WindowScrolling             *bool `json:"window-scrolling"`
	}

	SeparatorConfig struct {
		Expand *bool `json:"expand"`
		Style  *int  `json:"style"`
	}

	PagerConfig struct {
		MiniatureView      *bool `json:"miniature-view"`
		Numbering          *bool `json:"numbering"`
		Rows               *int  `json:"rows"`
		WorkspaceScrolling *bool `json:"workspace-scrolling"`
		WrapWorkspaces     *bool `json:"wrap-workspaces"`
	}

	SystrayConfig struct {
		HideNewItems  *bool `json:"hide-new-items"`
		IconSize      *int  `json:"icon-size"`
		MenuIsPrimary *bool `json:"menu-is-primary"`
		ShowFrame     *bool `json:"show-frame"`
		SingleRow     *bool `json:"single-row"`
		SquareIcons   *bool `json:"square-icons"`
		SymbolicIcons *bool `json:"symbolic-icons"`
	}

	PulseaudioConfig struct {
		EnableKeyboardShortcuts *bool   `json:"enable-keyboard-shortcuts"`
		EnableMpris             *bool   `json:"enable-mpris"`
		EnableMultimediaKeys    *bool   `json:"enable-multimedia-keys"`
		EnableWnck              *bool   `json:"enable-wnck"`
		MixerCommand            *string `json:"mixer-command"`
		MultimediaKeysToAll     *bool   `json:"multimedia-keys-to-all"`
		PlaySound               *bool   `json:"play-sound"`
		RecIndicatorPersistent  *bool   `json:"rec-indicator-persistent"`
		ShowNotifications       *int    `json:"show-notifications"`
		VolumeMax               *int    `json:"volume-max"`
		VolumeStep              *int    `json:"volume-step"`
	}

	SystemloadConfig struct {
		CPU                  *ResourceConfig `json:"cpu"`
		Memory               *ResourceConfig `json:"memory"`
		Network              *ResourceConfig `json:"network"`
		Swap                 *ResourceConfig `json:"swap"`
		Uptime               *UptimeConfig   `json:"uptime"`
		SystemMonitorCommand *string         `json:"system-monitor-command"`
		Timeout              *int            `json:"timeout"`
		TimeoutSeconds       *int            `json:"timeout-seconds"`
	}

	ResourceConfig struct {
		Color   *[]float64 `json:"color"`
		Enabled *bool      `json:"enabled"`
		Label   *string    `json:"label"`
	}

	UptimeConfig struct {
		Enabled *bool `json:"enabled"`
	}

	ClockConfig struct {
		Command           *string `json:"command"`
		DigitalFormat     *string `json:"digital-format"`
		DigitalLayout     *int    `json:"digital-layout"`
		DigitalTimeFont   *string `json:"digital-time-font"`
		DigitalTimeFormat *string `json:"digital-time-format"`
		FlashSeparators   *bool   `json:"flash-separators"`
		Fuzziness         *int    `json:"fuzziness"`
		Mode              *int    `json:"mode"`
		ShowGrid          *bool   `json:"show-grid"`
		ShowInactive      *bool   `json:"show-inactive"`
		ShowMeridiem      *bool   `json:"show-meridiem"`
		ShowMilitary      *bool   `json:"show-military"`
		ShowSeconds       *bool   `json:"show-seconds"`
		ShowWeekNumbers   *bool   `json:"show-week-numbers"`
		Timezone          *string `json:"timezone"`
		TooltipFormat     *string `json:"tooltip-format"`
	}

	SchemaNode map[string]any
)

var (
	configFor = map[string]func() any{
		"whiskermenu":          func() any { return &WhiskermenuConfig{} },
		"launcher":             func() any { return &LauncherConfig{} },
		"tasklist":             func() any { return &TasklistConfig{} },
		"separator":            func() any { return &SeparatorConfig{} },
		"pager":                func() any { return &PagerConfig{} },
		"systray":              func() any { return &SystrayConfig{} },
		"pulseaudio":           func() any { return &PulseaudioConfig{} },
		"systemload":           func() any { return &SystemloadConfig{} },
		"clock":                func() any { return &ClockConfig{} },
		"power-manager-plugin": func() any { return &NoConfig{} },
		"notification-plugin":  func() any { return &NoConfig{} },
		"xfce4-clipman-plugin": func() any { return &NoConfig{} },
	}

	panelSchemaCmd = &cobra.Command{
		Use:   "schema",
		Short: "Print the JSON Schema for XFCE panel and plugin configurations",
		RunE: func(cmd *cobra.Command, args []string) error {
			schema, err := generateSchema()
			if err != nil {
				return err
			}
			fmt.Println(string(schema))
			return nil
		},
	}
)

func generateSchema() ([]byte, error) {
	variants := []SchemaNode{}
	for typeName, factory := range configFor {
		cfg := factory()
		_, isNoConfig := cfg.(*NoConfig)

		variant := SchemaNode{
			"type":                 "object",
			"additionalProperties": false,
			"required":             []string{"type"},
			"properties": SchemaNode{
				"type": SchemaNode{"const": typeName},
			},
		}
		if !isNoConfig {
			variant["properties"].(SchemaNode)["config"] =
				reflectSchema(cfg, nil)
		}
		variants = append(variants, variant)
	}

	overrides := map[string]SchemaNode{
		"plugins": {
			"type":  "array",
			"items": SchemaNode{"oneOf": variants},
		},
	}

	schema := SchemaNode{
		"additionalProperties": false,

		"properties": reflectSchema(&Panel{}, overrides)["properties"],
		"type":       "object",
	}

	return json.MarshalIndent(schema, "", "    ")
}

func reflectSchema(v any, overrides map[string]SchemaNode) map[string]any {
	t := reflect.TypeOf(v).Elem()
	if t.NumField() == 0 {
		return SchemaNode{
			"type":                 "object",
			"additionalProperties": false,
		}
	}

	properties := SchemaNode{}

	for i := range t.NumField() {
		f := t.Field(i)
		name := strings.Split(f.Tag.Get("json"), ",")[0]

		if name == "" || name == "-" {
			continue
		}

		override, ok := overrides[name]
		if ok {
			properties[name] = override
			continue
		}

		t := f.Type
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		var fieldSchema SchemaNode
		switch t.Kind() {
		case reflect.String:
			fieldSchema = SchemaNode{"type": "string"}
		case reflect.Bool:
			fieldSchema = SchemaNode{"type": "boolean"}
		case reflect.Float32, reflect.Float64:
			fieldSchema = SchemaNode{"type": "number"}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
			reflect.Int64:
			fieldSchema = SchemaNode{"type": "integer"}
		case reflect.Slice:
			var item SchemaNode
			elem := t.Elem()
			if elem.Kind() == reflect.Ptr {
				elem = elem.Elem()
			}
			switch elem.Kind() {
			case reflect.String:
				item = SchemaNode{"type": "string"}
			case reflect.Bool:
				item = SchemaNode{"type": "boolean"}
			case reflect.Float32, reflect.Float64:
				item = SchemaNode{"type": "number"}
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
				reflect.Int64:
				item = SchemaNode{"type": "integer"}
			}
			fieldSchema = SchemaNode{"type": "array", "items": item}
		case reflect.Struct:
			fieldSchema = reflectSchema(reflect.New(t).Interface(), nil)
		}

		properties[name] = fieldSchema
	}

	return SchemaNode{
		"type":                 "object",
		"additionalProperties": false,
		"properties":           properties,
	}
}

func init() {
	panelCmd.AddCommand(panelSchemaCmd)
}
