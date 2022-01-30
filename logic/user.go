package logic

import (
	"blog/dao/mysql"
	"blog/models"
	"blog/pkg/snowflake"
)

func SignUp(p *models.SignUpParam) {
	// 判断用户是否存在，查询语句
	mysql.QueryUserByUsername("")

	// 生成UID
	snowflake.GenID()

}
