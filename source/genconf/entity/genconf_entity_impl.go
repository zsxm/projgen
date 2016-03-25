//scgen
package entity

import (
	"bytes"
	"fmt"

	"github.com/zsxm/scgo/data"
)

//----------------------GenconfBean begin--------------------------------------

type GenconfBean struct {
	bean  *Genconf
	beans *Genconfs
}

func NewGenconfBean() *GenconfBean {
	e := &GenconfBean{}
	return e
}

func (this *GenconfBean) NewEntity() data.EntityInterface {
	return NewGenconf()
}

func (this *GenconfBean) NewEntitys(cap int) data.EntitysInterface {
	return NewGenconfs(cap)
}

func (this *GenconfBean) Entity() data.EntityInterface {
	if this.bean == nil {
		return nil
	}
	return this.bean
}

func (this *GenconfBean) Entitys() data.EntitysInterface {
	if this.beans == nil {
		return nil
	}
	return this.beans
}

func (this *GenconfBean) Datas() *Genconfs {
	if this.beans == nil {
		return nil
	}
	return this.beans
}

func (this *GenconfBean) Table() data.TableInformation {
	return genconfTableInformation
}

func (this *GenconfBean) FieldNames() data.FieldNames {
	return genconfFieldNames
}

func (this *GenconfBean) SetEntity(bean data.EntityInterface) {
	this.bean = bean.(*Genconf)
}

func (this *GenconfBean) SetEntitys(beans data.EntitysInterface) {
	this.beans = beans.(*Genconfs)
}

//------------------------------------------------------------

//------------------------------------------------------------
type Genconfs struct {
	datas []data.EntityInterface
	page  *data.Page
}

func NewGenconfs(cap int) *Genconfs {
	e := &Genconfs{}
	e.datas = make([]data.EntityInterface, 0, cap)
	return e
}

func (this *Genconfs) SetPage(page *data.Page) {
	this.page = page
}

func (this *Genconfs) Add(e data.EntityInterface) {
	this.datas = append(this.datas, e.(*Genconf))
}

func (this *Genconfs) Values() []data.EntityInterface {
	return this.datas
}

func (this *Genconfs) Len() int {
	return len(this.datas)
}

func (this *Genconfs) Get(i int) *Genconf {
	return this.datas[i].(*Genconf)
}

func (this *Genconfs) Table() data.TableInformation {
	return genconfTableInformation
}

func (this *Genconfs) FieldNames() data.FieldNames {
	return genconfFieldNames
}

func (this *Genconfs) JSON() string {
	var wr bytes.Buffer
	wr.WriteString("[")
	for i, v := range this.Values() {
		if i > 0 {
			wr.WriteString(",")
		}
		wr.WriteString(v.JSON())
	}
	wr.WriteString("]")
	return wr.String()
}

//----------------------GenconfBean end--------------------------------------

//----------------------Genconf begin--------------------------------------
func NewGenconf() *Genconf {
	return &Genconf{}
}

func (this *Genconf) Gid() *data.String {
	return &this.gid
}

func (this *Genconf) SetGid(value string) {
	this.gid.SetValue(value)
}

func (this *Genconf) SetValue(filedName, value string) {
	this.Field(filedName).SetValue(value)
}

func (this *Genconf) Table() data.TableInformation {
	return genconfTableInformation
}

func (this *Genconf) FieldNames() data.FieldNames {
	return genconfFieldNames
}

func (this *Genconf) Field(filedName string) data.EntityField {
	switch filedName {
	case "gid":
		this.gid.SetPrimaryKey(true)
		return this.gid.StructType()
	}
	return nil
}

func (this *Genconf) JSON() string {
	var b bytes.Buffer
	b.WriteString("{")
	b.WriteString(fmt.Sprintf(`"gid":%q`, this.gid.Value()))
	b.WriteString("}")
	return b.String()
}

//----------------------NewGenconf end--------------------------------------

//----------------------init() end--------------------------------------
func init() {
	genconfTableInformation.SetTableName("genconf")
	genconfColumnsArr := []string{
		"gid",
	}
	genconfTableInformation.SetColumns(genconfColumnsArr)
	genconfFieldNamesArr := []string{
		"gid",
	}
	genconfFieldNames.SetNames(genconfFieldNamesArr)
}

var genconfTableInformation data.TableInformation
var genconfFieldNames data.FieldNames

//----------------------init() end--------------------------------------
