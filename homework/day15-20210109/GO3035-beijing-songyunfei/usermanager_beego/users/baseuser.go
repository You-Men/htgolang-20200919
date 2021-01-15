package users

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
	"time"
)

type Userinfo struct {
	Id int64 		`orm:"column(id);pk;auto"`
	Name string		`orm:"name"`
	Sex string		`orm:"sex"`
	Addr string		`orm:"addr"`
	Tel string		`orm:"tel"`
	Birthday  time.Time `orm:"type(date)"`
	Passwd string	`orm:"password"`
	CreatedAt *time.Time `orm:"auto_now_add"`
	UpdatedAt *time.Time `orm:"auto_now"`
	DeletedAt time.Time `orm:"type(datetime);null"`
}

type UserRole struct {
	Id int `orm:"column(id);pk;auto"`
	UserId int64
	RoleId int `orm:"not null"`
}

type RolePermission struct {
	Id int `orm:"column(id);pk;auto"`
	RoleId int
	PermissionId int
}
type Permission struct {
	Id int `orm:"column(id);pk;auto"`
	PName string `orm:"PName"`
	Path string
}
type RoleInfo struct {
	Id int `orm:"column(id);pk;auto"`
	RoleName string
}

//log
type FileInfo struct {
	Id int `orm:"column(id);pk;auto"`
	FileName string
	Size int64
	UploadAt *time.Time `orm:"auto_now_add"`
}

type Recorder struct {
	Id int64 `orm:"column(id);pk;auto"`
	FileId int64
	IpAddr string
	Method string
	Status int
	InsertAt *time.Time `orm:"auto_now_add"`
}

//通过ID查找用户
func (u *Userinfo) FindByid(id int64) (userinfo *Userinfo,err error) {
	quser := &Userinfo{Id:id}
	qs := orm.NewOrm()
	if err := qs.Read(quser,"id"); err != nil{
		return &Userinfo{}, fmt.Errorf("未找到")
	}
	return quser,nil
}

// 新增用户
func (u *Userinfo) Add(name,addr,sex,tel,passwd,birthday,roleid string  )  error {
	rid,err := strconv.Atoi(roleid)
	if err != nil {
		return fmt.Errorf("%s 非法ID.\n",roleid)
	}
	user := &Userinfo{Name:name}
	qs := orm.NewOrm()
	if err := qs.Read(user,"Name"); err == nil {
		return fmt.Errorf("%s 已存在.\n",name)
	}
	hashed, _ := bcrypt.GenerateFromPassword([]byte(passwd), 10)
	br,err := time.Parse("2006-01-02",birthday)
	if err != nil{
		return err
	}
	iuser := &Userinfo{
		Name:      name,
		Sex:       sex,
		Addr:      addr,
		Tel:       tel,
		Birthday:  br,
		Passwd:    string(hashed),
	}
	if uid, err := qs.Insert(iuser); err != nil {
		return err
	}else {
		ur := &UserRole{
			UserId: uid,
			RoleId: rid,
		}
		_,err := qs.Insert(ur)
		if err != nil {
			return err
		}
	}

	return nil
}

//通过ID删除用户
func (u *Userinfo) Del(id int64 ) error {
	_,err := u.FindByid(id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	qs := orm.NewOrm()
	duser := &Userinfo{Id:id}
	if err = qs.Read(duser);err == nil {
		//now := time.Now().Format("2006-01-02 15:04:05")
		//duser.DeletedAt, _ = time.Parse("2006-01-02 15:04:05",now)
		duser.DeletedAt = time.Now()
		if _,err := qs.Update(duser);err != nil {
			return err
		}
	}else {
		return err
	}
	//if _,err := qs.Delete(&Userinfo{Id:id}); err != nil{
	//	return err
	//}
	return nil
}
// 修改索引为index的用户信息
func (u *Userinfo) Modify(id int64, name,addr,sex,tel,bri,passwd string ) error {
	qs := orm.NewOrm()
	var tmu Userinfo
	if _,err := qs.QueryTable(&Userinfo{}).Exclude("Id__exact",id).Filter("name__exact",name).All(&tmu);err != nil {
		return err
	}
	if tmu.Id != 0 {
		return fmt.Errorf("用户名:%s已存在.\n",name)
	}
	nuser := &Userinfo{Id:id}
	br,err := time.Parse("2006-01-02",bri)
	if err != nil{
		return err
	}
	uqs := orm.NewOrm()
	hashed, _ := bcrypt.GenerateFromPassword([]byte(passwd), 10)
	if err := uqs.Read(nuser); err == nil {
		nuser.Name = name
		nuser.Addr = addr
		nuser.Sex = sex
		nuser.Tel = tel
		nuser.Birthday = br
		nuser.Passwd = string(hashed)
	}
	if _,err := uqs.Update(nuser);err != nil {
		return err
	}
	return nil
}

//通过关键字查找用户 返回userinfo结构体和error
func (u *Userinfo) QueryUser(str string) (f Userinfo, ok bool) {
	for _,v := range u.Getall() {
		if strings.Contains(v.Name,str) || strings.Contains(v.Addr,str)|| strings.Contains(v.Tel,str) {
			return v,true
		}
	}
	return Userinfo{}, false
}

//认证功能
func (u *Userinfo) Auth(username, passwd string) bool {
	auser :=  &Userinfo{Name:username}
	auser.Name = username
	qs := orm.NewOrm()
	if err := qs.Read(auser,"Name"); err == nil {
		if err := bcrypt.CompareHashAndPassword([]byte(auser.Passwd),[]byte(passwd));err == nil {
			return true
		}
		return false
	}
	return false
}
//Get all
func (u *Userinfo) Getall() []Userinfo  {
	var userslice  []Userinfo
	ormer := orm.NewOrm()
	qs := ormer.QueryTable(&Userinfo{})
	_,err := qs.Filter("DeletedAt__isnull",true).All(&userslice)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return userslice
}

//Get role
func (u *Userinfo) GetRole() (roleslice []RoleInfo,err error) {
	ormer := orm.NewOrm()
	qs := ormer.QueryTable(&RoleInfo{})
	_,err = qs.All(&roleslice)
	if err != nil {
		return nil ,err
	}
	return roleslice,nil
}

func (u *Userinfo) InitDB(dbtype,dsn string) error {
	//orm.Debug = true
	if dbtype == "mysql"{
		if err := orm.RegisterDriver("mysql",orm.DRMySQL); err !=nil {
			return err
		}
	}

	if err := orm.RegisterDataBase("default","mysql",dsn); err != nil {
		return err
	}
	orm.RegisterModel(&Userinfo{},&Permission{},&RoleInfo{},&UserRole{},&RolePermission{},&FileInfo{},&Recorder{})
	if err := orm.RunSyncdb("default", false, true); err != nil {
		return err
	}
	qs := orm.NewOrm()
	if count, err := qs.QueryTable(&Userinfo{}).Count(); err != nil {
		return err
	}else {
		if count == 0 {
			if err := u.Add("admin","pek","0","110","admin","2020-10-10","1"); err != nil {
				return err
			}
		}
		return nil
	}
}
