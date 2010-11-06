package fork

import (
	"exec"
	"fmt"
	"os"
)

// convert from the more accessible env map to a list of strings in env format
func envToEnvv(env map[string]string) []string {
	envv := make([]string, 0, len(env))
	for name, value := range env {
		envv = append(envv, fmt.Sprintf("%s=%s", name, value))
	}
	return envv
}

// fork and exec
func Exec(s string, args []string, env map[string]string) (*os.Waitmsg, os.Error) {
	cmd, err := exec.Run(s, args, envToEnvv(env), ".", exec.PassThrough, exec.PassThrough, exec.PassThrough)
	if err != nil {
		return nil, err
	}
	msg, err := cmd.Wait(0)
	if err != nil {
		return nil, err
	}
	defer cmd.Close()
	return msg, nil
}