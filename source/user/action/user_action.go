//scgen
package action

import (
	"projgen/source/user/entity"
	"projgen/source/user/log"

	"github.com/zsxm/scgo/chttp"
)

func init() {
	chttp.Action("/user/login", login).Post()
}

//gen
func index(c chttp.Context) {
	e := entity.NewUser()
	c.BindData(e)
	c.JSON(e.JSON(), true)
}

func login(c chttp.Context) {
	log.Println("login")

	c.Redirect("/genconf/index")
}
