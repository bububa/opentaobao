package btrip

import (
	"sync"
)

// OpenApiNewApplyRs 结构体
type OpenApiNewApplyRs struct {
	// 用户传入审批单id
	ThirdpartApplyId string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	// 商旅审批单id
	ApplyId string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
}

var poolOpenApiNewApplyRs = sync.Pool{
	New: func() any {
		return new(OpenApiNewApplyRs)
	},
}

// GetOpenApiNewApplyRs() 从对象池中获取OpenApiNewApplyRs
func GetOpenApiNewApplyRs() *OpenApiNewApplyRs {
	return poolOpenApiNewApplyRs.Get().(*OpenApiNewApplyRs)
}

// ReleaseOpenApiNewApplyRs 释放OpenApiNewApplyRs
func ReleaseOpenApiNewApplyRs(v *OpenApiNewApplyRs) {
	v.ThirdpartApplyId = ""
	v.ApplyId = ""
	poolOpenApiNewApplyRs.Put(v)
}
