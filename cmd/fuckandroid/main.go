package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	fa "github.com/ClarkGuan/fuckandroid"
)

var helpDesc = `使用方法：
    fuckandroid sub-command args...

sub-command 可以是：
	init      创建指定路径的 workspace
	app       在指定的 workspace 下创建一个 Android App 子工程
	lib       在指定的 workspace 下创建一个 Android Library 子工程
	plainlib  在指定的 workspace 下创建一个 Java 或者 kotlin 子工程
	help      打印本帮主信息

打印 sub-command 的帮助信息：
    fuckandroid sub-command --help
例如
    fuckandroid init --help`

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

	case "plainlib":
		makePlainLibrary(os.Args[2:], cmd)

	case "help", "-help", "-h", "--help", "--h":
		printHelpDesc()

	default:
		fmt.Fprintln(os.Stderr, "Unknown sub-command:", strconv.Quote(os.Args[1]))
		printHelpDesc()
		os.Exit(1)
	}

	//fmt.Printf("fuckandroid %s finished without error.\n", cmd)
}

func printHelpDesc() {
	fmt.Println(helpDesc)
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
	var kotlin bool
	appFlagSet.StringVar(&dir, "p", ".", "Path to search workspace")
	appFlagSet.StringVar(&name, "name", "", "Display name of application")
	appFlagSet.StringVar(&appID, "id", "com.demo.app", "Id of android application. Default: \"com.demo.app\"")
	appFlagSet.BoolVar(&kotlin, "kotlin", false, "using kotlin")
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
	if err := fa.MakeAndroidApplication(dir, fa.ApplicationPro{Name: name, AppID: appID, Path: relativePath, Kotlin: kotlin}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func makeAndroidLibrary(args []string, cmd string) {
	libFlagSet := flag.NewFlagSet(fmt.Sprintf("fuckandroid %s", cmd), flag.ExitOnError)
	var dir string
	var packageName string
	var relativePath string
	var kotlin bool
	libFlagSet.StringVar(&dir, "p", ".", "Path to search workspace")
	libFlagSet.StringVar(&packageName, "pkg", "com.demo.lib", "Java package name for library. Default: \"com.demo.lib\"")
	libFlagSet.BoolVar(&kotlin, "kotlin", false, "using kotlin")
	libFlagSet.Parse(args)

	appArgs := libFlagSet.Args()
	if len(appArgs) == 0 {
		fmt.Fprintf(os.Stderr, "Sub-command `%s` need a `relativePath`\n", cmd)
		libFlagSet.PrintDefaults()
		os.Exit(1)
	}
	relativePath = appArgs[0]
	if err := fa.MakeAndroidLibrary(dir, fa.LibraryPro{Package: packageName, Path: relativePath, Kotlin: kotlin}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func makePlainLibrary(args []string, cmd string) {
	libFlagSet := flag.NewFlagSet(fmt.Sprintf("fuckandroid %s", cmd), flag.ExitOnError)
	var dir string
	var kotlin bool
	var relativePath string
	libFlagSet.StringVar(&dir, "p", ".", "Path to search workspace")
	libFlagSet.BoolVar(&kotlin, "kotlin", false, "using kotlin")
	libFlagSet.Parse(args)

	appArgs := libFlagSet.Args()
	if len(appArgs) == 0 {
		fmt.Fprintf(os.Stderr, "Sub-command `%s` need a `relativePath`\n", cmd)
		libFlagSet.PrintDefaults()
		os.Exit(1)
	}
	relativePath = appArgs[0]
	if err := fa.MakePlainLibrary(dir, relativePath, kotlin); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
