package utils

import (
	gitlab "github.com/xanzy/go-gitlab"
	"reflect"
	"testing"
)

func TestGetRepositoriesNames(t *testing.T) {
	repositories := []*gitlab.RegistryRepository{
		{ID: '1', Path: "path/to/repo1"},
		{ID: '2', Path: "path/to/repo2"},
	}

	result := GetRepositoriesNames(repositories)
	expected := []string{"path/to/repo1", "path/to/repo2"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %s want %s", result, expected)
	}
}
func TestBuildGitlabBulkDeleteFromConf(t *testing.T) {
	expectedRegexp := ".*"
	expectedOlderThan := "1w"
	expectedKeepN := 5

	sampleConfiguration := &Configuration{
		Token:   "SampleToken",
		Project: "SampleProject",

		TagsRegexp: expectedRegexp,
		OlderThan:  expectedOlderThan,
		KeepN:      expectedKeepN,
	}

	result := BuildGitlabBulkDeleteFromConf(sampleConfiguration)

	if *result.NameRegexp != expectedRegexp {
		t.Errorf("got %s want %s", *result.NameRegexp, expectedRegexp)
	}

	if *result.OlderThan != expectedOlderThan {
		t.Errorf("got %s want %s", *result.OlderThan, expectedOlderThan)
	}

	if *result.KeepN != expectedKeepN {
		t.Errorf("got %d want %d", result.KeepN, expectedKeepN)
	}
}
