package entity

import (
	"github.com/zsxm/scgo/data"
)

//go:generate $GOPATH/src/github.com/zsxm/scgo/tools/scgen/scgen.exe -fileDir=$GOFILE -projectDir=projgen -moduleName=user -goSource=source
//go:@Table value=genconf
type User struct {

	//go:@Column value=uid
	//go:@Identif
	uid data.String

	//go:@Column value=name
	name data.String

	//go:@Column value=pass
	pass data.String
}
