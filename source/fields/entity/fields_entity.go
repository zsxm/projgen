package entity

import (
	"github.com/zsxm/scgo/data"
)

//go:generate $GOPATH/src/github.com/zsxm/scgo/tools/scgen/scgen.exe -fileDir=$GOFILE -projectDir=projgen -moduleName=fields -goSource=source
//go:@Table value=project
type Fields struct {

	//go:@Column value=id
	//go:@Identif
	id data.String

	//go:@Column value=name
	name data.String

	//go:@Column value=business_id
	businessId data.String

	//go:@Column value=deleted
	deleted data.Integer
}
