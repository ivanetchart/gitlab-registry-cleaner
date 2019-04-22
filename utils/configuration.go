package utils

import (
	"flag"
	"fmt"
)

// Configuration containing the Gitlab token to access the API
//	and the project name
type Configuration struct {
	// Required fields
	Token   string
	Project string
	Debug   bool

	// Optional fields
	TagsRegexp string
	OlderThan  string
	KeepN      int
}

// Validate checks if it has the required configuration values
func (c *Configuration) Validate() (*Configuration, error) {
	if c.Project == "" {
		return nil, requiredAttrError("project")
	}

	if c.Token == "" {
		return nil, requiredAttrError("token")
	}

	return c, nil
}

func (c *Configuration) String() string {
	return fmt.Sprintf(
		"\n  - Tags Regexp: '%v'\n  - Older than: '%v'\n  - Keep last: '%v'",
		c.TagsRegexp, c.OlderThan, c.KeepN)
}

func requiredAttrError(attr string) error {
	return fmt.Errorf("--%s flag required", attr)
}

// LoadConfiguration creates a configuration struct to be used internally based on the flags passed
//	to the application CLI
func LoadConfiguration() *Configuration {
	gitlabToken := flag.String(
		"token",
		"",
		"Private token to authenticate against GitLab",
	)
	gitlabProject := flag.String(
		"project",
		"",
		"Project path including namespace (eg: my_company/software/my_project)",
	)
	debug := flag.Bool(
		"debug",
		false,
		"Set log level to debug",
	)
	tagsRegexp := flag.String(
		"tags-regexp",
		".*",
		"",
	)
	olderThan := flag.String(
		"older-than",
		"1w",
		"",
	)
	keepN := flag.Int(
		"keep-n",
		5,
		"",
	)

	flag.Parse()

	return &Configuration{
		Token:   *gitlabToken,
		Project: *gitlabProject,
		Debug:   *debug,

		TagsRegexp: *tagsRegexp,
		OlderThan:  *olderThan,
		KeepN:      *keepN,
	}
}
