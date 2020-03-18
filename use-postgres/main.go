package main

import (
	"encoding/json"
	"go-learn/use-postgres/db"
	"log"
)

type UserInfo struct {
	Id       int             `gorm:"column:id;default:"`
	Username string          `gorm:"column:username;default:"`
	Age      int             `gorm:"column:age;default:"`
	Realname string          `gorm:"column:realname;default:"`
	Raw      json.RawMessage `gorm:"column:raw;default:"`
}

func (ui UserInfo) TableName() string {
	return "user_info"
}

func main() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
	var user UserInfo
	var users = make([]UserInfo, 0, 50)

	// 删
	defer func() {
		// 删
		if e := db.DB.Model(&user).Where("username=?", "pig1").Delete(&UserInfo{}).Error; e != nil {
			log.Println(e.Error())
			return
		}
	}()

	// 增
	if e := db.DB.Model(&user).Create(&UserInfo{
		Username: "pig1",
		Age:      1,
		Realname: "猪",
		Raw: json.RawMessage([]byte(`
        {
            "info": "这是一条附加信息"
        }
    `)),
	}).Error; e != nil {
		log.Println(e.Error())
		return
	}
	log.Println("插入成功")
	// 查
	if e := db.DB.Model(&user).Where("username=?", "pig1").First(&user).Error; e != nil {
		log.Println(e.Error())
		return
	}
	log.Println("查询单个成功", user)

	if e := db.DB.Model(&user).Where("id>?", "0").First(&users).Error; e != nil {
		log.Println(e.Error())
		return
	}
	log.Println("查询列表成功", users)

	// 改
	if e := db.DB.Model(&user).Where("username=?", "pig1").Updates(UserInfo{
		Age: 1900,
	}).Error; e != nil {
		log.Println(e.Error())
		return
	}
	log.Println("修改成功")
}
