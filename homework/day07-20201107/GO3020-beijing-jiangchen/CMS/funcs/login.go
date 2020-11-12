package funcs

import (
	"CMS/consts"
	"CMS/models"
	"crypto/md5"
	"errors"
	"fmt"
	"strings"
)

//Login ...
// Login function. the default username & password is admin/admin.
// Get input from stdin and check whether the input equals the default password,
// force exit if type the wrong password in 3 consecutive times, or start the program normally if not so
func Login() (e error) {

	CMSBanner()
	line := models.NewLiner()
	defer line.Close()

	models.MUE.PersistentStorage.GetNonAdmin()

	fmt.Printf("\nWelcome to CMS System.\nPlease input login username.\n")
	var UserName string
	for {
		username, err := line.State.Prompt("username : ")
		if err != nil {
			e = err
			return
		}
		username = strings.TrimSpace(username)
		if username == "" {
			continue
		} else if !models.MUE.CheckUserName(username) {
			fmt.Println("user did not exists...Abort...")
			e = errors.New("user did not exists")
			return
		} else {
			UserName = username
			break
		}
	}
	user := models.MUE.QueryElementName(UserName)
	for {
		password, err := line.State.PasswordPrompt("please input password: ")
		if err != nil {
			e = err
			return
		}
		if md5.Sum([]byte(password)) == user[0].Password {
			fmt.Printf("\nlogin success.\n")
			break
		} else {
			fmt.Println("password incorrect...")
			consts.LoginCount++
		}
		if consts.LoginCount == 3 {
			fmt.Println("Bye :(")
			e = errors.New("login failure time reach 3")
			break
		}
	}
	return
}
