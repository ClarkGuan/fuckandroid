package fa

import (
	"os"
	"path/filepath"

	rice "github.com/GeertJohan/go.rice"
)

func MakeWorkspace(name, dir string) (err error) {
	rootDir := filepath.Join(dir, name)

	if _, err = os.Stat(rootDir); !os.IsNotExist(err) {
		return errAlreadyExist
	}

	if err = makeDirs(
		filepath.Join(rootDir, "buildsystem"),
		filepath.Join(rootDir, "workspace"),
	); err != nil {
		return
	}

	var box *rice.Box
	box, err = rice.FindBox("data")
	if err != nil {
		return
	}

	listFix := func(l []*struct {
		name string
		perm os.FileMode
	}) (ret []*struct {
		from string
		to   string
		perm os.FileMode
	}) {
		for _, i := range l {
			s := struct {
				from string
				to   string
				perm os.FileMode
			}{i.name, filepath.Join(rootDir, i.name), i.perm}
			ret = append(ret, &s)
		}
		return
	}

	list := listFix([]*struct {
		name string
		perm os.FileMode
	}{
		{"gradle/wrapper/gradle-wrapper.jar", 0664},
		{"gradle/wrapper/gradle-wrapper.properties", 0664},
		{"gradle/libs.versions.toml", 0664},
		{".gitignore", 0664},
		{"build.gradle", 0664},
		{"gradle.properties", 0664},
		{"gradlew", 0774},
		{"gradlew.bat", 0664},
	})

	if err = boxCopyAll(list, box); err != nil {
		return
	}

	if err = boxCopyTemplate(box, "local.properties", filepath.Join(rootDir, "local.properties"), 0664,
		map[string]string{"AndroidSdkHome": androidHome()}); err != nil {
		return
	}

	if err = boxCopyTemplate(box, "settings.gradle", filepath.Join(rootDir, "settings.gradle"), 0664,
		map[string]string{"RootProjectName": name}); err != nil {
		return
	}

	return
}

func androidHome() string {
	if val, found := os.LookupEnv("ANDROID_HOME"); found {
		return val
	}
	homeDir, _ := os.UserHomeDir()
	// TODO 默认路径
	return filepath.Join(homeDir, "Android/sdk")
}
