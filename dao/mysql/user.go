package mysql

import "fmt"

func QueryUserIsExist(username string) {
	sqlStr := fmt.Sprintf("select * from user where username=%s", username)
	db.QueryRow(sqlStr)
}
