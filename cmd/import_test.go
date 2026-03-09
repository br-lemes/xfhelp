package cmd

import (
	"reflect"
	"testing"
)

func TestAnyToXfconfBoolTrue(t *testing.T) {
	if anyToXfconf(true) != "true" {
		t.Error("expected \"true\"")
	}
}

func TestAnyToXfconfBoolFalse(t *testing.T) {
	if anyToXfconf(false) != "false" {
		t.Error("expected \"false\"")
	}
}

func TestAnyToXfconfInt(t *testing.T) {
	if anyToXfconf(float64(42)) != "42" {
		t.Error("expected \"42\"")
	}
}

func TestAnyToXfconfFloat(t *testing.T) {
	if anyToXfconf(float64(3.14)) != "3.14" {
		t.Error("expected \"3.14\"")
	}
}

func TestAnyToXfconfString(t *testing.T) {
	if anyToXfconf("hello world") != "hello world" {
		t.Error("expected \"hello world\"")
	}
}

func TestAnyToXfconfArray(t *testing.T) {
	result := anyToXfconf([]any{float64(1), float64(2), float64(3)})
	if result != "[1,2,3]" {
		t.Errorf("expected \"[1,2,3]\", got %q", result)
	}
}

func TestAnyToXfconfArrayString(t *testing.T) {
	result := anyToXfconf([]any{"foo", "bar"})
	if result != "[foo,bar]" {
		t.Errorf("expected \"[foo,bar]\", got %q", result)
	}
}

func TestValidateImportValueInt(t *testing.T) {
	err := validateImportValue("/prop", float64(42), TypeInt)
	assertNoError(t, err)
}

func TestValidateImportValueIntFloat(t *testing.T) {
	err := validateImportValue("/prop", float64(42.5), TypeInt)
	assertError(t, err)
}

func TestValidateImportValueIntWrongType(t *testing.T) {
	err := validateImportValue("/prop", "42", TypeInt)
	assertError(t, err)
}

func TestValidateImportValueFloat(t *testing.T) {
	err := validateImportValue("/prop", float64(3.14), TypeFloat)
	assertNoError(t, err)
}

func TestValidateImportValueFloatWrongType(t *testing.T) {
	err := validateImportValue("/prop", "3.14", TypeFloat)
	assertError(t, err)
}

func TestValidateImportValueBool(t *testing.T) {
	err := validateImportValue("/prop", true, TypeBool)
	assertNoError(t, err)
}

func TestValidateImportValueBoolWrongType(t *testing.T) {
	err := validateImportValue("/prop", "true", TypeBool)
	assertError(t, err)
}

func TestValidateImportValueString(t *testing.T) {
	err := validateImportValue("/prop", "hello", TypeString)
	assertNoError(t, err)
}

func TestValidateImportValueStringWrongType(t *testing.T) {
	err := validateImportValue("/prop", float64(42), TypeString)
	assertError(t, err)
}

func TestValidateImportValueArrayInt(t *testing.T) {
	err := validateImportValue(
		"/prop", []any{float64(1), float64(2)}, TypeArrayInt,
	)
	assertNoError(t, err)
}

func TestValidateImportValueArrayIntInvalid(t *testing.T) {
	err := validateImportValue("/prop", []any{float64(1), "abc"}, TypeArrayInt)
	assertError(t, err)
}

func TestValidateImportValueArrayWrongType(t *testing.T) {
	err := validateImportValue("/prop", "not_array", TypeArrayInt)
	assertError(t, err)
}

func TestValidateImportChannelNotInSchema(t *testing.T) {
	data := ConfigMap{"abobrinha": {"/some/prop": "value"}}
	err := validateImport(data, testSchema)
	assertError(t, err)
}

func TestValidateImportPropertyNotInSchema(t *testing.T) {
	data := ConfigMap{"xfwm4": {"/general/abobrinha": "value"}}
	err := validateImport(data, testSchema)
	assertError(t, err)
}

func TestValidateImportWrongType(t *testing.T) {
	data := ConfigMap{"xfwm4": {"/general/zoom": "not_an_int"}}
	err := validateImport(data, testSchema)
	assertError(t, err)
}

func TestValidateImportValid(t *testing.T) {
	data := ConfigMap{
		"xfwm4": {
			"/general/theme": "Default",
			"/general/zoom":  float64(0),
		},
	}
	err := validateImport(data, testSchema)
	assertNoError(t, err)
}

func TestApplyPropertyScalar(t *testing.T) {
	var capturedArgs []string
	query := func(args ...string) ([]byte, error) {
		capturedArgs = append([]string{"xfconf-query"}, args...)
		return nil, nil
	}
	applyProperty(query, "xfwm4", "/general/theme", "Default", TypeString)
	expected := []string{
		"xfconf-query", "-c", "xfwm4", "-n", "-p", "/general/theme",
		"-s", "Default", "-t", "string",
	}
	if !reflect.DeepEqual(capturedArgs, expected) {
		t.Errorf("expected %v, got %v", expected, capturedArgs)
	}
}

func TestApplyPropertyNew(t *testing.T) {
	var capturedArgs []string
	query := func(args ...string) ([]byte, error) {
		capturedArgs = append([]string{"xfconf-query"}, args...)
		return nil, nil
	}
	applyProperty(query, "xfwm4", "/general/theme", "Default", TypeString)
	expected := []string{
		"xfconf-query", "-c", "xfwm4", "-n", "-p", "/general/theme",
		"-s", "Default", "-t", "string",
	}
	if !reflect.DeepEqual(capturedArgs, expected) {
		t.Errorf("expected %v, got %v", expected, capturedArgs)
	}
}

func TestApplyPropertyArray(t *testing.T) {
	var capturedArgs []string
	query := func(args ...string) ([]byte, error) {
		capturedArgs = append([]string{"xfconf-query"}, args...)
		return nil, nil
	}
	applyProperty(query, "xfwm4", "/general/workspace_names",
		[]any{"Workspace 1", "Workspace 2", "Workspace 3"}, TypeArrayString)
	expected := []string{
		"xfconf-query", "-c", "xfwm4", "-n", "-p", "/general/workspace_names",
		"-s", "Workspace 1", "-s", "Workspace 2", "-s", "Workspace 3",
		"-t", "string", "-t", "string", "-t", "string",
	}
	if !reflect.DeepEqual(capturedArgs, expected) {
		t.Errorf("expected %v, got %v", expected, capturedArgs)
	}
}
