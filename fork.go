package fork

import (
	// "exec"
	"fmt"
)

func envToEnvv(env map[string]string) []string {
	envv := make([]string, 0, len(env))
	for name, value := range env {
		envv = append(envv, fmt.Sprintf("%s=%s", name, value))
	}
	return envv
}