package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fileInfo, _ := os.Stat("/usr/local/Cellar/go/1.10.1/src/github.com/watcher/test.file")
	// 获取文件的变更时间
	modTime := fileInfo.ModTime()
	fmt.Println(modTime)
	reload()
}

func reload() {
	ticker := time.NewTicker(time.Second * 2)
	for _ = range ticker.C {
		// 可以动态变更
		go func() {
			checkFile()
		}()
	}
}

func checkFile() {
	fileInfo, _ := os.Stat("/usr/local/Cellar/go/1.10.1/src/github.com/watcher/test.file")
	// 获取文件的变更时间
	modTime := fileInfo.ModTime().Unix()
	fmt.Println(modTime)
}
