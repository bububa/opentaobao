package alihealthcrm

import (
	"sync"
)

// AlibabaAlihealthPregnancyNavigateinfoQueryModel 结构体
type AlibabaAlihealthPregnancyNavigateinfoQueryModel struct {
	// list
	List []Content `json:"list,omitempty" xml:"list>content,omitempty"`
}

var poolAlibabaAlihealthPregnancyNavigateinfoQueryModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthPregnancyNavigateinfoQueryModel)
	},
}

// GetAlibabaAlihealthPregnancyNavigateinfoQueryModel() 从对象池中获取AlibabaAlihealthPregnancyNavigateinfoQueryModel
func GetAlibabaAlihealthPregnancyNavigateinfoQueryModel() *AlibabaAlihealthPregnancyNavigateinfoQueryModel {
	return poolAlibabaAlihealthPregnancyNavigateinfoQueryModel.Get().(*AlibabaAlihealthPregnancyNavigateinfoQueryModel)
}

// ReleaseAlibabaAlihealthPregnancyNavigateinfoQueryModel 释放AlibabaAlihealthPregnancyNavigateinfoQueryModel
func ReleaseAlibabaAlihealthPregnancyNavigateinfoQueryModel(v *AlibabaAlihealthPregnancyNavigateinfoQueryModel) {
	v.List = v.List[:0]
	poolAlibabaAlihealthPregnancyNavigateinfoQueryModel.Put(v)
}
