package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var depth int

func listAnyOne(path string) {
	finfo, err := os.Stat(path) // *os.fileStat
	// path不存在
	if os.IsNotExist(err) {
		fmt.Println(path, "IS NOT EXIST")
		return
	}
	//判断dir还是file
	if !finfo.IsDir() {
		fmt.Println("↑__ " + path + "*")
	} else {
		//如果path是目录，进入path
		fmt.Println("↑__ " + path + "/")
		os.Chdir(path)
		depth++
		flies, _ := filepath.Glob("*") // files = []string
		for _, f := range flies {
			for i := 1; i <= depth; i++ {
				fmt.Print("\t")
			}
			//每次遇到目录就再进一层
			listAnyOne(f)
		}
		//返回上一层目录
		os.Chdir("..")
		depth--
		// fmt.Println(os.Getwd())
	}
	return
}

func main() {
	tmpdir := flag.String("", ".", "File/Dir path") //pointer
	help := flag.Bool("h", false, "Help Info")      //pointer

	flag.Usage = func() {
		fmt.Println(`
Usage: ls [OPTION]... [FILE|DIR]...
Options:`)
		flag.PrintDefaults()
	}
	flag.Parse()
	if *help {
		flag.Usage()
	} else {
		if len(flag.Args()) == 0 {
			fmt.Println("Args is empty, listing PWD...")
			listAnyOne(*tmpdir)
		} else {
			// flag.Args() == ["a","b","c"]
			for _, arg := range flag.Args() {
				listAnyOne(arg)
			}
		}
	}
}
