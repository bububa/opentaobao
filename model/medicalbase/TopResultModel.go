package medicalbase

import (
	"sync"
)

// TopResultModel 结构体
type TopResultModel struct {
	// 操作说明
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 操作码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// isvId
	Model string `json:"model,omitempty" xml:"model,omitempty"`
}

var poolTopResultModel = sync.Pool{
	New: func() any {
		return new(TopResultModel)
	},
}

// GetTopResultModel() 从对象池中获取TopResultModel
func GetTopResultModel() *TopResultModel {
	return poolTopResultModel.Get().(*TopResultModel)
}

// ReleaseTopResultModel 释放TopResultModel
func ReleaseTopResultModel(v *TopResultModel) {
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Model = ""
	poolTopResultModel.Put(v)
}
