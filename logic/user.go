package logic

import (
	"blog/dao/mysql"
	"blog/pkg/snowflake"
)

func SignUp() {
	// 判断用户是否存在，查询语句
	mysql.QueryUserByUsername("")

	// 生成UID
	snowflake.GenID()

}
