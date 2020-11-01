package funcs

import (
	"crypto/md5"
	"fmt"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day06-20201031/Go3028-Beijing-lisuo/define"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day06-20201031/Go3028-Beijing-lisuo/utils"
)

// Login implement login func
func Login() bool {
	var logCount int8 = 0
	var logged bool = true
	for {
		fmt.Print("Input the UserName[admin]: \n> ")
		input := utils.Read()
		user, err := NameFindUser(&define.UserList, input)
		if err != nil {
			fmt.Println("No such user.")
			logCount++
			if logCount == 3 {
				logged = false
				break
			}
		} else {
			fmt.Print("Input the PassWord[qwert]: \n> ")
			input := utils.Read()
			inputPasswd := md5.Sum([]byte(fmt.Sprintf("%x", md5.Sum([]byte(input)))))
			if user.Passwd == inputPasswd {
				fmt.Printf("User %s logged in.\n", user.Name)
				return logged
			}
			fmt.Println("Wrong password...")
			logCount++
			if logCount == 3 {
				logged = false
				break
			}
			continue
		}
	}
	return logged
}
