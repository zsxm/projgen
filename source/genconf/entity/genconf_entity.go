package entity

import (
	"github.com/zsxm/scgo/data"
)

//go:generate $GOPATH/src/github.com/zsxm/scgo/tools/scgen/scgen.exe -fileDir=$GOFILE -projectDir=projgen -moduleName=genconf -goSource=source
//go:@Table value=genconf
type Genconf struct {

	//go:@Column value=gid
	//go:@Identif
	gid data.String
}
