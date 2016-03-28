//scgen
package action

import (
	"github.com/zsxm/scgo/chttp"
	"projgen/source/project/entity"
)

func init() {
	chttp.Action("/project/index", index).Get()
}

//gen
func index(c chttp.Context) {
	e := entity.NewProject()
	c.BindData(e)
	c.JSON(e.JSON(), true)
}
