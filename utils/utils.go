package utils

import (
	gitlab "github.com/xanzy/go-gitlab"
)

// GetRepositoriesNames ..
func GetRepositoriesNames(repositories []*gitlab.RegistryRepository) []string {
	vsm := make([]string, len(repositories))
	for i, v := range repositories {
		vsm[i] = v.Path
	}
	return vsm
}

// BuildGitlabBulkDeleteFromConf b
func BuildGitlabBulkDeleteFromConf(c *Configuration) *gitlab.DeleteRegistryRepositoryTagsOptions {
	return &gitlab.DeleteRegistryRepositoryTagsOptions{
		NameRegexp: gitlab.String(c.TagsRegexp),
		OlderThan:  gitlab.String(c.OlderThan),
		KeepN:      gitlab.Int(c.KeepN),
	}
}
