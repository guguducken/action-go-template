package github

// WorkflowInputs allow you to specify data that the action expects to use during runtime.
// GitHub stores input parameters as environment variables.
var WorkflowInputs = Inputs{}

// WorkflowOutputs allow you to declare data that an action sets.
// Actions that run later in a workflow can use the output data set in previously run actions.
var WorkflowOutputs = Outputs{}
