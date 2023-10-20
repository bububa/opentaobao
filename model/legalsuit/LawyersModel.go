package legalsuit

import (
	"sync"
)

// LawyersModel 结构体
type LawyersModel struct {
	// 律师列表
	Lawyers []Lawyers `json:"lawyers,omitempty" xml:"lawyers>lawyers,omitempty"`
	// 操作类型
	OperationType string `json:"operation_type,omitempty" xml:"operation_type,omitempty"`
}

var poolLawyersModel = sync.Pool{
	New: func() any {
		return new(LawyersModel)
	},
}

// GetLawyersModel() 从对象池中获取LawyersModel
func GetLawyersModel() *LawyersModel {
	return poolLawyersModel.Get().(*LawyersModel)
}

// ReleaseLawyersModel 释放LawyersModel
func ReleaseLawyersModel(v *LawyersModel) {
	v.Lawyers = v.Lawyers[:0]
	v.OperationType = ""
	poolLawyersModel.Put(v)
}
