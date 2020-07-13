package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	fa "github.com/ClarkGuan/fuckandroid"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Please insert sub-command!")
		os.Exit(1)
	}

	var cmd string
	switch cmd = os.Args[1]; cmd {
	case "init":
		makeWorkspace(os.Args[2:], cmd)

	case "app":
		makeAndroidApplication(os.Args[2:], cmd)

	case "lib":
		makeAndroidLibrary(os.Args[2:], cmd)

	case "noandroidlib":
		makePlainLibrary(os.Args[2:], cmd)

	default:
		fmt.Fprintln(os.Stderr, "Unknown sub-command:", strconv.Quote(os.Args[1]))
		os.Exit(1)
	}

	//fmt.Printf("fuckandroid %s finished without error.\n", cmd)
}

func makeWorkspace(args []string, cmd string) {
	initFlagSet := flag.NewFlagSet(fmt.Sprintf("fuckandroid %s", cmd), flag.ExitOnError)
	var dir string
	initFlagSet.StringVar(&dir, "p", ".", "Root workspace's parent directory path")
	initFlagSet.Parse(args)

	initArgs := initFlagSet.Args()
	if len(initArgs) == 0 {
		fmt.Fprintf(os.Stderr, "Sub-command `%s` need a `name`\n", cmd)
		initFlagSet.PrintDefaults()
		os.Exit(1)
	}
	name := initArgs[0]

	if err := fa.MakeWorkspace(name, dir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func makeAndroidApplication(args []string, cmd string) {
	appFlagSet := flag.NewFlagSet(fmt.Sprintf("fuckandroid %s", cmd), flag.ExitOnError)
	var dir string
	var name string
	var appID string
	var relativePath string
	var nokotlin bool
	appFlagSet.StringVar(&dir, "p", ".", "Path to search workspace")
	appFlagSet.StringVar(&name, "name", "", "Display name of application")
	appFlagSet.StringVar(&appID, "id", "com.demo.app", "Id of android application. Default: \"com.demo.app\"")
	appFlagSet.BoolVar(&nokotlin, "nokotlin", false, "not using kotlin")
	appFlagSet.Parse(args)

	appArgs := appFlagSet.Args()
	if len(appArgs) == 0 {
		fmt.Fprintf(os.Stderr, "Sub-command `%s` need a `relativePath`\n", cmd)
		appFlagSet.PrintDefaults()
		os.Exit(1)
	}
	relativePath = appArgs[0]
	if len(name) == 0 {
		name = filepath.Base(relativePath)
	}
	if err := fa.MakeAndroidApplication(dir, fa.ApplicationPro{Name: name, AppID: appID, Path: relativePath, Kotlin: !nokotlin}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func makeAndroidLibrary(args []string, cmd string) {
	libFlagSet := flag.NewFlagSet(fmt.Sprintf("fuckandroid %s", cmd), flag.ExitOnError)
	var dir string
	var packageName string
	var relativePath string
	var nokotlin bool
	libFlagSet.StringVar(&dir, "p", ".", "Path to search workspace")
	libFlagSet.StringVar(&packageName, "pkg", "com.demo.lib", "Java package name for library. Default: \"com.demo.lib\"")
	libFlagSet.BoolVar(&nokotlin, "nokotlin", false, "not using kotlin")
	libFlagSet.Parse(args)

	appArgs := libFlagSet.Args()
	if len(appArgs) == 0 {
		fmt.Fprintf(os.Stderr, "Sub-command `%s` need a `relativePath`\n", cmd)
		libFlagSet.PrintDefaults()
		os.Exit(1)
	}
	relativePath = appArgs[0]
	if err := fa.MakeAndroidLibrary(dir, fa.LibraryPro{Package: packageName, Path: relativePath, Kotlin: !nokotlin}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func makePlainLibrary(args []string, cmd string) {
	libFlagSet := flag.NewFlagSet(fmt.Sprintf("fuckandroid %s", cmd), flag.ExitOnError)
	var dir string
	var nokotlin bool
	var relativePath string
	libFlagSet.StringVar(&dir, "p", ".", "Path to search workspace")
	libFlagSet.BoolVar(&nokotlin, "nokotlin", false, "not using kotlin")
	libFlagSet.Parse(args)

	appArgs := libFlagSet.Args()
	if len(appArgs) == 0 {
		fmt.Fprintf(os.Stderr, "Sub-command `%s` need a `relativePath`\n", cmd)
		libFlagSet.PrintDefaults()
		os.Exit(1)
	}
	relativePath = appArgs[0]
	if err := fa.MakePlainLibrary(dir, relativePath, !nokotlin); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
