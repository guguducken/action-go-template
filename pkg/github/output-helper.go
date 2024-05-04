package github

import (
	"errors"
	"fmt"
	"os"
)

type Output struct {
	Description string `yaml:"description"`
}

type Outputs map[string]Output

func SetOutput(key, value string) error {
	outputs := os.Getenv("GITHUB_OUTPUT")
	if outputs == "" {
		return deprecatedSetOutput(key, value)
	}
	f, err := os.OpenFile(outputs, os.O_APPEND, os.ModeAppend)
	if err != nil {
		return errors.New("failed to open output file")
	}
	_, err = f.WriteString(prepareKeyValueMessage(key, value))
	return err
}

// deprecated:
func deprecatedSetOutput(key, value string) error {
	return deprecatedCommand("set-output", key, value)
}

func deprecatedCommand(command, key, value string) error {
	_, err := os.Stdout.WriteString(fmt.Sprintf("\n::%s name={%s}::{%s}\n", command, key, value))
	return err
}
