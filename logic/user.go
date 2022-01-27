package logic

import (
	"blog/dao/mysql"
	"fmt"
)

func SignUp() {
	mysql.QueryUserIsExist("")

	fmt.Println("hello")
}
