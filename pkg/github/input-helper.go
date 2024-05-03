package github

import (
	"fmt"
	"os"
	"strings"
)

type Input struct {
	Name        string `yaml:"-"`
	Description string `yaml:"description"`
	Required    bool   `yaml:"required"`
	Default     string `yaml:"default"`
}

type Inputs map[string]Input

func GetInput(name string) string {
	name = fmt.Sprintf("INPUT_%s", strings.ToUpper(strings.ReplaceAll(name, " ", "_")))
	return os.Getenv(name)
}

func MustGetInput(name string) string {
	result := GetInput(name)
	if result == "" {
		panic("cannot get input: " + name)
	}
	return result
}
