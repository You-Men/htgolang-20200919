package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
)

/*
cmds: ls, put, get, rm
path: /tmp/xxx
file:
*/

var (
	addr          = ":8889"
	servPath      = "/opt/tmp/"
	contentLenStr = 5
)

type ResponseBody struct {
	Cmd      string `json:"cmd"`
	FilePath string `json:"filePath"`
	FileName string `json:"fileName"`
	FileSize int64  `json:"fileSize"`
	Status   int    `json:"status"`
}

func main() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server starts up, listenling: %s\n", addr)

	//for {
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}

	rec := ReadHead(conn)
	fmt.Println("cmd from client: ", rec)
	if rec.Cmd == "put" {
		// HandlePUT()
		// make([]byte, rec.FileSize)
		// conn.Read()
	}
	if rec.Cmd == "ls" {
		// HandleLS()
		// ListFiles(filepath.Join(rec.FilePath, rec.FileName))
		// or ListFiles("/tmp/")
	}
	if rec.Cmd == "get" {
		// HandleGET()
	}
	if rec.Cmd == "rm" {
		// HandleRM()
	}
	WriteHead(conn, rec)
	//WriteHeadLen(conn, rec)
	//WriteHeadBody(conn, rec)

	//HandleLS(conn, &rec)
	HandleGET(conn, &rec)
	conn.Close()
	fmt.Println("Server closed.")

}

// =========== protocol ===========
// WriteHeadLen ...
func WriteHeadLen(c net.Conn, response ResponseBody) {
	bt, err := json.Marshal(response)
	if err != nil {
		c.Close()
		panic(err)
	}
	contentLen := len(string(bt))
	lenStr := fmt.Sprintf("%05d", contentLen)
	_, errW := c.Write([]byte(lenStr))
	if errW != nil {
		c.Close()
		panic(errW)
	}
}

// WriteHead wrap WriteHeadBody and WriteHeadLen
func WriteHead(c net.Conn, resBody ResponseBody) {
	WriteHeadLen(c, resBody)
	WriteHeadBody(c, resBody)
}

// WriteHeadBody send responseBody to client
func WriteHeadBody(c net.Conn, resBody ResponseBody) {
	b, errMar := json.Marshal(resBody)
	if errMar != nil {
		panic(errMar)
	}
	_, errW := c.Write(b)
	if errW != nil {
		c.Close()
		panic(errW)
	}
}

// ReadHeadLen Do not use alone, read json head len from client
func ReadHeadLen(c net.Conn) int {
	var buf = make([]byte, contentLenStr)
	_, errRead := c.Read(buf)
	if errRead != nil {
		c.Close()
		panic(errRead)
	}
	len, err := strconv.Atoi(string(buf))
	if err != nil {
		panic(err)
	}
	return len
}

// ReadHeadBody read json body from client
func ReadHead(c net.Conn) ResponseBody {
	conLen := ReadHeadLen(c)
	var d = make([]byte, conLen)
	buf := bytes.NewBuffer(d)
	_, errR := c.Read(buf.Bytes())
	if errR != nil {
		c.Close()
		panic(errR)
	}
	cmdBodyBytes := buf.Bytes()
	var resBeforeSend = ResponseBody{}
	errUnmarshal := json.Unmarshal(cmdBodyBytes, &resBeforeSend)
	if errUnmarshal != nil {
		panic(errUnmarshal)
	}
	return resBeforeSend
}

// =========== data transfer ===========
// HandleLS send fileList to client
func HandleLS(c net.Conn, cmd *ResponseBody) {
	if cmd.Cmd == "ls" {
		var files []string
		var path = servPath
		if cmd.FilePath != "/" && cmd.FileName != "" {
			path = filepath.Join(cmd.FilePath, cmd.FileName)
		} else {
			path = filepath.Join(servPath, cmd.FileName)
		}
		files = ListFiles(cmd, path)
		fileListToWrite := []byte(fmt.Sprintf("Files are: \n%v\n", files))
		len := len(fileListToWrite)
		cmd.FileSize = int64(len)
		cmd.Status = 200
		WriteHeadLen(c, *cmd)
		WriteHeadBody(c, *cmd)
		fmt.Println("len: ", len)
		c.Write([]byte(fmt.Sprintf("Files are: \n%v\n", files)))
	}
}

// HandleGET tell client file size and then send file content
func HandleGET(c net.Conn, res *ResponseBody) {
	//ReadHeadLen(c)
	response := ReadHead(c)
	var filePath string
	if response.Cmd == "get" && response.FileName != "" {
		if response.FilePath == "/" {
			filePath = filepath.Join(servPath, response.FileName)
			response.FilePath = servPath
		} else {
			filePath = filepath.Join(response.FilePath, response.FileName)
		}
	} else {
		filePath = ""
		response.Status = 404
		WriteHead(c, response)
		//WriteHeadLen(c, response)
		//WriteHeadBody(c, response)
		c.Close()
		return
	}
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		panic(err)
	}
	fileLen := fileInfo.Size()
	response.FileSize = fileLen
	response.Status = 200
	WriteHead(c, response)
	//WriteHeadLen(c, response)
	//WriteHeadBody(c, response)
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	var buf = make([]byte, fileLen)
	reader := bufio.NewReader(f)
	reader.Read(buf)
	c.Write(buf)
	c.Close()

}

// =========== tools ===========
// ListFiles list files in a dir recursively
func ListFiles(cmd *ResponseBody, path string) []string {
	var files []string
	if cmd.Cmd == "ls" {
		var err error
		_, files, err = walkReturnDirSlice(path)
		if err != nil {
			panic(err)
		}
	} else {
		return files
	}
	return files
}

// return walked dirs slice and walked files slice
func walkReturnDirSlice(dir string) ([]string, []string, error) {
	var walkedFiles, walkedDirs []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if isDir(path) == 1 {
			walkedDirs = append(walkedDirs, path)
		} else if isDir(path) == 0 {
			walkedFiles = append(walkedFiles, path)
		}
		return nil
	})

	if err != nil {
		return []string{}, []string{}, err
	}

	return walkedDirs, walkedFiles, nil
}

// dir return 1, file return 0, not exist return -1
func isDir(file string) int {
	f, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			return -1
		}
	}
	if f.IsDir() {
		return 1
	}
	return 0
}
