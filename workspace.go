package fa

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	rice "github.com/GeertJohan/go.rice"
)

func MakeWorkspace(name, dir string) (err error) {
	rootDir := filepath.Join(dir, name)
	if err = os.MkdirAll(rootDir, 0775); err != nil {
		return
	}
	if err = os.MkdirAll(filepath.Join(rootDir, "buildsystem"), 0775); err != nil {
		return
	}
	if err = os.MkdirAll(filepath.Join(rootDir, "workspace"), 0775); err != nil {
		return
	}
	var box *rice.Box
	box, err = rice.FindBox("data")
	if err != nil {
		return
	}

	list := []*struct {
		name string
		perm os.FileMode
	}{
		{"gradle/wrapper/gradle-wrapper.jar", 0664},
		{"gradle/wrapper/gradle-wrapper.properties", 0664},
		{".gitignore", 0664},
		{"build.gradle", 0664},
		{"gradle.properties", 0664},
		{"gradlew", 0774},
		{"gradlew.bat", 0664},
	}

	if err = boxCopyAll(list, box, rootDir); err != nil {
		return
	}

	if err = boxCopyTemplate(box, rootDir, "local.properties", 0664,
		map[string]string{"AndroidSdkHome": androidHome()}); err != nil {
		return
	}

	if err = boxCopyTemplate(box, rootDir, "settings.gradle", 0664,
		map[string]string{"RootProjectName": name}); err != nil {
		return
	}

	return
}

func boxCopyAll(list []*struct {
	name string
	perm os.FileMode
}, box *rice.Box, dir string) error {
	for _, entry := range list {
		if strings.Contains(entry.name, "/") {
			if err := os.MkdirAll(filepath.Dir(filepath.Join(dir, entry.name)), 0775); err != nil {
				return err
			}
		}
		if err := boxCopy(box, dir, entry.name, entry.perm); err != nil {
			return err
		}
	}
	return nil
}

func boxCopy(box *rice.Box, dir, name string, perm os.FileMode) error {
	return ioutil.WriteFile(filepath.Join(dir, name), box.MustBytes(name), perm)
}

func boxCopyTemplate(box *rice.Box, dir, name string, perm os.FileMode, holder map[string]string) error {
	tmpl := template.New(name)
	tmpl, _ = tmpl.Parse(box.MustString(name))
	output, err := os.OpenFile(filepath.Join(dir, name), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	defer output.Close()
	return tmpl.Execute(output, holder)
}

func androidHome() string {
	if val, found := os.LookupEnv("ANDROID_HOME"); found {
		return val
	}
	homeDir, _ := os.UserHomeDir()
	// TODO 默认路径
	return filepath.Join(homeDir, "Android/Sdk")
}
