package wdk

import (
	"sync"
)

// OperatorVo 结构体
type OperatorVo struct {
	// 操作人id
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 操作人name
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// 操作人type
	OperatorType int64 `json:"operator_type,omitempty" xml:"operator_type,omitempty"`
}

var poolOperatorVo = sync.Pool{
	New: func() any {
		return new(OperatorVo)
	},
}

// GetOperatorVo() 从对象池中获取OperatorVo
func GetOperatorVo() *OperatorVo {
	return poolOperatorVo.Get().(*OperatorVo)
}

// ReleaseOperatorVo 释放OperatorVo
func ReleaseOperatorVo(v *OperatorVo) {
	v.OperatorId = ""
	v.OperatorName = ""
	v.OperatorType = 0
	poolOperatorVo.Put(v)
}
