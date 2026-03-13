package cmd

import (
	"errors"
	"testing"
)

func TestGetPanels_SinglePanel(t *testing.T) {
	output := "/panels/panel-1/output-name     Primary\n" +
		"/panels/panel-1/position        p=8;x=0;y=0\n" +
		"/panels/panel-1/length          100\n" +
		"/panels/panel-1/position-locked true\n" +
		"/panels/panel-1/size            30\n" +
		"/panels/panel-1/plugin-ids      [1, 2, 3, 4, 5]"

	panels, err := getPanels(fakeQuery(output))
	assertNoError(t, err)

	if len(panels) != 1 {
		t.Errorf("expected 1 panel, got %d", len(panels))
	}

	panel := panels[1]
	if panel["output-name"] != "Primary" {
		t.Errorf("expected output-name 'Primary', got %v", panel["output-name"])
	}
	if panel["position"] != "p=8;x=0;y=0" {
		t.Errorf("expected position 'p=8;x=0;y=0', got %v", panel["position"])
	}
	if panel["length"] != 100.0 {
		t.Errorf("expected length 100.0, got %v", panel["length"])
	}
	if panel["position-locked"] != true {
		t.Errorf(
			"expected position-locked true, got %v",
			panel["position-locked"],
		)
	}
	if panel["size"] != int64(30) {
		t.Errorf("expected size 30, got %v", panel["size"])
	}
}

func TestGetPanels_MultiplePanels(t *testing.T) {
	output := "/panels/panel-1/output-name Primary\n" +
		"/panels/panel-1/position    p=8;x=0;y=0\n" +
		"/panels/panel-1/length      100\n" +
		"/panels/panel-2/output-name HDMI2\n" +
		"/panels/panel-2/position    p=10;x=1920;y=0\n" +
		"/panels/panel-2/length      75\n" +
		"/panels/panel-2/size        24"

	panels, err := getPanels(fakeQuery(output))
	assertNoError(t, err)

	if len(panels) != 2 {
		t.Errorf("expected 2 panels, got %d", len(panels))
	}

	panel1 := panels[1]
	if panel1["output-name"] != "Primary" {
		t.Errorf(
			"expected panel 1 output-name 'Primary', got %v",
			panel1["output-name"],
		)
	}

	panel2 := panels[2]
	if panel2["output-name"] != "HDMI2" {
		t.Errorf(
			"expected panel 2 output-name 'HDMI2', got %v",
			panel2["output-name"],
		)
	}
	if panel2["position"] != "p=10;x=1920;y=0" {
		t.Errorf(
			"expected panel 2 position 'p=10;x=1920;y=0', got %v",
			panel2["position"],
		)
	}
}

func TestGetPanels_EmptyOutput(t *testing.T) {
	panels, err := getPanels(fakeQuery(""))
	assertNoError(t, err)

	if len(panels) != 0 {
		t.Errorf("expected 0 panels, got %d", len(panels))
	}
}

func TestGetPanels_InvalidLines(t *testing.T) {
	output := "/panels/panel-1/output-name    Primary\n" +
		"invalid line without proper    format\n" +
		"/panels/panel-2/position       p=10;x=1920;y=0\n" +
		"/panels/invalid-id/output-name Secondary\n" +
		"/panels/panel-abc/output-name  Tertiary"

	panels, err := getPanels(fakeQuery(output))
	assertNoError(t, err)

	if len(panels) != 2 {
		t.Errorf("expected 2 panels, got %d", len(panels))
	}

	if panels[1] == nil || panels[1]["output-name"] != "Primary" {
		t.Error("panel 1 not parsed correctly")
	}

	if panels[2] == nil || panels[2]["position"] != "p=10;x=1920;y=0" {
		t.Error("panel 2 not parsed correctly")
	}
}

func TestGetPanels_QueryError(t *testing.T) {
	_, err := getPanels(func(args ...string) ([]byte, error) {
		return nil, errors.New("test error")
	})
	assertError(t, err)
}

func TestSetPanelOutput_InvalidPanelID(t *testing.T) {
	err := setPanelOutput(fakeQuery(""), 999, "Primary")
	assertError(t, err)
}

func TestSetPanelOutput_Success(t *testing.T) {
	err := setPanelOutput(
		fakeQuery("/panels/panel-1/output-name  Primary"), 1, "HDMI2",
	)
	assertNoError(t, err)
}
