package github

type Output struct {
	Description string `yaml:"description"`
}

type Outputs map[string]Output
