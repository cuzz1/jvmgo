package classpath

import (
	"os"
	"path/filepath"
)

type ClassPath struct {
	bootClassPath Entry
	extClassPath  Entry
	userClassPath Entry
}

// 双亲委派
func (self *ClassPath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClassPath.ReadClass(className); err == nil {
		return data, entry, err
	}

	if data, entry, err := self.extClassPath.ReadClass(className); err == nil {
		return data, entry, err
	}

	return self.userClassPath.ReadClass(className)
}

func (self *ClassPath) String() string {
	return self.userClassPath.String()
}

func Parse(jreOption, cpOption string) *ClassPath {
	cp := &ClassPath{}

	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}
func (self *ClassPath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClassPath = newWildcardEntry(jreLibPath)

	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClassPath = newWildcardEntry(jreExtPath)

}

func (self *ClassPath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}

	self.userClassPath = newEntry(cpOption)
}

func getJreDir(jreOption string) string {
	// -Xjre 参数作为 jre 目录
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	// 如果找不到, 就找当前目录下 jre
	if exists("./jre") {
		return "./jre"
	}
	// 还找不到, 就寻找 JAVA_HOME 中的 jre
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

// 用于判断目录是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
