# fuckandroid

## 安装

fuckandroid 依赖 go.rice，需要安装

```bash
$ go install github.com/GeertJohan/go.rice/rice@latest
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

例如

```bash
$ fuckandroid init hello
$ tree -p hello/
hello/
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

4 directories, 8 files
```

- buildsystem：存放各种脚本文件以及配置文件
- workspace：存放各种工程目录。例如 clean 架构中 presentation、domain 和 data 子工程

### 2. 创建 Android Application 子工程

```bash
$ fuckandroid app [-p 包含 workspace 目录路径] [-name 程序显示名称] [-id 程序唯一ID] [-nokotlin] projectPath
```

例如

```bash
$ fuckandroid app -name 你好 -id com.clark.app.hello -nokotlin hello
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

![alt Android设备截图](doc/device-img.png)

### 3. 创建 Android Library 子工程

```bash
$ fuckandroid lib [-p 包含 workspace 目录路径] [-pkg Java 包名] [-nokotlin] projectPath
```

例如

```bash
$ fuckandroid lib mylib
$ tree -p workspace/mylib
workspace/mylib
├── [-rw-rw-r--]  build.gradle
├── [-rw-rw-r--]  consumer-rules.pro
├── [drwxrwxr-x]  libs
├── [-rw-rw-r--]  proguard-rules.pro
└── [drwxrwxr-x]  src
    ├── [drwxrwxr-x]  androidTest
    │   └── [drwxrwxr-x]  java
    ├── [drwxrwxr-x]  main
    │   ├── [-rw-rw-r--]  AndroidManifest.xml
    │   └── [drwxrwxr-x]  java
    └── [drwxrwxr-x]  test
        └── [drwxrwxr-x]  java

8 directories, 4 files
```

### 4. 创建 Java or Kotlin Library 子工程

```bash
$ fuckandroid plainlib [-p 包含 workspace 目录路径] [-nokotlin] projectPath
```

例如

```bash
$ fuckandroid plainlib -nokotlin myjavalib
$ tree -p workspace/myjavalib/
workspace/myjavalib/
├── [-rw-rw-r--]  build.gradle
├── [drwxrwxr-x]  libs
└── [drwxrwxr-x]  src
    ├── [drwxrwxr-x]  main
    │   └── [drwxrwxr-x]  java
    └── [drwxrwxr-x]  test
        └── [drwxrwxr-x]  java

6 directories, 1 file
```
