package mysql

import (
	"blog/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

const secret = "test"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)

// CheckUserExist 根据username查询用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

func InsertUser(u *models.User) (err error) {
	// 对密码进行加密
	u.Password = encryptPassword(u.Password)
	sqlStr := `insert into user(user_id, username, password) values (?, ?, ?)`
	_, err = db.Exec(sqlStr, u.UserID, u.Username, u.Password)

	return
}

// encryptPassword 加密密码
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(u *models.User) (err error) {
	oPassword := u.Password
	sqlStr := `select user_id, username, password from user where username=?`
	err = db.Get(u, sqlStr, u.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	password := encryptPassword(oPassword)
	if password != u.Password {
		return ErrorInvalidPassword
	}
	return
}
