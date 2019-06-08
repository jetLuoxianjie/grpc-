package handlers

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
	"html/template"
	"log"
	"net/http"
	"xxm/web/proto"
)

type HomePage struct {
	Name string
}

type UserPage struct {
	Name string
}
type Message struct {
	Name string
}
//连接另一个服务器

func connect()(conn *grpc.ClientConn,c proto.HelloServiceClient){

	//// 发起链接
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	//
	//// 创建pb包的客户端

	c = proto.NewHelloServiceClient(conn)

	return
}
func connect1()(conn *grpc.ClientConn,c proto.SpiderServiceClient){

	//// 发起链接
	conn, err := grpc.Dial("127.0.0.1:9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	//
	//// 创建pb包的客户端

	c = proto.NewSpiderServiceClient(conn)

	return
}


func HomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		p:=&HomePage{Name: ""}
		t, e:=template.ParseFiles("./templates/home.html")
		if e!=nil {
			log.Printf("Parsing template home.html error: %s", e)
			return
		}
		t.Execute(w, p)
		return
}

func LoginHandler(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {



	conn,c:=connect()
	defer conn.Close()

	request.ParseForm()
	username:=request.Form.Get("username")
	pwd:=request.Form.Get("pwd")
	fmt.Println(username,pwd,"  ------loginhandler")

	req:=proto.ReqUser{Username:username,Pwd:pwd}
	res,err:=c.Login(context.Background(),&req)

	fmt.Println(	res,"   --------res")


	if err!=nil{
		log.Println("登录失败")
		p:=&Message{Name: "用户名或密码错误"}
		t, e:=template.ParseFiles("./templates/home.html")
		if e!=nil {
			log.Printf("Parsing template home.html error: %s", e)
			return
		}


		t.Execute(writer, p)


	}else{
      p:=&UserPage{Name:username}
		t, e:=template.ParseFiles("./templates/userhome.html")
		if e!=nil{
			log.Printf("Parsing userhome.html error: %s", e)
			return
		}
		t.Execute(writer, p)

	}


}

func RegisterHandler(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {

	conn,c:=connect()
	defer conn.Close()
	//conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//defer conn.Close()
	//c:=proto.NewHelloServiceClient(conn)

	request.ParseForm()
	username:=request.Form.Get("username")
	pwd:=request.Form.Get("pwd")

	req:=proto.ReqUser{}
	req.Username=username
	req.Pwd=pwd

	r2,err:=c.CreateUser(context.Background(),&req)

	if err!=nil{
		fmt.Println("添加失败")
		fmt.Println(r2.Username)
		log.Printf("%s", err)
	}

	fmt.Println(r2)

	writer.Write([]byte("注册成功"))


}


func GetHandler(writer http.ResponseWriter, request *http.Request, ps httprouter.Params){
	request.ParseForm()
	start:=request.Form.Get("start")
	end:=request.Form.Get("end")
	//传开始页
   req:= proto.Req{Start:start,End:end}
	conn,c:=connect1()
	defer conn.Close()
     r2,err:=c.GetXiaohua(context.Background(),&req)

	if err!=nil {
		fmt.Println(err)
	}

     if r2.Message=="success"{
     	writer.Write([]byte("爬取成功"))
	 }else{
		 writer.Write([]byte("爬取失败"))
	 }




}