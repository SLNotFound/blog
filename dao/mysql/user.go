package mysql

import "fmt"

// QueryUserByUsername 根据username查询用户是否存在
func QueryUserByUsername(username string) {
	sqlStr := fmt.Sprintf("select * from user where username=%s", username)
	db.QueryRow(sqlStr)
}
