package main

import (
    "fmt"
    "jvmgo/ch02/classpath"
    "strings"
)

func main() {
    // cmd := parseCmd()
    cmd := &Cmd{
        helpFlag:    false,
        versionFlag: false,
        cpOption:    "",
        XjreOption:  `C:\Program Files\Java\jdk1.8.0_131\jre`,
        class:       "java.lang.Object",
        args:        nil,
    }
    if cmd.versionFlag {
        fmt.Println("version 0.0.1")
    } else if cmd.helpFlag || cmd.class == "" {
        printUsage()
    } else {
        startJVM(cmd)
    }
}

func startJVM(cmd *Cmd) {
    // 获取 ClassPath
    cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
    fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.class, cmd.args)

    className := strings.Replace(cmd.class, ".", "/", -1)

    classData, _, err := cp.ReadClass(className)

    if err != nil {
        fmt.Printf("Could not find or load main class %s\n", cmd.class)
        return
    }
    fmt.Printf("class data:%v\n", classData)
}
