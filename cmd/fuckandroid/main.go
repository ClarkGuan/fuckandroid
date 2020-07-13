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

	switch os.Args[1] {
	case "init":
		makeWorkspace(os.Args[2:])

	case "newapp":
		makeAndroidApplication(os.Args[2:])

	case "newlib":
		makeAndroidLibrary(os.Args[2:])

	case "newjavalib":
		makeJavaLibrary(os.Args[2:])

	default:
		fmt.Fprintln(os.Stderr, "Unknown sub-command:", strconv.Quote(os.Args[1]))
	}
}

func makeWorkspace(args []string) {
	initFlagSet := flag.NewFlagSet("fuckandroid init", flag.ExitOnError)
	var dir string
	initFlagSet.StringVar(&dir, "p", ".", "Root workspace's parent directory path")
	initFlagSet.Parse(args)

	initArgs := initFlagSet.Args()
	if len(initArgs) == 0 {
		fmt.Fprintln(os.Stderr, "Sub-command `init` need a `name`")
		initFlagSet.PrintDefaults()
		os.Exit(1)
	}
	name := initArgs[0]

	if err := fa.MakeWorkspace(name, dir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func makeAndroidApplication(args []string) {
	appFlagSet := flag.NewFlagSet("fuckandroid newapp", flag.ExitOnError)
	var dir string
	var name string
	var appID string
	var relativePath string
	appFlagSet.StringVar(&dir, "p", ".", "Path to search workspace")
	appFlagSet.StringVar(&name, "name", "", "Display name of application")
	appFlagSet.StringVar(&appID, "id", "com.demo.app", "Id of android application. Default: \"com.demo.app\"")
	appFlagSet.Parse(args)

	appArgs := appFlagSet.Args()
	if len(appArgs) == 0 {
		fmt.Fprintln(os.Stderr, "Sub-command `newapp` need a `relativePath`")
		appFlagSet.PrintDefaults()
		os.Exit(1)
	}
	relativePath = appArgs[0]
	if len(name) == 0 {
		name = filepath.Base(relativePath)
	}
	if err := fa.MakeAndroidApplication(dir, fa.ApplicationPro{Name: name, AppID: appID, Path: relativePath}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func makeAndroidLibrary(args []string) {
	libFlagSet := flag.NewFlagSet("fuckandroid newlib", flag.ExitOnError)
	var dir string
	var packageName string
	var relativePath string
	libFlagSet.StringVar(&dir, "p", ".", "Path to search workspace")
	libFlagSet.StringVar(&packageName, "pkg", "com.demo.lib", "Java package name for library. Default: \"com.demo.lib\"")
	libFlagSet.Parse(args)

	appArgs := libFlagSet.Args()
	if len(appArgs) == 0 {
		fmt.Fprintln(os.Stderr, "Sub-command `newlib` need a `relativePath`")
		libFlagSet.PrintDefaults()
		os.Exit(1)
	}
	relativePath = appArgs[0]
	if err := fa.MakeAndroidLibrary(dir, fa.LibraryPro{Package: packageName, Path: relativePath}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func makeJavaLibrary(args []string) {
	libFlagSet := flag.NewFlagSet("fuckandroid newjavalib", flag.ExitOnError)
	var dir string
	var relativePath string
	libFlagSet.StringVar(&dir, "p", ".", "Path to search workspace")
	libFlagSet.Parse(args)

	appArgs := libFlagSet.Args()
	if len(appArgs) == 0 {
		fmt.Fprintln(os.Stderr, "Sub-command `newjavalib` need a `relativePath`")
		libFlagSet.PrintDefaults()
		os.Exit(1)
	}
	relativePath = appArgs[0]
	if err := fa.MakeJavaLibrary(dir, relativePath); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
