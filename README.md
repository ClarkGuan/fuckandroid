# fuckandroid

## 安装

fuckandroid 依赖 go.rice，需要安装

```bash
$ go get github.com/GeertJohan/go.rice
$ go get github.com/GeertJohan/go.rice/rice
```

然后

```bash
$ git clone https://github.com/ClarkGuan/fuckandroid
$ cd fuckandroid
$ ./install.sh
```

即可。

## 使用

### 1. 创建 Android 工作区目录

```bash
$ fuckandroid init [-p 父目录路径] workspaceName
```

即可创建如下 Android 工作区：

```bash
.
├── [-rw-rw-r--]  build.gradle
├── [drwxrwxr-x]  buildsystem
├── [drwxrwxr-x]  gradle
│   └── [drwxrwxr-x]  wrapper
│       ├── [-rw-rw-r--]  gradle-wrapper.jar
│       └── [-rw-rw-r--]  gradle-wrapper.properties
├── [-rw-rw-r--]  gradle.properties
├── [-rwxrwxr--]  gradlew
├── [-rw-rw-r--]  gradlew.bat
├── [-rw-rw-r--]  local.properties
├── [-rw-rw-r--]  settings.gradle
└── [drwxrwxr-x]  workspace
```

- buildsystem：存放各种脚本文件以及配置文件
- workspace：存放各种工程目录。例如 clean 架构中 presentation、domain 和 data 子工程

### 2. 创建 Android Application 子工程

```bash
$ fuckandroid newapp [-p 包含 workspace 目录路径] [-name 程序显示名称] [-id 程序唯一ID] projectPath
```

例如

```bash
$ fuckandroid newapp -name 你好 -id com.clark.app.hello hello
```

生成文件如下：

```bash
$ tree -p workspace
workspace/
└── [drwxrwxr-x]  hello
    ├── [-rw-rw-r--]  build.gradle
    ├── [drwxrwxr-x]  libs
    ├── [-rw-rw-r--]  proguard-rules.pro
    └── [drwxrwxr-x]  src
        ├── [drwxrwxr-x]  androidTest
        │   └── [drwxrwxr-x]  java
        ├── [drwxrwxr-x]  main
        │   ├── [-rw-rw-r--]  AndroidManifest.xml
        │   ├── [drwxrwxr-x]  java
        │   │   └── [drwxrwxr-x]  com
        │   │       └── [drwxrwxr-x]  clark
        │   │           └── [drwxrwxr-x]  app
        │   │               └── [drwxrwxr-x]  hello
        │   │                   └── [-rw-rw-r--]  MainActivity.java
        │   └── [drwxrwxr-x]  res
        │       ├── [drwxrwxr-x]  drawable
        │       │   └── [-rw-rw-r--]  ic_launcher_background.xml
        │       ├── [drwxrwxr-x]  drawable-v24
        │       │   └── [-rw-rw-r--]  ic_launcher_foreground.xml
        │       ├── [drwxrwxr-x]  layout
        │       │   └── [-rw-rw-r--]  layout_main.xml
        │       ├── [drwxrwxr-x]  mipmap-anydpi-v26
        │       │   ├── [-rw-rw-r--]  ic_launcher_round.xml
        │       │   └── [-rw-rw-r--]  ic_launcher.xml
        │       ├── [drwxrwxr-x]  mipmap-hdpi
        │       │   ├── [-rw-rw-r--]  ic_launcher.png
        │       │   └── [-rw-rw-r--]  ic_launcher_round.png
        │       ├── [drwxrwxr-x]  mipmap-mdpi
        │       │   ├── [-rw-rw-r--]  ic_launcher.png
        │       │   └── [-rw-rw-r--]  ic_launcher_round.png
        │       ├── [drwxrwxr-x]  mipmap-xhdpi
        │       │   ├── [-rw-rw-r--]  ic_launcher.png
        │       │   └── [-rw-rw-r--]  ic_launcher_round.png
        │       ├── [drwxrwxr-x]  mipmap-xxhdpi
        │       │   ├── [-rw-rw-r--]  ic_launcher.png
        │       │   └── [-rw-rw-r--]  ic_launcher_round.png
        │       ├── [drwxrwxr-x]  mipmap-xxxhdpi
        │       │   ├── [-rw-rw-r--]  ic_launcher.png
        │       │   └── [-rw-rw-r--]  ic_launcher_round.png
        │       └── [drwxrwxr-x]  values
        │           ├── [-rw-rw-r--]  colors.xml
        │           ├── [-rw-rw-r--]  strings.xml
        │           └── [-rw-rw-r--]  styles.xml
        └── [drwxrwxr-x]  test
            └── [drwxrwxr-x]  java

24 directories, 22 files
```

安装运行：

```bash
$ ./gradlew installDebug
```

如图：

[alt Android设备截图](doc/device-img.png)
