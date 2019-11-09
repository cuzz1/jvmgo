package main

import "fmt"
import "strings"
import "jvmgo/ch10/classpath"
import "jvmgo/ch10/rtda/heap"

func main() {
	// cmd := parseCmd()
	cmd := &Cmd{
		helpFlag:         false,
		versionFlag:      false,
		verboseClassFlag: false,
		verboseInstFlag:  false,
		cpOption:         `/home/cuzz/go/src/jvmgo/java/ch10`,
		// cpOption:         `C:\Users\cuzz\go\src\jvmgo\java\ch10`,
		XjreOption:  `/usr/local/lib/jdk1.8/jre`,
		// XjreOption: `Program Files\Java\jdk1.8.0_131\jre`,
		class: "ParseIntTest",
		args:  []string{},
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
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag) // 初始化ClassLoader

	className := strings.Replace(cmd.class, ".", "/", -1) // 获取需要加载的类名
	mainClass := classLoader.LoadClass(className)         // 加载类
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod, cmd.verboseInstFlag, cmd.args)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}
