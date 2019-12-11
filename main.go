package main

import(
	"net/http"
	"github.com/labstack/echo"
	"github.com/Misodat85/echotest"
)

func main(){
	//Creating New instance
	e := echo.New()
	//Routing
	e.GET("/",helloHandler)
	e.POST("/user",user)
	//Start Server
	e.Logger.Fatal(e.Start(":1323"))
}
  