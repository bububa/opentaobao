package iot

import (
	"sync"
)

// BusinessRecipeStepActionOpenParam 结构体
type BusinessRecipeStepActionOpenParam struct {
	// 指令名称
	ActionName string `json:"action_name,omitempty" xml:"action_name,omitempty"`
	// 指令值
	ActionValue string `json:"action_value,omitempty" xml:"action_value,omitempty"`
	// 指令顺序号
	Sequence int64 `json:"sequence,omitempty" xml:"sequence,omitempty"`
}

var poolBusinessRecipeStepActionOpenParam = sync.Pool{
	New: func() any {
		return new(BusinessRecipeStepActionOpenParam)
	},
}

// GetBusinessRecipeStepActionOpenParam() 从对象池中获取BusinessRecipeStepActionOpenParam
func GetBusinessRecipeStepActionOpenParam() *BusinessRecipeStepActionOpenParam {
	return poolBusinessRecipeStepActionOpenParam.Get().(*BusinessRecipeStepActionOpenParam)
}

// ReleaseBusinessRecipeStepActionOpenParam 释放BusinessRecipeStepActionOpenParam
func ReleaseBusinessRecipeStepActionOpenParam(v *BusinessRecipeStepActionOpenParam) {
	v.ActionName = ""
	v.ActionValue = ""
	v.Sequence = 0
	poolBusinessRecipeStepActionOpenParam.Put(v)
}
