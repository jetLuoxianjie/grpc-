package impl

import (
	"context"
	"fmt"
	"log"
	"xxm/user/dao"
	"xxm/user/proto"
)

type UserServer struct {
}

func (us *UserServer)CreateUser(ctx context.Context, in *proto.ReqUser) (res *proto.ResUser,err error){
  res=new(proto.ResUser)
  //数据库操作
  fmt.Println(in.Username,in.Pwd,"-----impl")

  err=dao.CreateUser(in.Username,in.Pwd)
  if err!=nil{
	  log.Println("CreateUser: %v", err)
	  return
  }
  res.Username=in.Username
  res.Username=in.Pwd

  return

}

func (us *UserServer)Login(ctx context.Context, in *proto.ReqUser) (res *proto.ResUser,err error){
	res=new(proto.ResUser)
	//需要查询数据库
	fmt.Println(in.Username,in.Pwd,"------impllogin")
    username,err:=  dao.Login(in.Username,in.Pwd)
 // err还需处理
    res.Username=username

	if err!=nil{
		fmt.Println(username,"uername ",res.Username)
		log.Println("login: %v", err)
		return  res,err
	}


	return res,nil


}


