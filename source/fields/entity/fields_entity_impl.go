//scgen
package entity

import (
	"bytes"
	"fmt"
	"github.com/zsxm/scgo/data"
	"strconv"
)

//----------------------FieldsBean begin--------------------------------------

type FieldsBean struct {
	bean  *Fields
	beans *Fieldss
}

func NewFieldsBean() *FieldsBean {
	e := &FieldsBean{}
	return e
}

func (this *FieldsBean) NewEntity() data.EntityInterface {
	return NewFields()
}

func (this *FieldsBean) NewEntitys(cap int) data.EntitysInterface {
	return NewFieldss(cap)
}

func (this *FieldsBean) Entity() data.EntityInterface {
	if this.bean == nil {
		return nil
	}
	return this.bean
}

func (this *FieldsBean) Entitys() data.EntitysInterface {
	if this.beans == nil {
		return nil
	}
	return this.beans
}

func (this *FieldsBean) Datas() *Fieldss {
	if this.beans == nil {
		return nil
	}
	return this.beans
}

func (this *FieldsBean) Table() data.TableInformation {
	return fieldsTableInformation
}

func (this *FieldsBean) FieldNames() data.FieldNames {
	return fieldsFieldNames
}

func (this *FieldsBean) SetEntity(bean data.EntityInterface) {
	this.bean = bean.(*Fields)
}

func (this *FieldsBean) SetEntitys(beans data.EntitysInterface) {
	this.beans = beans.(*Fieldss)
}

//------------------------------------------------------------

//------------------------------------------------------------
type Fieldss struct {
	datas []data.EntityInterface
	page  *data.Page
}

func NewFieldss(cap int) *Fieldss {
	e := &Fieldss{}
	e.datas = make([]data.EntityInterface, 0, cap)
	return e
}

func (this *Fieldss) SetPage(page *data.Page) {
	this.page = page
}

func (this *Fieldss) Add(e data.EntityInterface) {
	this.datas = append(this.datas, e.(*Fields))
}

func (this *Fieldss) Values() []data.EntityInterface {
	return this.datas
}

func (this *Fieldss) Len() int {
	return len(this.datas)
}

func (this *Fieldss) Get(i int) *Fields {
	return this.datas[i].(*Fields)
}

func (this *Fieldss) Table() data.TableInformation {
	return fieldsTableInformation
}

func (this *Fieldss) FieldNames() data.FieldNames {
	return fieldsFieldNames
}

func (this *Fieldss) JSON() string {
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

//----------------------FieldsBean end--------------------------------------

//----------------------Fields begin--------------------------------------
func NewFields() *Fields {
	return &Fields{}
}

func (this *Fields) Id() *data.String {
	return &this.id
}

func (this *Fields) SetId(value string) {
	this.id.SetValue(value)
}

func (this *Fields) Name() *data.String {
	return &this.name
}

func (this *Fields) SetName(value string) {
	this.name.SetValue(value)
}

func (this *Fields) BusinessId() *data.String {
	return &this.businessId
}

func (this *Fields) SetBusinessId(value string) {
	this.businessId.SetValue(value)
}

func (this *Fields) Deleted() *data.Integer {
	return &this.deleted
}

func (this *Fields) SetDeleted(value int) {
	this.deleted.SetValue(strconv.Itoa(value))
}

func (this *Fields) SetValue(filedName, value string) {
	this.Field(filedName).SetValue(value)
}

func (this *Fields) Table() data.TableInformation {
	return fieldsTableInformation
}

func (this *Fields) FieldNames() data.FieldNames {
	return fieldsFieldNames
}

func (this *Fields) Field(filedName string) data.EntityField {
	switch filedName {
	case "id":
		this.id.SetPrimaryKey(true)
		return this.id.StructType()
	case "name":
		return this.name.StructType()
	case "businessid", "business_id":
		return this.businessId.StructType()
	case "deleted":
		return this.deleted.StructType()
	}
	return nil
}

func (this *Fields) JSON() string {
	var b bytes.Buffer
	b.WriteString("{")
	b.WriteString(fmt.Sprintf(`"id":%q`, this.id.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"name":%q`, this.name.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"businessId":%q`, this.businessId.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"deleted":%q`, this.deleted.Value()))
	b.WriteString("}")
	return b.String()
}

//----------------------NewFields end--------------------------------------

//----------------------init() end--------------------------------------
func init() {
	fieldsTableInformation.SetTableName("project")
	fieldsColumnsArr := []string{
		"id", "name", "business_id", "deleted", 
	}
	fieldsTableInformation.SetColumns(fieldsColumnsArr)
	fieldsFieldNamesArr := []string{
		"id", "name", "businessId", "deleted", 
	}
	fieldsFieldNames.SetNames(fieldsFieldNamesArr)
}

var fieldsTableInformation data.TableInformation
var fieldsFieldNames data.FieldNames

//----------------------init() end--------------------------------------
