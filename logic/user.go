package logic

import (
	"blog/dao/mysql"
	"blog/models"
	"blog/pkg/jwt"
	"blog/pkg/snowflake"
)

func SignUp(p *models.SignUpParam) (err error) {
	// 判断用户是否存在，查询语句
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 生成UID
	userID := snowflake.GenID()
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.InsertUser(user)
}

func Login(p *models.LoginParam) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	//return mysql.Login(user)
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	return jwt.GenToken(user.UserID, user.Username)
}
