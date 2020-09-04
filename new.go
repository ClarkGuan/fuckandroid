package fa

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	rice "github.com/GeertJohan/go.rice"
)

func MakePlainLibrary(dir, path string, kotlin bool) (err error) {
	if filepath.IsAbs(path) {
		return fmt.Errorf("%q is absolute path", path)
	}

	var workspace string
	workspace, err = checkWorkspaceDir(dir)
	if err != nil {
		return err
	}

	libPath := filepath.Join(workspace, path)
	if _, err := os.Stat(libPath); !os.IsNotExist(err) {
		return errAlreadyExist
	}

	var box *rice.Box
	box, err = rice.FindBox("data")
	if err != nil {
		return err
	}

	if err = makeDirs(
		filepath.Join(libPath, "libs"),
		filepath.Join(libPath, "src/test/java"),
		filepath.Join(libPath, "src/main/java"),
	); err != nil {
		return
	}

	if err = boxCopyTemplate(box, "noandroidlib/build.gradle", filepath.Join(libPath, "build.gradle"), 0664,
		map[string]string{"Kotlin": parseBoolean(kotlin)}); err != nil {
		return
	}

	if err = appendSubProject(filepath.Join(workspace, "../settings.gradle"), libPath); err != nil {
		return
	}

	return
}

type LibraryPro struct {
	Package string
	Path    string
	Kotlin  bool
}

func MakeAndroidLibrary(dir string, lib LibraryPro) (err error) {
	var workspace string
	workspace, err = checkWorkspaceDir(dir)
	if err != nil {
		return err
	}

	if filepath.IsAbs(lib.Path) {
		return fmt.Errorf("%q is absolute path", lib.Path)
	}

	libPath := filepath.Join(workspace, lib.Path)
	if _, err := os.Stat(libPath); !os.IsNotExist(err) {
		return errAlreadyExist
	}

	var box *rice.Box
	box, err = rice.FindBox("data")
	if err != nil {
		return err
	}

	if err = makeDirs(
		filepath.Join(libPath, "libs"),
		filepath.Join(libPath, "src/androidTest/java"),
		filepath.Join(libPath, "src/test/java"),
		filepath.Join(libPath, "src/main/java"),
	); err != nil {
		return
	}

	listFix := func(l []string) []*struct {
		from string
		to   string
		perm os.FileMode
	} {
		ret := make([]*struct {
			from string
			to   string
			perm os.FileMode
		}, len(l))

		for i := range l {
			ret[i] = new(struct {
				from string
				to   string
				perm os.FileMode
			})
			ret[i].from = l[i]
			ret[i].to = filepath.Join(libPath, l[i][4:]) // lib/...
			ret[i].perm = 0664
		}
		return ret
	}

	list := listFix([]string{
		"lib/consumer-rules.pro",
		"lib/proguard-rules.pro",
	})

	if err = boxCopyAll(list, box); err != nil {
		return
	}

	if err = boxCopyTemplate(box, "lib/src/main/AndroidManifest.xml",
		filepath.Join(libPath, "src/main/AndroidManifest.xml"), 0664,
		map[string]string{"PackageName": lib.Package}); err != nil {
		return
	}

	if err = boxCopyTemplate(box, "lib/build.gradle",
		filepath.Join(libPath, "build.gradle"), 0664,
		map[string]string{"Kotlin": parseBoolean(lib.Kotlin)}); err != nil {
		return
	}

	if err = appendSubProject(filepath.Join(workspace, "../settings.gradle"), libPath); err != nil {
		return
	}

	return
}

type ApplicationPro struct {
	Name   string
	AppID  string
	Path   string
	Kotlin bool
}

func (app *ApplicationPro) GetPath() string {
	if len(app.Path) == 0 {
		return app.Name
	}
	return app.Path
}

func MakeAndroidApplication(dir string, app ApplicationPro) error {
	workspace, err := checkWorkspaceDir(dir)
	if err != nil {
		return err
	}

	if filepath.IsAbs(app.GetPath()) {
		return fmt.Errorf("%q is absolute path", app.GetPath())
	}

	appPath := filepath.Join(workspace, app.GetPath())
	if _, err := os.Stat(appPath); !os.IsNotExist(err) {
		return errAlreadyExist
	}

	var box *rice.Box
	box, err = rice.FindBox("data")
	if err != nil {
		return err
	}

	if err := makeDirs(
		filepath.Join(appPath, "libs"),
		filepath.Join(appPath, "src/androidTest/java"),
		filepath.Join(appPath, "src/test/java"),
	); err != nil {
		return err
	}

	listFix := func(l []string) []*struct {
		from string
		to   string
		perm os.FileMode
	} {
		ret := make([]*struct {
			from string
			to   string
			perm os.FileMode
		}, len(l))

		for i := range l {
			ret[i] = new(struct {
				from string
				to   string
				perm os.FileMode
			})
			ret[i].from = l[i]
			ret[i].to = filepath.Join(appPath, l[i][4:]) // app/...
			ret[i].perm = 0664
		}
		return ret
	}

	list := listFix([]string{
		"app/src/main/res/drawable/ic_launcher_background.xml",
		"app/src/main/res/drawable-v24/ic_launcher_foreground.xml",
		"app/src/main/res/layout/layout_main.xml",
		"app/src/main/res/mipmap-anydpi-v26/ic_launcher.xml",
		"app/src/main/res/mipmap-anydpi-v26/ic_launcher_round.xml",
		"app/src/main/res/mipmap-hdpi/ic_launcher.png",
		"app/src/main/res/mipmap-hdpi/ic_launcher_round.png",
		"app/src/main/res/mipmap-mdpi/ic_launcher.png",
		"app/src/main/res/mipmap-mdpi/ic_launcher_round.png",
		"app/src/main/res/mipmap-xhdpi/ic_launcher.png",
		"app/src/main/res/mipmap-xhdpi/ic_launcher_round.png",
		"app/src/main/res/mipmap-xxhdpi/ic_launcher.png",
		"app/src/main/res/mipmap-xxhdpi/ic_launcher_round.png",
		"app/src/main/res/mipmap-xxxhdpi/ic_launcher.png",
		"app/src/main/res/mipmap-xxxhdpi/ic_launcher_round.png",
		"app/src/main/res/values/colors.xml",
		"app/src/main/res/values/styles.xml",
		"app/proguard-rules.pro",
	})

	if err = boxCopyAll(list, box); err != nil {
		return err
	}

	if err = boxCopyTemplate(box, "app/src/main/res/values/strings.xml",
		filepath.Join(appPath, "src/main/res/values/strings.xml"), 0664,
		map[string]string{"ApplicationName": app.Name}); err != nil {
		return err
	}

	if err = boxCopyTemplate(box, "app/src/main/AndroidManifest.xml",
		filepath.Join(appPath, "src/main/AndroidManifest.xml"), 0664,
		map[string]string{"AndroidApplicationID": app.AppID}); err != nil {
		return err
	}

	if err = boxCopyTemplate(box, "app/build.gradle",
		filepath.Join(appPath, "build.gradle"), 0664,
		map[string]string{"AndroidApplicationID": app.AppID, "Kotlin": parseBoolean(app.Kotlin)}); err != nil {
		return err
	}

	if app.Kotlin {
		if err = boxCopyTemplate(box, "app/src/main/java/MainActivity.kt",
			filepath.Join(appPath, fmt.Sprintf("src/main/java/%s/MainActivity.kt", strings.ReplaceAll(app.AppID, ".", "/"))),
			0664, map[string]string{"AndroidApplicationID": app.AppID}); err != nil {
			return err
		}
	} else {
		if err = boxCopyTemplate(box, "app/src/main/java/MainActivity.java",
			filepath.Join(appPath, fmt.Sprintf("src/main/java/%s/MainActivity.java", strings.ReplaceAll(app.AppID, ".", "/"))),
			0664, map[string]string{"AndroidApplicationID": app.AppID}); err != nil {
			return err
		}
	}

	if err = appendSubProject(filepath.Join(workspace, "../settings.gradle"), appPath); err != nil {
		return err
	}

	return nil
}

func checkWorkspaceDir(dir string) (string, error) {
	absPath, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}
	for len(absPath) > 1 {
		ws := filepath.Join(absPath, "workspace")
		if _, err := os.Stat(ws); err == nil {
			if _, err = os.Stat(filepath.Join(absPath, "settings.gradle")); err == nil {
				return ws, nil
			}
		}
		absPath = filepath.Dir(absPath)
	}

	return "", errNoWorkspace
}

func appendSubProject(filename, subProPath string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Seek(0, 2)
	if err != nil {
		return err
	}
	_, err = f.WriteString(fmt.Sprintf("\ninclude ':%s'", gradleDir(subProPath)))
	return err
}

func gradleDir(s string) string {
	index := strings.Index(s, "workspace")
	s = s[index:]
	// Unix for '/'  Windows for '\\'
	return strings.ReplaceAll(s, string(os.PathSeparator), ":")
}

func parseBoolean(b bool) string {
	if b {
		return "true"
	}
	return ""
}
