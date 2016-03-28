//scgen
package action

import (
	"github.com/zsxm/scgo/chttp"
	"projgen/source/fields/entity"
)

func init() {
	chttp.Action("/fields/index", index).Get()
}

//gen
func index(c chttp.Context) {
	e := entity.NewFields()
	c.BindData(e)
	c.JSON(e.JSON(), true)
}
