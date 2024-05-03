package github

// WorkflowInputs is the comparison between the input name and its detailed content.
// please make sure key == values.Name
var WorkflowInputs = Inputs{}

var WorkflowOutputs = Outputs{}

type Output struct {
	Description string `yaml:"description"`
}

type Outputs map[string]Output
