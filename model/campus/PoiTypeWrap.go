package campus

import (
	"sync"
)

// PoiTypeWrap 结构体
type PoiTypeWrap struct {
	// description
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// fullName
	FullName string `json:"full_name,omitempty" xml:"full_name,omitempty"`
	// 分类名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 分类编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 大类名称
	SecondTypeName string `json:"second_type_name,omitempty" xml:"second_type_name,omitempty"`
	// 大类编码
	SecondTypeCode string `json:"second_type_code,omitempty" xml:"second_type_code,omitempty"`
	// modifier
	Modifier string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// creator
	Creator string `json:"creator,omitempty" xml:"creator,omitempty"`
	// gmtModified
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// gmtCreate
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 类别的实际名称,空间或分组
	Classify string `json:"classify,omitempty" xml:"classify,omitempty"`
	// 大类ID
	SecondTypeId int64 `json:"second_type_id,omitempty" xml:"second_type_id,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// isDelete
	IsDelete bool `json:"is_delete,omitempty" xml:"is_delete,omitempty"`
}

var poolPoiTypeWrap = sync.Pool{
	New: func() any {
		return new(PoiTypeWrap)
	},
}

// GetPoiTypeWrap() 从对象池中获取PoiTypeWrap
func GetPoiTypeWrap() *PoiTypeWrap {
	return poolPoiTypeWrap.Get().(*PoiTypeWrap)
}

// ReleasePoiTypeWrap 释放PoiTypeWrap
func ReleasePoiTypeWrap(v *PoiTypeWrap) {
	v.Description = ""
	v.FullName = ""
	v.Name = ""
	v.Code = ""
	v.SecondTypeName = ""
	v.SecondTypeCode = ""
	v.Modifier = ""
	v.Creator = ""
	v.GmtModified = ""
	v.GmtCreate = ""
	v.Classify = ""
	v.SecondTypeId = 0
	v.Id = 0
	v.IsDelete = false
	poolPoiTypeWrap.Put(v)
}
