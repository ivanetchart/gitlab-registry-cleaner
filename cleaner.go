package main

import (
	"log"
	"strings"

	utils "github.com/ivanetchart/gitlab-registry-cleaner/utils"
	gitlab "github.com/xanzy/go-gitlab"
)

func main() {
	gitlabConfig, err := utils.LoadConfiguration().Validate()
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Gitlab Token: %s", gitlabConfig.Token)
	log.Printf("Gitlab Project: %s", gitlabConfig.Project)

	git := gitlab.NewClient(nil, gitlabConfig.Token)

	repositories, _, err := git.ContainerRegistry.ListRegistryRepositories(
		gitlabConfig.Project,
		&gitlab.ListRegistryRepositoriesOptions{},
	)

	if err != nil {
		log.Println(err)
		return
	}

	log.Printf(
		"Found these repositories at %s (it will include also the default one for the project):\n  - %+v",
		gitlabConfig.Project,
		strings.Join(utils.GetRepositoriesNames(repositories), "\n  - "),
	)

	log.Printf(
		"Configuration used to delete tags: %s", gitlabConfig,
	)

	for _, repo := range repositories {
		log.Printf(
			"Starting to delete tags from repository: %s",
			repo.Path,
		)

		bulkDeleteOpts := utils.BuildGitlabBulkDeleteFromConf(gitlabConfig)

		_, bulkErr := git.ContainerRegistry.DeleteRegistryRepositoryTags(
			gitlabConfig.Project,
			repo.ID,
			bulkDeleteOpts,
		)

		if bulkErr != nil {
			log.Println(bulkErr)
			return
		}
	}
}
