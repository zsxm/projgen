//scgen
package entity

import (
	"bytes"
	"fmt"

	"github.com/zsxm/scgo/data"
)

//----------------------UserBean begin--------------------------------------

type UserBean struct {
	bean  *User
	beans *Users
}

func NewUserBean() *UserBean {
	e := &UserBean{}
	return e
}

func (this *UserBean) NewEntity() data.EntityInterface {
	return NewUser()
}

func (this *UserBean) NewEntitys(cap int) data.EntitysInterface {
	return NewUsers(cap)
}

func (this *UserBean) Entity() data.EntityInterface {
	if this.bean == nil {
		return nil
	}
	return this.bean
}

func (this *UserBean) Entitys() data.EntitysInterface {
	if this.beans == nil {
		return nil
	}
	return this.beans
}

func (this *UserBean) Datas() *Users {
	if this.beans == nil {
		return nil
	}
	return this.beans
}

func (this *UserBean) Table() data.TableInformation {
	return userTableInformation
}

func (this *UserBean) FieldNames() data.FieldNames {
	return userFieldNames
}

func (this *UserBean) SetEntity(bean data.EntityInterface) {
	this.bean = bean.(*User)
}

func (this *UserBean) SetEntitys(beans data.EntitysInterface) {
	this.beans = beans.(*Users)
}

//------------------------------------------------------------

//------------------------------------------------------------
type Users struct {
	datas []data.EntityInterface
	page  *data.Page
}

func NewUsers(cap int) *Users {
	e := &Users{}
	e.datas = make([]data.EntityInterface, 0, cap)
	return e
}

func (this *Users) SetPage(page *data.Page) {
	this.page = page
}

func (this *Users) Add(e data.EntityInterface) {
	this.datas = append(this.datas, e.(*User))
}

func (this *Users) Values() []data.EntityInterface {
	return this.datas
}

func (this *Users) Len() int {
	return len(this.datas)
}

func (this *Users) Get(i int) *User {
	return this.datas[i].(*User)
}

func (this *Users) Table() data.TableInformation {
	return userTableInformation
}

func (this *Users) FieldNames() data.FieldNames {
	return userFieldNames
}

func (this *Users) JSON() string {
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

//----------------------UserBean end--------------------------------------

//----------------------User begin--------------------------------------
func NewUser() *User {
	return &User{}
}

func (this *User) Uid() *data.String {
	return &this.uid
}

func (this *User) SetUid(value string) {
	this.uid.SetValue(value)
}

func (this *User) Name() *data.String {
	return &this.name
}

func (this *User) SetName(value string) {
	this.name.SetValue(value)
}

func (this *User) Pass() *data.String {
	return &this.pass
}

func (this *User) SetPass(value string) {
	this.pass.SetValue(value)
}

func (this *User) SetValue(filedName, value string) {
	this.Field(filedName).SetValue(value)
}

func (this *User) Table() data.TableInformation {
	return userTableInformation
}

func (this *User) FieldNames() data.FieldNames {
	return userFieldNames
}

func (this *User) Field(filedName string) data.EntityField {
	switch filedName {
	case "uid":
		this.uid.SetPrimaryKey(true)
		return this.uid.StructType()
	case "name":
		return this.name.StructType()
	case "pass":
		return this.pass.StructType()
	}
	return nil
}

func (this *User) JSON() string {
	var b bytes.Buffer
	b.WriteString("{")
	b.WriteString(fmt.Sprintf(`"uid":%q`, this.uid.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"name":%q`, this.name.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"pass":%q`, this.pass.Value()))
	b.WriteString("}")
	return b.String()
}

//----------------------NewUser end--------------------------------------

//----------------------init() end--------------------------------------
func init() {
	userTableInformation.SetTableName("genconf")
	userColumnsArr := []string{
		"uid", "name", "pass",
	}
	userTableInformation.SetColumns(userColumnsArr)
	userFieldNamesArr := []string{
		"uid", "name", "pass",
	}
	userFieldNames.SetNames(userFieldNamesArr)
}

var userTableInformation data.TableInformation
var userFieldNames data.FieldNames

//----------------------init() end--------------------------------------
