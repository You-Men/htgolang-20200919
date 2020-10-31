package funcs

import (
	"errors"
	"fmt"
	"os"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day06-20201031/Go3028-Beijing-lisuo/utils"
	"github.com/olekukonko/tablewriter"
)

// CmdToFuncMap mapping the cmd to corresponding func
var CmdToFuncMap = map[string]func(){}

// FuncList contains the funcs
var FuncList = map[string]string{
	"add":   "AddUser",
	"del":   "DelUser",
	"mod":   "ModifyUser",
	"get":   "QueryUser",
	"show":  "ShowCurrentUserList",
	"help":  "ShowHelp",
	"h":     "ShowHelp",
	"cls":   "utils.ClearScreen",
	"q":     "utils.Quit",
	"Q":     "utils.Quit",
	"quit":  "utils.Quit",
	"exit":  "utils.Quit",
	"mycmd": "DoMap",
}

// CmdList contains default commands binding to funcs
var CmdList = []string{}

// CmdToFunc mapping cmd to corresponding funcs
func CmdToFunc(cmd string, f func()) {
	if _, ok := CmdToFuncMap[cmd]; ok {
		panic(fmt.Sprintf("command %s already exists", cmd))
	}
	CmdToFuncMap[cmd] = f
	return
}

// DoMap mapping a unused cmd to a exists func
func DoMap() {
	f := &FuncList
	fmt.Print("Input the cmd you want use?\n> ")
	cmd := utils.Read()
	if cmd == "" {
		fmt.Print("You input a blank string, default func will bind to utils.ShowHelp\n> ")
		CmdToFunc(cmd, ShowHelp)
		(*f)[cmd] = "ShowHelp"
	}
	fmt.Print("Chose the func you want map with?\n> ")
	ShowFuncList()
	c := utils.Read()
	switch c {
	case "1":
		CmdToFunc(cmd, AddUser)
		(*f)[cmd] = "AddUser"
	case "2":
		CmdToFunc(cmd, DelUser)
		(*f)[cmd] = "DelUser"
	case "3":
		CmdToFunc(cmd, ModifyUser)
		(*f)[cmd] = "ModifyUser"
	case "4":
		CmdToFunc(cmd, QueryUser)
		(*f)[cmd] = "QueryUser"
	case "5":
		CmdToFunc(cmd, ShowCurrentUserList)
		(*f)[cmd] = "ShowCurrentUserList"
	case "6":
		CmdToFunc(cmd, ShowHelp)
		(*f)[cmd] = "ShowHelp"
	case "7":
		CmdToFunc(cmd, utils.ClearScreen)
		(*f)[cmd] = "utils.ClearScreen"
	case "8":
		CmdToFunc(cmd, utils.Quit)
		(*f)[cmd] = "utils.Quit"
	}

}

// ShowFuncList display the func and what they do
func ShowFuncList() {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetAutoFormatHeaders(false)
	t.SetAutoWrapText(false)
	t.SetReflowDuringAutoWrap(false)

	t.SetHeader([]string{"Nu", "Function", "What they do", "Current binding cmds"})
	t.Append([]string{"1", "AddUser", "Add a User", utils.GetKeyByValue(FuncList, "AddUser")})
	t.Append([]string{"2", "DelUser", "Delete a User", utils.GetKeyByValue(FuncList, "DelUser")})
	t.Append([]string{"3", "ModifyUser", "Modify a User", utils.GetKeyByValue(FuncList, "ModifyUser")})
	t.Append([]string{"4", "QueryUser", "Search User", utils.GetKeyByValue(FuncList, "QueryUser")})
	t.Append([]string{"5", "ShowCurrentUserList", "Show User List", utils.GetKeyByValue(FuncList, "ShowCurrentUserList")})
	t.Append([]string{"6", "ShowHelp", "Show help list", utils.GetKeyByValue(FuncList, "ShowHelp")})
	t.Append([]string{"7", "utils.ClearScreen", "Clean the terminal", utils.GetKeyByValue(FuncList, "utils.ClearScreen")})
	t.Append([]string{"8", "utils.Quit", "Exit this program", utils.GetKeyByValue(FuncList, "utils.Quit")})
	t.Render()
}

// AddFunc register cmd to func
func AddFunc() {
	CmdToFunc("add", AddUser)
	CmdToFunc("del", DelUser)
	CmdToFunc("mod", ModifyUser)
	CmdToFunc("get", QueryUser)
	CmdToFunc("show", ShowCurrentUserList)
	CmdToFunc("help", ShowHelp)
	CmdToFunc("h", ShowHelp)
	CmdToFunc("cls", utils.ClearScreen)
	CmdToFunc("q", utils.Quit)
	CmdToFunc("Q", utils.Quit)
	CmdToFunc("quit", utils.Quit)
	CmdToFunc("exit", utils.Quit)
	CmdToFunc("mycmd", DoMap)
}

// ExecFunc execute the func of a cmd
func ExecFunc(input string) error {
	if f, ok := CmdToFuncMap[input]; ok {
		f()
	} else {
		return errors.New("[wrong cmd type \"h\" for help]\n> ")
	}
	return nil
}
