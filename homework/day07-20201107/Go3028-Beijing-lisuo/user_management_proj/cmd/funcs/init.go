package funcs

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day07-20201107/Go3028-Beijing-lisuo/user_management_proj/cmd/db"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day07-20201107/Go3028-Beijing-lisuo/user_management_proj/define"
)

// NewUser make a new user contains user's info
func NewUser(id int64, name, cell, address, born, passwd string) define.User {
	return define.User{
		ID:      id,
		Name:    name,
		Cell:    cell,
		Address: address,
		Born: func() time.Time {
			t, _ := time.Parse("2006.01.02", born)
			return t
		}(),
		Passwd: fmt.Sprintf("%x", md5.Sum([]byte(passwd))),
	}
}

// Init add some users to define.UserList
func Init(ul *[]define.User) {
	user0 := NewUser(0, "admin", "18811992299", "HaidianDistrict,BeijingXinParkRestaurants,BeixiaguanSubdistrict,HaidianDistrict,China",
		time.Now().Format("2006.01.02"), "qwert")
	(*ul) = append((*ul), user0)
	AddFunc()
	db.ReadUsers()
}
