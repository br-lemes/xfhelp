package cmd

import "testing"

var testSchema = SchemaSpec{
	"xfwm4": {
		"/general/theme": {Type: TypeString},
		"/general/zoom":  {Type: TypeInt},
	},
	"xsettings": {
		"/Xft/RGBA": {Type: TypeString},
	},
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Error("expected error but got nil")
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func fakeQuery(output string) queryFunc {
	return func(args ...string) ([]byte, error) {
		return []byte(output), nil
	}
}

func fakeQueryMap(outputs map[string]string) queryFunc {
	return func(args ...string) ([]byte, error) {
		for i, arg := range args {
			if arg == "-c" && i+1 < len(args) {
				channel := args[i+1]
				output, ok := outputs[channel]
				if ok {
					return []byte(output), nil
				}
			}
		}
		output, ok := outputs["__channels__"]
		if ok {
			return []byte(output), nil
		}
		return []byte{}, nil
	}
}
