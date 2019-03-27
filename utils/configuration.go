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

// LoadConfiguration conf
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
	tagsRegexp := flag.String(
		"tags-regexp",
		".*",
		"Project path including namespace (eg: my_company/software/my_project)",
	)
	olderThan := flag.String(
		"older-than",
		"1w",
		"Project path including namespace (eg: my_company/software/my_project)",
	)
	keepN := flag.Int(
		"keep-n",
		5,
		"Project path including namespace (eg: my_company/software/my_project)",
	)

	flag.Parse()

	return &Configuration{
		Token:   *gitlabToken,
		Project: *gitlabProject,

		TagsRegexp: *tagsRegexp,
		OlderThan:  *olderThan,
		KeepN:      *keepN,
	}
}
