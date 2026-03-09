package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var exportSchema = SchemaSpec{
	"keyboard-layout": {
		"/Default/XkbDisable": {Type: TypeBool},
	},
	"keyboards": {
		"/Default/KeyRepeat":       {Type: TypeBool},
		"/Default/KeyRepeat/Delay": {Type: TypeInt},
		"/Default/KeyRepeat/Rate":  {Type: TypeInt},
		"/Default/Numlock":         {Type: TypeBool},
		"/Default/RestoreNumlock":  {Type: TypeBool},
	},
	"xfce4-keyboard-shortcuts": {
		"/commands/custom/<Alt>F2/startup-notify":  {Type: TypeBool},
		"/commands/custom/<Alt>F3/startup-notify":  {Type: TypeBool},
		"/commands/custom/<Primary><Alt>Delete":    {Type: TypeString},
		"/commands/custom/<Primary><Alt>Escape":    {Type: TypeString},
		"/commands/custom/<Primary><Alt>t":         {Type: TypeString},
		"/commands/custom/<Super>f":                {Type: TypeString},
		"/commands/custom/<Super>l":                {Type: TypeString},
		"/commands/custom/<Super>m":                {Type: TypeString},
		"/commands/custom/<Super>p":                {Type: TypeString},
		"/commands/custom/<Super>t":                {Type: TypeString},
		"/commands/custom/<Super>v":                {Type: TypeString},
		"/commands/custom/<Super>w":                {Type: TypeString},
		"/commands/custom/Print":                   {Type: TypeString},
		"/commands/custom/override":                {Type: TypeBool},
		"/providers":                               {Type: TypeArrayString},
		"/xfwm4/custom/<Alt><Shift>Tab":            {Type: TypeString},
		"/xfwm4/custom/<Alt>Delete":                {Type: TypeString},
		"/xfwm4/custom/<Alt>F10":                   {Type: TypeString},
		"/xfwm4/custom/<Alt>F11":                   {Type: TypeString},
		"/xfwm4/custom/<Alt>F12":                   {Type: TypeString},
		"/xfwm4/custom/<Alt>F4":                    {Type: TypeString},
		"/xfwm4/custom/<Alt>F6":                    {Type: TypeString},
		"/xfwm4/custom/<Alt>F7":                    {Type: TypeString},
		"/xfwm4/custom/<Alt>F8":                    {Type: TypeString},
		"/xfwm4/custom/<Alt>F9":                    {Type: TypeString},
		"/xfwm4/custom/<Alt>Insert":                {Type: TypeString},
		"/xfwm4/custom/<Alt>Tab":                   {Type: TypeString},
		"/xfwm4/custom/<Alt>space":                 {Type: TypeString},
		"/xfwm4/custom/<Primary><Alt>End":          {Type: TypeString},
		"/xfwm4/custom/<Primary><Alt>Home":         {Type: TypeString},
		"/xfwm4/custom/<Primary><Alt>d":            {Type: TypeString},
		"/xfwm4/custom/<Primary><Shift><Alt>Left":  {Type: TypeString},
		"/xfwm4/custom/<Primary><Shift><Alt>Right": {Type: TypeString},
		"/xfwm4/custom/<Primary><Shift><Alt>Up":    {Type: TypeString},
		"/xfwm4/custom/<Primary>F1":                {Type: TypeString},
		"/xfwm4/custom/<Primary>F10":               {Type: TypeString},
		"/xfwm4/custom/<Primary>F11":               {Type: TypeString},
		"/xfwm4/custom/<Primary>F12":               {Type: TypeString},
		"/xfwm4/custom/<Primary>F2":                {Type: TypeString},
		"/xfwm4/custom/<Primary>F3":                {Type: TypeString},
		"/xfwm4/custom/<Primary>F4":                {Type: TypeString},
		"/xfwm4/custom/<Primary>F5":                {Type: TypeString},
		"/xfwm4/custom/<Primary>F6":                {Type: TypeString},
		"/xfwm4/custom/<Primary>F7":                {Type: TypeString},
		"/xfwm4/custom/<Primary>F8":                {Type: TypeString},
		"/xfwm4/custom/<Primary>F9":                {Type: TypeString},
		"/xfwm4/custom/<Shift><Alt>Page_Down":      {Type: TypeString},
		"/xfwm4/custom/<Shift><Alt>Page_Up":        {Type: TypeString},
		"/xfwm4/custom/<Super>Tab":                 {Type: TypeString},
		"/xfwm4/custom/Down":                       {Type: TypeString},
		"/xfwm4/custom/Escape":                     {Type: TypeString},
		"/xfwm4/custom/Left":                       {Type: TypeString},
		"/xfwm4/custom/Right":                      {Type: TypeString},
		"/xfwm4/custom/Up":                         {Type: TypeString},
		"/xfwm4/custom/override":                   {Type: TypeBool},
	},
	"xfce4-notifyd": {
		"/date-time-custom-format":  {Type: TypeString},
		"/log-level":                {Type: TypeString},
		"/log-level-apps":           {Type: TypeString},
		"/log-max-size-enabled":     {Type: TypeBool},
		"/notification-log":         {Type: TypeBool},
		"/notify-location":          {Type: TypeString},
		"/plugin/after-menu-shown":  {Type: TypeString},
		"/plugin/hide-clear-prompt": {Type: TypeBool},
		"/plugin/hide-on-read":      {Type: TypeBool},
		"/show-notifications-on":    {Type: TypeString},
	},
	"xfce4-power-manager": {
		"/xfce4-power-manager/battery-button-action":       {Type: TypeInt},
		"/xfce4-power-manager/brightness-level-on-ac":      {Type: TypeInt},
		"/xfce4-power-manager/brightness-level-on-battery": {Type: TypeInt},
		"/xfce4-power-manager/brightness-on-ac":            {Type: TypeInt},
		"/xfce4-power-manager/brightness-on-battery":       {Type: TypeInt},
		"/xfce4-power-manager/brightness-step-count":       {Type: TypeInt},
		"/xfce4-power-manager/brightness-switch":           {Type: TypeInt},
		"/xfce4-power-manager/critical-power-action":       {Type: TypeInt},
		"/xfce4-power-manager/critical-power-level":        {Type: TypeInt},
		"/xfce4-power-manager/dpms-enabled":                {Type: TypeBool},
		"/xfce4-power-manager/dpms-on-ac-off":              {Type: TypeInt},
		"/xfce4-power-manager/dpms-on-ac-sleep":            {Type: TypeInt},
		"/xfce4-power-manager/dpms-on-battery-off":         {Type: TypeInt},
		"/xfce4-power-manager/dpms-on-battery-sleep":       {Type: TypeInt},
		"/xfce4-power-manager/general-notification":        {Type: TypeBool},
		"/xfce4-power-manager/handle-brightness-keys":      {Type: TypeBool},
		"/xfce4-power-manager/hibernate-button-action":     {Type: TypeInt},
		"/xfce4-power-manager/inactivity-on-ac":            {Type: TypeInt},
		"/xfce4-power-manager/inactivity-on-battery":       {Type: TypeInt},
		"/xfce4-power-manager/inactivity-sleep-mode-on-ac": {Type: TypeInt},
		"/xfce4-power-manager/inactivity-sleep-mode-on-battery"://
		{Type: TypeInt},
		"/xfce4-power-manager/lid-action-on-ac":              {Type: TypeInt},
		"/xfce4-power-manager/lid-action-on-battery":         {Type: TypeInt},
		"/xfce4-power-manager/lock-screen-suspend-hibernate": {Type: TypeBool},
		"/xfce4-power-manager/logind-handle-lid-switch":      {Type: TypeBool},
		"/xfce4-power-manager/power-button-action":           {Type: TypeInt},
		"/xfce4-power-manager/show-panel-label":              {Type: TypeInt},
		"/xfce4-power-manager/show-presentation-indicator":   {Type: TypeBool},
		"/xfce4-power-manager/show-tray-icon":                {Type: TypeBool},
		"/xfce4-power-manager/sleep-button-action":           {Type: TypeInt},
	},
	"xfce4-screensaver": {
		"/lock/embedded-keyboard/displayed": {Type: TypeBool},
		"/lock/embedded-keyboard/enabled":   {Type: TypeBool},
		"/lock/enabled":                     {Type: TypeBool},
		"/lock/logout/enabled":              {Type: TypeBool},
		"/lock/saver-activation/delay":      {Type: TypeInt},
		"/lock/saver-activation/enabled":    {Type: TypeBool},
		"/lock/sleep-activation":            {Type: TypeBool},
		"/lock/status-messages/enabled":     {Type: TypeBool},
		"/lock/user-switching/enabled":      {Type: TypeBool},
		"/saver/enabled":                    {Type: TypeBool},
		"/saver/fullscreen-inhibit":         {Type: TypeBool},
		"/saver/idle-activation/enabled":    {Type: TypeBool},
		"/saver/mode":                       {Type: TypeInt},
	},
	"xfce4-session": {
		"/compat/LaunchGNOME":  {Type: TypeBool},
		"/general/SaveOnExit":  {Type: TypeBool},
		"/shutdown/LockScreen": {Type: TypeBool},
	},
	"xfwm4": {
		"/general/activate_action":                {Type: TypeString},
		"/general/borderless_maximize":            {Type: TypeBool},
		"/general/box_move":                       {Type: TypeBool},
		"/general/box_resize":                     {Type: TypeBool},
		"/general/button_layout":                  {Type: TypeString},
		"/general/button_offset":                  {Type: TypeInt},
		"/general/button_spacing":                 {Type: TypeInt},
		"/general/click_to_focus":                 {Type: TypeBool},
		"/general/cycle_apps_only":                {Type: TypeBool},
		"/general/cycle_draw_frame":               {Type: TypeBool},
		"/general/cycle_hidden":                   {Type: TypeBool},
		"/general/cycle_minimized":                {Type: TypeBool},
		"/general/cycle_minimum":                  {Type: TypeBool},
		"/general/cycle_preview":                  {Type: TypeBool},
		"/general/cycle_raise":                    {Type: TypeBool},
		"/general/cycle_tabwin_mode":              {Type: TypeInt},
		"/general/cycle_workspaces":               {Type: TypeBool},
		"/general/double_click_action":            {Type: TypeString},
		"/general/double_click_distance":          {Type: TypeInt},
		"/general/double_click_time":              {Type: TypeInt},
		"/general/easy_click":                     {Type: TypeString},
		"/general/focus_delay":                    {Type: TypeInt},
		"/general/focus_hint":                     {Type: TypeBool},
		"/general/focus_new":                      {Type: TypeBool},
		"/general/frame_border_top":               {Type: TypeInt},
		"/general/frame_opacity":                  {Type: TypeInt},
		"/general/full_width_title":               {Type: TypeBool},
		"/general/horiz_scroll_opacity":           {Type: TypeBool},
		"/general/inactive_opacity":               {Type: TypeInt},
		"/general/maximized_offset":               {Type: TypeInt},
		"/general/mousewheel_rollup":              {Type: TypeBool},
		"/general/move_opacity":                   {Type: TypeInt},
		"/general/placement_mode":                 {Type: TypeString},
		"/general/placement_ratio":                {Type: TypeInt},
		"/general/popup_opacity":                  {Type: TypeInt},
		"/general/prevent_focus_stealing":         {Type: TypeBool},
		"/general/raise_delay":                    {Type: TypeInt},
		"/general/raise_on_click":                 {Type: TypeBool},
		"/general/raise_on_focus":                 {Type: TypeBool},
		"/general/raise_with_any_button":          {Type: TypeBool},
		"/general/repeat_urgent_blink":            {Type: TypeBool},
		"/general/resize_opacity":                 {Type: TypeInt},
		"/general/scroll_workspaces":              {Type: TypeBool},
		"/general/shadow_delta_height":            {Type: TypeInt},
		"/general/shadow_delta_width":             {Type: TypeInt},
		"/general/shadow_delta_x":                 {Type: TypeInt},
		"/general/shadow_delta_y":                 {Type: TypeInt},
		"/general/shadow_opacity":                 {Type: TypeInt},
		"/general/show_app_icon":                  {Type: TypeBool},
		"/general/show_dock_shadow":               {Type: TypeBool},
		"/general/show_frame_shadow":              {Type: TypeBool},
		"/general/show_popup_shadow":              {Type: TypeBool},
		"/general/snap_resist":                    {Type: TypeBool},
		"/general/snap_to_border":                 {Type: TypeBool},
		"/general/snap_to_windows":                {Type: TypeBool},
		"/general/snap_width":                     {Type: TypeInt},
		"/general/theme":                          {Type: TypeString},
		"/general/tile_on_move":                   {Type: TypeBool},
		"/general/title_alignment":                {Type: TypeString},
		"/general/title_font":                     {Type: TypeString},
		"/general/title_horizontal_offset":        {Type: TypeInt},
		"/general/title_shadow_active":            {Type: TypeBool},
		"/general/title_shadow_inactive":          {Type: TypeBool},
		"/general/title_vertical_offset_active":   {Type: TypeInt},
		"/general/title_vertical_offset_inactive": {Type: TypeInt},
		"/general/titleless_maximize":             {Type: TypeBool},
		"/general/toggle_workspaces":              {Type: TypeBool},
		"/general/unredirect_overlays":            {Type: TypeBool},
		"/general/urgent_blink":                   {Type: TypeBool},
		"/general/use_compositing":                {Type: TypeBool},
		"/general/workspace_count":                {Type: TypeInt},
		"/general/workspace_names":                {Type: TypeArrayString},
		"/general/wrap_cycle":                     {Type: TypeBool},
		"/general/wrap_layout":                    {Type: TypeBool},
		"/general/wrap_resistance":                {Type: TypeInt},
		"/general/wrap_windows":                   {Type: TypeBool},
		"/general/wrap_workspaces":                {Type: TypeBool},
		"/general/zoom_desktop":                   {Type: TypeBool},
		"/general/zoom_pointer":                   {Type: TypeBool},
	},
	"xsettings": {
		"/Gdk/WindowScalingFactor":       {Type: TypeInt},
		"/Gtk/ButtonImages":              {Type: TypeBool},
		"/Gtk/CanChangeAccels":           {Type: TypeBool},
		"/Gtk/ColorPalette":              {Type: TypeString},
		"/Gtk/CursorThemeName":           {Type: TypeString},
		"/Gtk/CursorThemeSize":           {Type: TypeInt},
		"/Gtk/DecorationLayout":          {Type: TypeString},
		"/Gtk/DialogsUseHeader":          {Type: TypeBool},
		"/Gtk/FontName":                  {Type: TypeString},
		"/Gtk/IconSizes":                 {Type: TypeString},
		"/Gtk/KeyThemeName":              {Type: TypeString},
		"/Gtk/MenuBarAccel":              {Type: TypeString},
		"/Gtk/MenuImages":                {Type: TypeBool},
		"/Gtk/MonospaceFontName":         {Type: TypeString},
		"/Gtk/TitlebarMiddleClick":       {Type: TypeString},
		"/Gtk/ToolbarStyle":              {Type: TypeString},
		"/Net/CursorBlink":               {Type: TypeBool},
		"/Net/CursorBlinkTime":           {Type: TypeInt},
		"/Net/DndDragThreshold":          {Type: TypeInt},
		"/Net/DoubleClickDistance":       {Type: TypeInt},
		"/Net/DoubleClickTime":           {Type: TypeInt},
		"/Net/EnableEventSounds":         {Type: TypeBool},
		"/Net/EnableInputFeedbackSounds": {Type: TypeBool},
		"/Net/IconThemeName":             {Type: TypeString},
		"/Net/SoundThemeName":            {Type: TypeString},
		"/Net/ThemeName":                 {Type: TypeString},
		"/Xfce/LastCustomDPI":            {Type: TypeInt},
		"/Xft/Antialias":                 {Type: TypeInt},
		"/Xft/DPI":                       {Type: TypeInt},
		"/Xft/HintStyle":                 {Type: TypeString},
		"/Xft/Hinting":                   {Type: TypeInt},
		"/Xft/RGBA":                      {Type: TypeString},
	},
}

type (
	PropertyType string
	PropertySpec struct{ Type PropertyType }
	ChannelSpec  map[string]PropertySpec
	SchemaSpec   map[string]ChannelSpec
	ConfigMap    map[string]map[string]any
)

const (
	TypeInt         PropertyType = "int"
	TypeBool        PropertyType = "bool"
	TypeFloat       PropertyType = "float"
	TypeString      PropertyType = "string"
	TypeArrayInt    PropertyType = "array:int"
	TypeArrayBool   PropertyType = "array:bool"
	TypeArrayFloat  PropertyType = "array:float"
	TypeArrayString PropertyType = "array:string"
)

func (t PropertyType) IsValid() bool {
	switch t {
	case TypeInt, TypeBool, TypeFloat, TypeString,
		TypeArrayInt, TypeArrayBool, TypeArrayFloat, TypeArrayString:
		return true
	default:
		return false
	}
}

var schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Print the JSON Schema of supported properties",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		root := map[string]any{
			"additionalProperties": false,
			"properties":           buildChannels(),
			"type":                 "object",
		}
		jsonData, err := json.MarshalIndent(root, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(jsonData))
	},
}

func buildChannels() map[string]any {
	channels := make(map[string]any)
	for channel, channelSpec := range exportSchema {
		channels[channel] = map[string]any{
			"additionalProperties": false,
			"properties":           buildProperties(channelSpec),
			"type":                 "object",
		}
	}
	return channels
}

func buildProperties(channelSpec ChannelSpec) map[string]any {
	properties := make(map[string]any)
	for prop, spec := range channelSpec {
		properties[prop] = buildType(spec.Type)
	}
	return properties
}

func buildType(t PropertyType) map[string]any {
	switch t {
	case TypeInt:
		return map[string]any{"type": "integer"}
	case TypeFloat:
		return map[string]any{"type": "number"}
	case TypeBool:
		return map[string]any{"type": "boolean"}
	case TypeString:
		return map[string]any{"type": "string"}
	case TypeArrayInt:
		return map[string]any{
			"type":  "array",
			"items": map[string]any{"type": "integer"},
		}
	case TypeArrayFloat:
		return map[string]any{
			"type":  "array",
			"items": map[string]any{"type": "number"},
		}
	case TypeArrayBool:
		return map[string]any{
			"type":  "array",
			"items": map[string]any{"type": "boolean"},
		}
	case TypeArrayString:
		return map[string]any{
			"type":  "array",
			"items": map[string]any{"type": "string"},
		}
	}
	return map[string]any{"type": "string"}
}

func init() {
	rootCmd.AddCommand(schemaCmd)
	for channel, channelSpec := range exportSchema {
		for property, propertySpec := range channelSpec {
			if propertySpec.Type == "" {
				panic(fmt.Sprintf(
					"channel %q has missing type for property %q",
					channel,
					property,
				))
			}
			if !propertySpec.Type.IsValid() {
				panic(fmt.Sprintf(
					"channel %q has invalid property type %q for property %q",
					channel,
					propertySpec.Type,
					property,
				))
			}
		}
	}
}
