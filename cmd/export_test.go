package cmd

import (
	"reflect"
	"testing"
)

func TestConvertValueInt(t *testing.T) {
	result, err := convertValue("/prop", "42", TypeInt)
	assertNoError(t, err)
	if result != int64(42) {
		t.Errorf("expected int64(42), got %v (%T)", result, result)
	}
}

func TestConvertValueIntInvalid(t *testing.T) {
	_, err := convertValue("/prop", "abc", TypeInt)
	assertError(t, err)
}

func TestConvertValueIntFloat(t *testing.T) {
	_, err := convertValue("/prop", "3.14", TypeInt)
	assertError(t, err)
}

func TestConvertValueFloat(t *testing.T) {
	result, err := convertValue("/prop", "3.14", TypeFloat)
	assertNoError(t, err)
	if result != float64(3.14) {
		t.Errorf("expected float64(3.14), got %v (%T)", result, result)
	}
}

func TestConvertValueFloatInvalid(t *testing.T) {
	_, err := convertValue("/prop", "abc", TypeFloat)
	assertError(t, err)
}

func TestConvertValueBoolTrue(t *testing.T) {
	result, err := convertValue("/prop", "true", TypeBool)
	assertNoError(t, err)
	if result != true {
		t.Errorf("expected true, got %v (%T)", result, result)
	}
}

func TestConvertValueBoolFalse(t *testing.T) {
	result, err := convertValue("/prop", "false", TypeBool)
	assertNoError(t, err)
	if result != false {
		t.Errorf("expected false, got %v (%T)", result, result)
	}
}

func TestConvertValueBoolInvalid(t *testing.T) {
	_, err := convertValue("/prop", "abc", TypeBool)
	assertError(t, err)
}

func TestConvertValueString(t *testing.T) {
	result, err := convertValue("/prop", "hello world", TypeString)
	assertNoError(t, err)
	if result != "hello world" {
		t.Errorf("expected \"hello world\", got %v (%T)", result, result)
	}
}

func TestConvertValueScalarGotArray(t *testing.T) {
	_, err := convertValue("/prop", "[1,2,3]", TypeInt)
	assertError(t, err)
}

func TestConvertValueArrayGotScalar(t *testing.T) {
	_, err := convertValue("/prop", "42", TypeArrayInt)
	assertError(t, err)
}

func TestConvertValueArrayInt(t *testing.T) {
	result, err := convertValue("/prop", "[1,2,3]", TypeArrayInt)
	assertNoError(t, err)
	expected := []any{int64(1), int64(2), int64(3)}
	slice, ok := result.([]any)
	if !ok {
		t.Fatalf("expected []any, got %T", result)
	}
	if len(slice) != len(expected) {
		t.Fatalf("expected len %d, got %d", len(expected), len(slice))
	}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("index %d: expected %v, got %v", i, v, slice[i])
		}
	}
}

func TestConvertValueArrayFloat(t *testing.T) {
	result, err := convertValue("/prop", "[1.1,2.2,3.3]", TypeArrayFloat)
	assertNoError(t, err)
	expected := []any{float64(1.1), float64(2.2), float64(3.3)}
	slice, ok := result.([]any)
	if !ok {
		t.Fatalf("expected []any, got %T", result)
	}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("index %d: expected %v, got %v", i, v, slice[i])
		}
	}
}

func TestConvertValueArrayBool(t *testing.T) {
	result, err := convertValue("/prop", "[true,false]", TypeArrayBool)
	assertNoError(t, err)
	expected := []any{true, false}
	slice, ok := result.([]any)
	if !ok {
		t.Fatalf("expected []any, got %T", result)
	}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("index %d: expected %v, got %v", i, v, slice[i])
		}
	}
}

func TestConvertValueArrayString(t *testing.T) {
	result, err := convertValue("/prop", "[foo,bar]", TypeArrayString)
	assertNoError(t, err)
	expected := []any{"foo", "bar"}
	slice, ok := result.([]any)
	if !ok {
		t.Fatalf("expected []any, got %T", result)
	}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("index %d: expected %v, got %v", i, v, slice[i])
		}
	}
}

func TestConvertValueArrayIntInvalid(t *testing.T) {
	_, err := convertValue("/prop", "[1,abc,3]", TypeArrayInt)
	assertError(t, err)
}

func TestGetPropertiesMultiple(t *testing.T) {
	query := fakeQuery(
		"/general/theme           Default\n" +
			"/general/use-compositing true\n" +
			"/general/zoom-desktop    0",
	)
	result, err := getProperties(query, "xfwm4")
	assertNoError(t, err)
	expected := map[string]string{
		"/general/theme":           "Default",
		"/general/use-compositing": "true",
		"/general/zoom-desktop":    "0",
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGetPropertiesValueWithSpaces(t *testing.T) {
	query := fakeQuery("/general/theme           Default")
	result, err := getProperties(query, "xfwm4")
	assertNoError(t, err)
	expected := map[string]string{
		"/general/theme": "Default",
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGetPropertiesIgnoresShortLines(t *testing.T) {
	query := fakeQuery(
		"/general/theme           Default\n" +
			"/general/empty\n" +
			"/general/zoom-desktop    0",
	)
	result, err := getProperties(query, "xfwm4")
	assertNoError(t, err)
	expected := map[string]string{
		"/general/theme":        "Default",
		"/general/zoom-desktop": "0",
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGetChannelsWithHeader(t *testing.T) {
	query := fakeQuery(
		"Channels:\n" +
			"  xfwm4\n" +
			"  xsettings\n" +
			"  keyboards",
	)
	result, err := getChannels(query)
	assertNoError(t, err)
	expected := []string{"xfwm4", "xsettings", "keyboards"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGetChannelsWithoutHeader(t *testing.T) {
	query := fakeQuery(
		"  xfwm4\n" +
			"  xsettings\n" +
			"  keyboards",
	)
	result, err := getChannels(query)
	assertNoError(t, err)
	expected := []string{"xfwm4", "xsettings", "keyboards"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGetListed(t *testing.T) {
	query := fakeQueryMap(map[string]string{
		"xfwm4": "/general/theme           Default\n" +
			"/general/zoom             0\n" +
			"/general/untracked         something",
		"xsettings": "/Xft/RGBA                 none",
	})
	result, err := getTracked(query, testSchema)
	assertNoError(t, err)
	expected := ConfigMap{
		"xfwm4": {
			"/general/theme": "Default",
			"/general/zoom":  int64(0),
		},
		"xsettings": {
			"/Xft/RGBA": "none",
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGetListedMissingProperty(t *testing.T) {
	query := fakeQueryMap(map[string]string{
		"xfwm4":     "/general/theme           Default",
		"xsettings": "",
	})
	result, err := getTracked(query, testSchema)
	assertNoError(t, err)
	expected := ConfigMap{
		"xfwm4": {
			"/general/theme": "Default",
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGetUntracked(t *testing.T) {
	query := fakeQueryMap(map[string]string{
		"__channels__": "xfwm4\nxsettings\nkeyboards",
		"xfwm4": "/general/theme           Default\n" +
			"/general/untracked         something",
		"xsettings": "/Xft/RGBA                 none\n" +
			"/Xft/untracked             other",
		"keyboards": "/Default/KeyRepeat        true",
	})
	result, err := getUntracked(query, testSchema)
	assertNoError(t, err)
	expected := ConfigMap{
		"xfwm4": {
			"/general/untracked": "something",
		},
		"xsettings": {
			"/Xft/untracked": "other",
		},
		"keyboards": {
			"/Default/KeyRepeat": "true",
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
