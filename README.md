# fuckandroid

## 安装

```bash
$ go get github.com/ClarkGuan/fuckandroid/...
```

## 使用

### 1. 创建 Android 工作区目录

```bash
$ fuckandroid init [-p dir] hello
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
$ fuckandroid new 
```