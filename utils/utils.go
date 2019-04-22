package utils

import (
	gitlab "github.com/xanzy/go-gitlab"
)

// GetRepositoriesNames makes a list of repositories names
func GetRepositoriesNames(repositories []*gitlab.RegistryRepository) []string {
	vsm := make([]string, len(repositories))
	for i, v := range repositories {
		vsm[i] = v.Path
	}
	return vsm
}

// BuildGitlabBulkDeleteFromConf based on the app configuration it creates the set of options
//	to be used by the delete registry tags command
func BuildGitlabBulkDeleteFromConf(c *Configuration) *gitlab.DeleteRegistryRepositoryTagsOptions {
	return &gitlab.DeleteRegistryRepositoryTagsOptions{
		NameRegexp: gitlab.String(c.TagsRegexp),
		OlderThan:  gitlab.String(c.OlderThan),
		KeepN:      gitlab.Int(c.KeepN),
	}
}
