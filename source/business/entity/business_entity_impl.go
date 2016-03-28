//scgen
package entity

import (
	"bytes"
	"fmt"
	"github.com/zsxm/scgo/data"
	"strconv"
)

//----------------------BusinessBean begin--------------------------------------

type BusinessBean struct {
	bean  *Business
	beans *Businesss
}

func NewBusinessBean() *BusinessBean {
	e := &BusinessBean{}
	return e
}

func (this *BusinessBean) NewEntity() data.EntityInterface {
	return NewBusiness()
}

func (this *BusinessBean) NewEntitys(cap int) data.EntitysInterface {
	return NewBusinesss(cap)
}

func (this *BusinessBean) Entity() data.EntityInterface {
	if this.bean == nil {
		return nil
	}
	return this.bean
}

func (this *BusinessBean) Entitys() data.EntitysInterface {
	if this.beans == nil {
		return nil
	}
	return this.beans
}

func (this *BusinessBean) Datas() *Businesss {
	if this.beans == nil {
		return nil
	}
	return this.beans
}

func (this *BusinessBean) Table() data.TableInformation {
	return businessTableInformation
}

func (this *BusinessBean) FieldNames() data.FieldNames {
	return businessFieldNames
}

func (this *BusinessBean) SetEntity(bean data.EntityInterface) {
	this.bean = bean.(*Business)
}

func (this *BusinessBean) SetEntitys(beans data.EntitysInterface) {
	this.beans = beans.(*Businesss)
}

//------------------------------------------------------------

//------------------------------------------------------------
type Businesss struct {
	datas []data.EntityInterface
	page  *data.Page
}

func NewBusinesss(cap int) *Businesss {
	e := &Businesss{}
	e.datas = make([]data.EntityInterface, 0, cap)
	return e
}

func (this *Businesss) SetPage(page *data.Page) {
	this.page = page
}

func (this *Businesss) Add(e data.EntityInterface) {
	this.datas = append(this.datas, e.(*Business))
}

func (this *Businesss) Values() []data.EntityInterface {
	return this.datas
}

func (this *Businesss) Len() int {
	return len(this.datas)
}

func (this *Businesss) Get(i int) *Business {
	return this.datas[i].(*Business)
}

func (this *Businesss) Table() data.TableInformation {
	return businessTableInformation
}

func (this *Businesss) FieldNames() data.FieldNames {
	return businessFieldNames
}

func (this *Businesss) JSON() string {
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

//----------------------BusinessBean end--------------------------------------

//----------------------Business begin--------------------------------------
func NewBusiness() *Business {
	return &Business{}
}

func (this *Business) Id() *data.String {
	return &this.id
}

func (this *Business) SetId(value string) {
	this.id.SetValue(value)
}

func (this *Business) Name() *data.String {
	return &this.name
}

func (this *Business) SetName(value string) {
	this.name.SetValue(value)
}

func (this *Business) ProjectId() *data.String {
	return &this.projectId
}

func (this *Business) SetProjectId(value string) {
	this.projectId.SetValue(value)
}

func (this *Business) Deleted() *data.Integer {
	return &this.deleted
}

func (this *Business) SetDeleted(value int) {
	this.deleted.SetValue(strconv.Itoa(value))
}

func (this *Business) SetValue(filedName, value string) {
	this.Field(filedName).SetValue(value)
}

func (this *Business) Table() data.TableInformation {
	return businessTableInformation
}

func (this *Business) FieldNames() data.FieldNames {
	return businessFieldNames
}

func (this *Business) Field(filedName string) data.EntityField {
	switch filedName {
	case "id":
		this.id.SetPrimaryKey(true)
		return this.id.StructType()
	case "name":
		return this.name.StructType()
	case "projectid", "project_id":
		return this.projectId.StructType()
	case "deleted":
		return this.deleted.StructType()
	}
	return nil
}

func (this *Business) JSON() string {
	var b bytes.Buffer
	b.WriteString("{")
	b.WriteString(fmt.Sprintf(`"id":%q`, this.id.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"name":%q`, this.name.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"projectId":%q`, this.projectId.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"deleted":%q`, this.deleted.Value()))
	b.WriteString("}")
	return b.String()
}

//----------------------NewBusiness end--------------------------------------

//----------------------init() end--------------------------------------
func init() {
	businessTableInformation.SetTableName("project")
	businessColumnsArr := []string{
		"id", "name", "project_id", "deleted", 
	}
	businessTableInformation.SetColumns(businessColumnsArr)
	businessFieldNamesArr := []string{
		"id", "name", "projectId", "deleted", 
	}
	businessFieldNames.SetNames(businessFieldNamesArr)
}

var businessTableInformation data.TableInformation
var businessFieldNames data.FieldNames

//----------------------init() end--------------------------------------
