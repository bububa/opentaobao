package alihealthcrm

import (
	"sync"
)

// AlibabaAlihealthPregnancyProductGetModel 结构体
type AlibabaAlihealthPregnancyProductGetModel struct {
	// 文章列表
	Contents []Contents `json:"contents,omitempty" xml:"contents>contents,omitempty"`
	// 总页数
	TotalPage string `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// 总条数
	TotalCount string `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolAlibabaAlihealthPregnancyProductGetModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthPregnancyProductGetModel)
	},
}

// GetAlibabaAlihealthPregnancyProductGetModel() 从对象池中获取AlibabaAlihealthPregnancyProductGetModel
func GetAlibabaAlihealthPregnancyProductGetModel() *AlibabaAlihealthPregnancyProductGetModel {
	return poolAlibabaAlihealthPregnancyProductGetModel.Get().(*AlibabaAlihealthPregnancyProductGetModel)
}

// ReleaseAlibabaAlihealthPregnancyProductGetModel 释放AlibabaAlihealthPregnancyProductGetModel
func ReleaseAlibabaAlihealthPregnancyProductGetModel(v *AlibabaAlihealthPregnancyProductGetModel) {
	v.Contents = v.Contents[:0]
	v.TotalPage = ""
	v.TotalCount = ""
	poolAlibabaAlihealthPregnancyProductGetModel.Put(v)
}
