/* ----------------------------------
*  @author suyame 2022-07-06 16:06:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package main

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"regexp"
	"time"
)

type User struct {
	ID            uint           `gorm:"primaryKey" json:"id,omitempty"`
	CreatedAt     time.Time      `json:"created_at,omitempty"`
	UpdatedAt     time.Time      `json:"created_at,omitempty"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Name          string         `json:"name,omitempty"`
	Password      string         `json:"password,omitempty"`
	FollowCount   int64          `json:"follow_count,omitempty"`
	FollowerCount int64          `json:"follower_count,omitempty"`
	FollowID      string         `json:"follow_id,omitempty"`
	FollowerID    string         `json:"follower_id,omitempty"`
	IsFollow      bool           `json:"is_follow,omitempty"`
	LikeVideosID  string         `json:"like_videos,omitempty"`
}

var db *gorm.DB

func init() {
	db, _ = ConnectDataBase()
	db.AutoMigrate(&User{})
}

// 连接数据库
func ConnectDataBase() (db *gorm.DB, err error) {
	// 记得在容器内修改
	dsn := "root:19990221@tcp(mysql:3306)/golang_mysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}

/***************************** 用户 ********************************/
// 检测用户是否存在
func checkUserName(username string) (User, bool) {
	user := User{}
	result := db.Select("ID").Where("name = ?", username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return User{}, false
	}
	return user, true
}

// 判断账号是否为邮箱格式

func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` // 匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// 获得全部用户的简要信息，用于初始化
func GetUsersBriefInfo() []User {
	users := []User{}
	db.Select("ID", "Name", "Password", "LikeVideosID", "FollowCount", "FollowerCount", "FollowID", "FollowerID").Find(&users)
	return users
}

// FindUserInfo 根据名字查找
func FindUserInfo(username string) (User, bool) {
	user := User{}
	result := db.Where("name = ? ", username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return User{}, false
	}
	return user, true
}

// 通过user_id查找用户
func FindUserByID(id uint) (User, bool) {
	user := User{}
	result := db.Select("ID", "Name", "FollowCount", "FollowerCount", "IsFollow", "FollowID", "FollowerID").Where("ID = ?", id).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return User{}, false
	}
	return user, true
}

// 更新用户
func UpdateUser(user User) {
	// 通过发送消息来更新
	db.Model(&user).Updates(user)
}

// 添加用户信息
func AddUserInfo(username string, password string) (User, error) {
	// 向数据库中插入一条数据
	newUser := User{
		Name:     username,
		Password: password,
	}
	result := db.Create(&newUser)
	if result.Error != nil {
		return User{}, result.Error
	}
	return newUser, nil
}

// 删除用户
func DeleteUser(username string) error {
	user := User{}
	result := db.Where("name = ? ", username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	db.Select(clause.Associations).Delete(&user)
	return nil
}
