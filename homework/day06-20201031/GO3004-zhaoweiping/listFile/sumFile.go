package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func flagMsg() string {
	//定义命令行参数方式1
	var srcFile string
	flag.StringVar(&srcFile, "s", "", "源文件")
	//解析命令行参数
	flag.Parse()

	if srcFile == "" {
		flag.Usage()
		os.Exit(1)
	}
	return srcFile
}
func fileSuffix(filename string) bool {
	fFix := path.Ext(filename)
	if fFix == ".go" || fFix == ".cgo" {
		return true
	} else {
		return false
	}
}
func lsFile(srcFile string) {
	f, err := os.Stat(srcFile)
	if err != nil {
		fmt.Println(err)
	}
	//判断是否目录
	if f.IsDir() {
		file, err := os.Open(srcFile)
		if err != nil {
			fmt.Println(err)
		}
		nf, err := file.Readdirnames(0)
		if err != nil {
			fmt.Println(err)
		}
		for _, v := range nf {
			nnf := path.Join(srcFile, v)
			lsFile(nnf)
		}
	} else {
		if fileSuffix(srcFile) {
			num := countFileLine(srcFile)

			fmt.Printf("%v：%v\n", f.Name(), num)
		}
	}
}

func countFileLine(filename string) (count int64) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	count = 0
	for i := 0; i < len(data); i++ {
		if data[i] == '\n' {
			count++
		}
	}
	return
}

func main() {
	srcFile := flagMsg()
	lsFile(srcFile)
}
