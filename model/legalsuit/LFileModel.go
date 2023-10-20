package legalsuit

import (
	"sync"
)

// LFileModel 结构体
type LFileModel struct {
	// 附件名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 附件key
	Key string `json:"key,omitempty" xml:"key,omitempty"`
}

var poolLFileModel = sync.Pool{
	New: func() any {
		return new(LFileModel)
	},
}

// GetLFileModel() 从对象池中获取LFileModel
func GetLFileModel() *LFileModel {
	return poolLFileModel.Get().(*LFileModel)
}

// ReleaseLFileModel 释放LFileModel
func ReleaseLFileModel(v *LFileModel) {
	v.Name = ""
	v.Key = ""
	poolLFileModel.Put(v)
}
