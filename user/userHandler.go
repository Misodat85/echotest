package user

import(
	"net/http"
	"github.com/labstack/echo/hello"
)
type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

func user(c echo.Context) (err error) {
	u := new(User)
	if err = c.Bind(u); err != nil {
	  return
	}
	return c.JSON(http.StatusOK, u)//JSON sends a JSON response with status code.
}
