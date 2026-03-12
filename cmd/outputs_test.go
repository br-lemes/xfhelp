package cmd

import "testing"

func TestGetActiveProfile_ReturnsProfile(t *testing.T) {
	got, err := getActiveProfile(fakeQuery("Default\n"))
	assertNoError(t, err)
	if got != "Default" {
		t.Errorf("expected Default, got %q", got)
	}
}

func TestGetActiveProfile_TrimsWhitespace(t *testing.T) {
	got, err := getActiveProfile(fakeQuery("  MyProfile  \n"))
	assertNoError(t, err)
	if got != "MyProfile" {
		t.Errorf("expected MyProfile, got %q", got)
	}
}

func TestGetActiveProfile_EmptyFallsBackToDefault(t *testing.T) {
	got, err := getActiveProfile(fakeQuery(""))
	assertNoError(t, err)
	if got != "Default" {
		t.Errorf("expected Default fallback, got %q", got)
	}
}

func TestGetActiveProfile_WhitespaceFallsBackToDefault(t *testing.T) {
	got, err := getActiveProfile(fakeQuery("   \n"))
	assertNoError(t, err)
	if got != "Default" {
		t.Errorf("expected Default fallback, got %q", got)
	}
}

func TestGetOutputs_AlwaysIncludesAutomaticAndPrimary(t *testing.T) {
	got, err := getOutputs(fakeQuery(""), "Default")
	assertNoError(t, err)
	if len(got) < 2 || got[0] != "Automatic" || got[1] != "Primary" {
		t.Errorf("expected first two outputs to be Automatic and Primary, got %v", got)
	}
}

func TestGetOutputs_DetectsActiveMonitors(t *testing.T) {
	xfconfOutput := "/Default/HDMI-1/Active  true\n" +
		"/Default/eDP-1/Active   true\n" +
		"/Default/VGA-1/Active   false\n"
	got, err := getOutputs(fakeQuery(xfconfOutput), "Default")
	assertNoError(t, err)

	want := []string{"Automatic", "Primary", "HDMI-1", "eDP-1"}
	if len(got) != len(want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
	for i, w := range want {
		if got[i] != w {
			t.Errorf("index %d: expected %q, got %q", i, w, got[i])
		}
	}
}

func TestGetOutputs_IgnoresInactiveMonitors(t *testing.T) {
	xfconfOutput := "/Default/VGA-1/Active   false\n" +
		"/Default/DVI-1/Active   false\n"
	got, err := getOutputs(fakeQuery(xfconfOutput), "Default")
	assertNoError(t, err)

	if len(got) != 2 {
		t.Errorf("expected only Automatic and Primary, got %v", got)
	}
}

func TestGetOutputs_IgnoresNonActiveProperties(t *testing.T) {
	xfconfOutput := "/Default/HDMI-1/Active       true\n" +
		"/Default/HDMI-1/Resolution  1920x1080\n" +
		"/Default/HDMI-1/Rotation    0\n"
	got, err := getOutputs(fakeQuery(xfconfOutput), "Default")
	assertNoError(t, err)

	want := []string{"Automatic", "Primary", "HDMI-1"}
	if len(got) != len(want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestGetOutputs_EmptyOutputReturnsOnlyDefaults(t *testing.T) {
	got, err := getOutputs(fakeQuery(""), "Default")
	assertNoError(t, err)

	if len(got) != 2 {
		t.Errorf("expected 2 defaults, got %v", got)
	}
}

func TestGetOutputs_RespectsProfile(t *testing.T) {
	xfconfOutput := "/Office/HDMI-1/Active  true\n" +
		"/Home/eDP-1/Active    true\n"
	got, err := getOutputs(fakeQuery(xfconfOutput), "Home")
	assertNoError(t, err)

	want := []string{"Automatic", "Primary", "eDP-1"}
	if len(got) != len(want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
	if got[2] != "eDP-1" {
		t.Errorf("expected eDP-1, got %q", got[2])
	}
}
