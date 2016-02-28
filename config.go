// Package prj Config contains the config reader/writer interface and any related config functions
package prj

type Config struct {
	Git struct {
		User       string
		Email      string
		GithubName string
		SSHKey     string
	}
	License struct {
		Type   string
		Author string
		Email  string
	}
	Readme struct {
		TmplFile string
	}
	Project struct {
		Title     string
		Lang      string
		Structure map[string]interface{}
	}
}
