package main

import (
    "fmt"
    "os"
)

func exists(path string) bool {
    if _, err := os.Stat(path); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}

func main() {
    b := exists(`‪F:\downloads\阿里巴巴Java开发...1528268103.pdf`)
    fmt.Println(b)
}
