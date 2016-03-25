//scgen
package action

import (
	"projgen/source/genconf/entity"

	"github.com/zsxm/scgo/chttp"
)

func init() {
	chttp.Action("/genconf/index", index).Get()
}

//gen
func index(c chttp.Context) {
	entity.NewGenconfBean()
	c.HTML("/genconf/genconf", nil)
}
