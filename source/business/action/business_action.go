//scgen
package action

import (
	"github.com/zsxm/scgo/chttp"
	"projgen/source/business/entity"
)

func init() {
	chttp.Action("/business/index", index).Get()
}

//gen
func index(c chttp.Context) {
	e := entity.NewBusiness()
	c.BindData(e)
	c.JSON(e.JSON(), true)
}
