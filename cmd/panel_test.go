package cmd

import (
	"errors"
	"testing"
)

func TestGetPanels_SinglePanel(t *testing.T) {
	output := "/panels/panel-1/output-name Primary"

	panels, err := getPanels(fakeQuery(output))
	assertNoError(t, err)

	if len(panels) != 1 {
		t.Errorf("expected 1 panel, got %d", len(panels))
	}

	if panels[1] != "Primary" {
		t.Errorf("expected output-name 'Primary', got %v", panels[1])
	}
}

func TestGetPanels_MultiplePanels(t *testing.T) {
	output := "/panels/panel-1/output-name Primary\n" +
		"/panels/panel-2/output-name HDMI2"

	panels, err := getPanels(fakeQuery(output))
	assertNoError(t, err)

	if len(panels) != 2 {
		t.Errorf("expected 2 panels, got %d", len(panels))
	}

	if panels[1] != "Primary" {
		t.Errorf(
			"expected panel 1 output-name 'Primary', got %v",
			panels[1],
		)
	}

	if panels[2] != "HDMI2" {
		t.Errorf(
			"expected panel 2 output-name 'HDMI2', got %v",
			panels[2],
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
	output := "/panels/panel-1/output-name Primary\n" +
		"invalid line without proper format\n" +
		"/panels/panel-2/output-name HDMI2\n" +
		"/panels/invalid-id/output-name Secondary\n" +
		"/panels/panel-abc/output-name Tertiary"

	panels, err := getPanels(fakeQuery(output))
	assertNoError(t, err)

	if len(panels) != 2 {
		t.Errorf("expected 2 panels, got %d", len(panels))
	}

	if panels[1] != "Primary" {
		t.Error("panel 1 not parsed correctly")
	}

	if panels[2] != "HDMI2" {
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
