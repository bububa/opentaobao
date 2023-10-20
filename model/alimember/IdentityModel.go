package alimember

import (
	"sync"
)

// IdentityModel 结构体
type IdentityModel struct {
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
}

var poolIdentityModel = sync.Pool{
	New: func() any {
		return new(IdentityModel)
	},
}

// GetIdentityModel() 从对象池中获取IdentityModel
func GetIdentityModel() *IdentityModel {
	return poolIdentityModel.Get().(*IdentityModel)
}

// ReleaseIdentityModel 释放IdentityModel
func ReleaseIdentityModel(v *IdentityModel) {
	v.StartTime = ""
	v.EndTime = ""
	poolIdentityModel.Put(v)
}
