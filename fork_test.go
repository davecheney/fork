package fork

import (
	"testing"
)

func TestEnvToEnvv(t *testing.T) {
	env := map[string]string { 
		"PATH": "/usr/bin",
		"USER": "foo", 
	}
	expected := []string {
		"USER=foo",
		"PATH=/usr/bin",
	}
	actual := envToEnvv(env)
	
	if len(actual) != len(expected) {
		t.Fatalf("envToEnvv() returned less elements, %d vs %d", len(actual),len(expected))
	}
	
	for index, value := range expected {
		if value != actual[index] {
			t.Fatalf("expected: %s, actual: %s", value, actual[index])
		}
	}
}