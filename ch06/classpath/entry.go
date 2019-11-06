package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	// 相对路径 java/lang/Object.class
	ReadClass(className string) (data []byte, entry Entry, err error)
	String() string
}

func newEntry(path string) Entry {
	//if strings.Contains(path, pathListSeparator) {
	//   return newCompositeEntry(path)
	//}
	//if strings.HasSuffix(path, "*") {
	//   return newWildcardEntry(path)
	//}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") || strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
