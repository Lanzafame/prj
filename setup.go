// Package prj Setup initializes the ~/.prj directory and the config.toml with defaults
package prj

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"runtime"
)

var (
	prjdir     = ".prj"
	configfile = "config.toml"
	homepath   = ""
)

var defaultconfig = `title = "prj config"

[license]
type = "MIT"
author = <insert name here>
email = <insert email here>

[git]
name = <insert git name here>
email = <insert git email here>`

func init() {
	if runtime.GOOS == "windows" {
		prjdir = "~prj"
	}
	curuser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	homepath = curuser.HomeDir
}

// SetupPrjDir checks for an existing ~/.prj dir
// if it doesn't exist it is created and the config.toml
// file is created inside. Returns an error if dir or file
// couldn't be created.
func SetupPrjDir() error {
	ok := dirExists(prjdir, homepath)
	if !ok {
		err := createPrjDir()
		if err != nil {
			return fmt.Errorf("SetupPrjDir: %s", err)
		}
	}
	ok = configExists(configfile, homepath+"/"+prjdir)
	if !ok {
		err := createConfigFile()
		if err != nil {
			return fmt.Errorf("SetupPrjDir: %s", err)
		}
	}
	return nil
}

// createPrjDir creates the prj directory in the users home directory
func createPrjDir() error {
	err := os.Mkdir(homepath+"/"+prjdir, 0777)
	if err != nil {
		return fmt.Errorf("createPrjDir: %s", err)
	}
	return nil
}

// createConfigFile creates the config.toml file in the prj directory
func createConfigFile() error {
	err := ioutil.WriteFile(homepath+"/"+prjdir+"/"+configfile, []byte(defaultconfig), 0777)
	if err != nil {
		return fmt.Errorf("createConfigFile: %s", err)
	}
	return nil
}

// dirExists determines whether a directory exists
// at a certain path.
func dirExists(dirname, path string) bool {
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalf("dirExists: %s", err)
	}

	for _, d := range dirs {
		if d.IsDir() && d.Name() == dirname {
			return true
		}
	}
	return false
}

// configExists determines whether a config file exists
// in a certain directory.
func configExists(file, dir string) bool {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalf("configExists: %s", err)
	}

	for _, f := range files {
		if !f.IsDir() && f.Name() == file {
			return true
		}
	}
	return false
}
