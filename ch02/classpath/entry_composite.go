package classpath

import (
    "errors"
    "os"
    "path/filepath"
    "strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
    var compositeEntry []Entry
    for _, path := range strings.Split(pathList, pathListSeparator) {
        entry := newEntry(path)
        compositeEntry = append(compositeEntry, entry)
    }
    return compositeEntry
}

func newWildcardEntry(path string) CompositeEntry {
    baseDir := path[:len(path)-1] // remove *
    compositeEntry := []Entry{}
    walkFn := func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() && path != baseDir {
            return filepath.SkipDir
        }
        if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
            jarEntry := newZipEntry(path)
            compositeEntry = append(compositeEntry, jarEntry)
        }
        return nil
    }
    filepath.Walk(baseDir, walkFn)
    return compositeEntry
}

func (self CompositeEntry) ReadClass(className string) ([]byte, Entry, error) {
    for _, entry := range self {
        data, from, err := entry.ReadClass(className)
        if err == nil {
            return data, from, nil
        }
    }
    return nil, nil, errors.New("class not found: " + className)
}
func (self CompositeEntry) String() string {
    strs := make([]string, len(self))
    for i, entry := range self {
        strs[i] = entry.String()
    }
    return strings.Join(strs, pathListSeparator)
}
