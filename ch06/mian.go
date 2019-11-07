package main

import (
	"fmt"
	"jvmgo/ch06/classpath"
	"jvmgo/ch06/rtda/heap"
	"strings"
)

func main() {
	// cmd := parseCmd()
	cmd := &Cmd{
		helpFlag:    false,
		versionFlag: false,
		cpOption:    `C:\Users\cuzz\go\src\jvmgo\java\ch06`,
		// XjreOption:  `/usr/local/lib/jdk1.8/jre`,
		XjreOption: `C:\Program Files\Java\jdk1.8.0_131\jre`,
		//class:      "java.lang.String",
		class: "MyObject",
		args:  nil,
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
	classLoader := heap.NewClassLoader(cp)

	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}
