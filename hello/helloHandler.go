package hello

import(
	"fmt"
	"github.com/labstack/echo"
)

func helloHandler(c echo.Context) error {
	return  c.String(http.StatusOK, "hello world!")
}
