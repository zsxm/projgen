//scgen
package entity

import (
	"bytes"
	"fmt"
	"github.com/zsxm/scgo/data"
	"strconv"
)

//----------------------ProjectBean begin--------------------------------------

type ProjectBean struct {
	bean  *Project
	beans *Projects
}

func NewProjectBean() *ProjectBean {
	e := &ProjectBean{}
	return e
}

func (this *ProjectBean) NewEntity() data.EntityInterface {
	return NewProject()
}

func (this *ProjectBean) NewEntitys(cap int) data.EntitysInterface {
	return NewProjects(cap)
}

func (this *ProjectBean) Entity() data.EntityInterface {
	if this.bean == nil {
		return nil
	}
	return this.bean
}

func (this *ProjectBean) Entitys() data.EntitysInterface {
	if this.beans == nil {
		return nil
	}
	return this.beans
}

func (this *ProjectBean) Datas() *Projects {
	if this.beans == nil {
		return nil
	}
	return this.beans
}

func (this *ProjectBean) Table() data.TableInformation {
	return projectTableInformation
}

func (this *ProjectBean) FieldNames() data.FieldNames {
	return projectFieldNames
}

func (this *ProjectBean) SetEntity(bean data.EntityInterface) {
	this.bean = bean.(*Project)
}

func (this *ProjectBean) SetEntitys(beans data.EntitysInterface) {
	this.beans = beans.(*Projects)
}

//------------------------------------------------------------

//------------------------------------------------------------
type Projects struct {
	datas []data.EntityInterface
	page  *data.Page
}

func NewProjects(cap int) *Projects {
	e := &Projects{}
	e.datas = make([]data.EntityInterface, 0, cap)
	return e
}

func (this *Projects) SetPage(page *data.Page) {
	this.page = page
}

func (this *Projects) Add(e data.EntityInterface) {
	this.datas = append(this.datas, e.(*Project))
}

func (this *Projects) Values() []data.EntityInterface {
	return this.datas
}

func (this *Projects) Len() int {
	return len(this.datas)
}

func (this *Projects) Get(i int) *Project {
	return this.datas[i].(*Project)
}

func (this *Projects) Table() data.TableInformation {
	return projectTableInformation
}

func (this *Projects) FieldNames() data.FieldNames {
	return projectFieldNames
}

func (this *Projects) JSON() string {
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

//----------------------ProjectBean end--------------------------------------

//----------------------Project begin--------------------------------------
func NewProject() *Project {
	return &Project{}
}

func (this *Project) Id() *data.String {
	return &this.id
}

func (this *Project) SetId(value string) {
	this.id.SetValue(value)
}

func (this *Project) Name() *data.String {
	return &this.name
}

func (this *Project) SetName(value string) {
	this.name.SetValue(value)
}

func (this *Project) Deleted() *data.Integer {
	return &this.deleted
}

func (this *Project) SetDeleted(value int) {
	this.deleted.SetValue(strconv.Itoa(value))
}

func (this *Project) SetValue(filedName, value string) {
	this.Field(filedName).SetValue(value)
}

func (this *Project) Table() data.TableInformation {
	return projectTableInformation
}

func (this *Project) FieldNames() data.FieldNames {
	return projectFieldNames
}

func (this *Project) Field(filedName string) data.EntityField {
	switch filedName {
	case "id":
		this.id.SetPrimaryKey(true)
		return this.id.StructType()
	case "name":
		return this.name.StructType()
	case "deleted":
		return this.deleted.StructType()
	}
	return nil
}

func (this *Project) JSON() string {
	var b bytes.Buffer
	b.WriteString("{")
	b.WriteString(fmt.Sprintf(`"id":%q`, this.id.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"name":%q`, this.name.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"deleted":%q`, this.deleted.Value()))
	b.WriteString("}")
	return b.String()
}

//----------------------NewProject end--------------------------------------

//----------------------init() end--------------------------------------
func init() {
	projectTableInformation.SetTableName("project")
	projectColumnsArr := []string{
		"id", "name", "deleted", 
	}
	projectTableInformation.SetColumns(projectColumnsArr)
	projectFieldNamesArr := []string{
		"id", "name", "deleted", 
	}
	projectFieldNames.SetNames(projectFieldNamesArr)
}

var projectTableInformation data.TableInformation
var projectFieldNames data.FieldNames

//----------------------init() end--------------------------------------
