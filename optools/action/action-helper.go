package main

import (
	"fmt"
	"os"

	"github.com/guguducken/action-go-template/pkg/github"
	"gopkg.in/yaml.v3"
)

var (
	Name        = ""
	Author      = ""
	Description = ""
	Icon        = ""
	Color       = ""
)

type Action struct {
	Name        string         `yaml:"name"`
	Author      string         `yaml:"author"`
	Description string         `yaml:"description"`
	Inputs      github.Inputs  `yaml:"inputs,omitempty"`
	Outputs     github.Outputs `yaml:"outputs,omitempty"`
	Runs        Runs           `yaml:"runs"`
	Branding    *Branding      `yaml:"branding,omitempty"`
}

type Runs struct {
	Using string   `yaml:"using"`
	Main  string   `yaml:"main,omitempty"`
	Image string   `yaml:"image,omitempty"`
	Args  []string `yaml:"args,omitempty"`
	Post  string   `yaml:"post,omitempty"`
}

type Branding struct {
	Icon  string `yaml:"icon"`
	Color string `yaml:"color"`
}

func main() {
	action := GenDefaultDockerAction(Name, Author, Description)
	action.Inputs = github.WorkflowInputs
	action.Outputs = github.WorkflowOutputs

	action.Runs.Args = GenDockerActionArgs(github.WorkflowInputs)

	action.Branding = &Branding{
		Icon:  Icon,
		Color: Color,
	}

	data, err := yaml.Marshal(action)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("action.yaml", data, 0644)
	if err != nil {
		panic(err)
	}
}

func GenDefaultDockerAction(name, author, description string) Action {
	return Action{
		Name:        name,
		Author:      author,
		Description: description,
		Runs: Runs{
			Using: "docker",
			Image: "Dockerfile",
			Args:  make([]string, 0),
		},
	}
}

func GenDockerActionArgs(inputs github.Inputs) []string {
	result := make([]string, 0, len(inputs))
	for _, input := range inputs {
		result = append(result, fmt.Sprintf("${{ inputs.%s }}", input.Name))
	}
	return result
}
