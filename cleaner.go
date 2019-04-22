package main

import (
	"strings"

	utils "github.com/ivanetchart/gitlab-registry-cleaner/utils"
	log "github.com/sirupsen/logrus"
	gitlab "github.com/xanzy/go-gitlab"
)

func main() {
	gitlabConfig, err := utils.LoadConfiguration().Validate()
	if err != nil {
		log.Errorln(err)
		return
	}

	if gitlabConfig.Debug {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugf("Gitlab Token: %s", gitlabConfig.Token)
	log.Debugf("Gitlab Project: %s", gitlabConfig.Project)

	git := gitlab.NewClient(nil, gitlabConfig.Token)

	repositories, _, err := git.ContainerRegistry.ListRegistryRepositories(
		gitlabConfig.Project,
		&gitlab.ListRegistryRepositoriesOptions{},
	)

	if err != nil {
		log.Errorln(err)
		return
	}

	log.Infof(
		"Found these repositories at %s (it will include also the default one for the project):\n  - %+v",
		gitlabConfig.Project,
		strings.Join(utils.GetRepositoriesNames(repositories), "\n  - "),
	)

	log.Debugf(
		"Configuration used to delete tags: %s", gitlabConfig,
	)

	for _, repo := range repositories {
		log.Infof(
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
			log.Errorln(bulkErr)
			return
		}
	}
}
