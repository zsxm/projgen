package entity

import (
	"github.com/zsxm/scgo/data"
)

//go:generate $GOPATH/src/github.com/zsxm/scgo/tools/scgen/scgen.exe -fileDir=$GOFILE -projectDir=projgen -moduleName=business -goSource=source
//go:@Table value=project
type Business struct {

	//go:@Column value=id
	//go:@Identif
	id data.String

	//go:@Column value=name
	name data.String

	//go:@Column value=project_id
	projectId data.String

	//go:@Column value=deleted
	deleted data.Integer
}
