package utils

import "testing"

func TestValidateConfiguration(t *testing.T) {
	confWithoutToken := &Configuration{
		Project: "SampleProject",
	}

	_, errWithoutToken := confWithoutToken.Validate()

	if errWithoutToken == nil {
		t.Errorf("Expected an error when validating a configuration without token")
	}

	confWithoutProject := &Configuration{
		Token: "Token",
	}

	_, errWithoutProj := confWithoutProject.Validate()

	if errWithoutProj == nil {
		t.Errorf("Expected an error when validating a configuration without project")
	}

	okConf := &Configuration{
		Token:   "Token",
		Project: "SampleProject",
	}

	_, err := okConf.Validate()

	if err != nil {
		t.Errorf("Configuration with project and token should be valid")
	}
}

func TestLoadConfiguration(t *testing.T) {
	conf := LoadConfiguration()

	if conf == nil {
		t.Errorf("LoadConfiguration should return a configuration")
	}
}
