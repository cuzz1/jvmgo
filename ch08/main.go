package main

import "fmt"
import "strings"
import "jvmgo/ch08/classpath"
import "jvmgo/ch08/rtda/heap"

func main() {
	// cmd := parseCmd()
	cmd := &Cmd{
		helpFlag:         false,
		versionFlag:      false,
		verboseClassFlag: false,
		verboseInstFlag:  false,
		cpOption:         `C:\Users\cuzz\go\src\jvmgo\java\ch08`,
		// XjreOption:  `/usr/local/lib/jdk1.8/jre`,
		XjreOption: `C:\Program Files\Java\jdk1.8.0_131\jre`,
		//class:      "java.lang.String",
		class: "HelloWorld",
		args:  []string{"hello, world", "你好, 世界"},
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
