package funcs

import (
	"fmt"
	"strconv"
	"strings"

	define "github.com/htgolang/htgolang-20200919/tree/master/homework/day03-20201008/Go3028-Beijing-lisuo/define"
	utils "github.com/htgolang/htgolang-20200919/tree/master/homework/day03-20201008/Go3028-Beijing-lisuo/utils"
)

// one can use Id or Name to find a user and del it
// if input a Id, find the user use this Id
// if input a Name, find the user use this Name
// if no this user, prompt error
func ModifyUser() {
	var name string
	var input string
	fmt.Print("Who you want to modify(Id/Name)?\n> ")
	fmt.Scanln(&input)
	if s, err := strconv.Atoi(strings.TrimSpace(input)); err == nil {
		id := int64(s)
		fmt.Printf("idType: %T  idValue: %v\n", id, id)
		user := utils.IdFindUser(define.UserList, id)
		if user.Name == "" {
			fmt.Println("No such user.", user)
		} else {
			fmt.Printf("Find user %v --> %v, are you sure to modify %v?(y/n) ", name, user, name)
			fmt.Scanln(&input)
			if strings.ToLower(input) == "y" {
				fmt.Println("modifying...........")
				utils.IdModUser(&define.UserList, id)
			} else if strings.ToLower(input) == "n" {
				fmt.Println("abort modify...........")
			}
		}
	} else {
		name = strings.ToLower(strings.TrimSpace(input))
		fmt.Printf("nameType: %T  nameValue: %v\n", name, name)
		user := utils.NameFindUser(name)
		if user == nil {
			fmt.Println("No such user.", user)
		} else {
			fmt.Printf("Find user %v --> %v, are you sure to delete %v?(y/n) ", name, user, name)
			fmt.Scanln(&input)
			if strings.ToLower(input) == "y" {
				fmt.Println("modifying...........")
				utils.NameModUser(&define.UserList, name)
			} else if strings.ToLower(input) == "n" {
				fmt.Println("abort modify...........")
			}
		}
	}
}
