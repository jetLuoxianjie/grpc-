package main


import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"xxm/web/handlers"
)

func RegisterHandler() *httprouter.Router {
	router:=httprouter.New()
	router.GET("/",handlers.HomeHandler)
	router.POST("/", handlers.HomeHandler)
	router.POST("/login",handlers.LoginHandler )
	router.POST("/register",handlers.RegisterHandler)
	router.POST("/get",handlers.GetHandler)

	router.ServeFiles("/statics/*filepath", http.Dir("./template"))
	return router
}

func main(){
      r:=RegisterHandler()
      http.ListenAndServe(":8080",r)

}
