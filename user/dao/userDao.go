package dao

import (
	"fmt"
	"log"
)

func CreateUser(loginName string, pwd string) error {

	fmt.Println(loginName,pwd,"----dao")
	stmtIns, err := dbConn.Prepare("INSERT INTO users(login_name, pwd) values(?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil


}
func  Login(loginName string, pwd string) (string,error){
	fmt.Println(loginName,pwd,"------daologin")
	stmtOut, err := dbConn.Prepare("SELECT login_name FROM users WHERE login_name = ? and pwd = ?")
	if err!=nil {
		log.Printf("%s", err)
		return "用户或密码错误",err
	}
	name:=""
	fmt.Println("到这里没有报错")

	err = stmtOut.QueryRow(loginName,pwd).Scan(&name)

	fmt.Println(err)
	if err != nil {
		return "用户或密码错误", err
	}

	fmt.Println("没有有错了")
	defer stmtOut.Close()
	return name, nil





}