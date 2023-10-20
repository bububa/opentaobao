package maitix

import (
	"sync"
)

// OpenApiPostFeeParam 结构体
type OpenApiPostFeeParam struct {
	// 地址国标
	Address *AddressDto `json:"address,omitempty" xml:"address,omitempty"`
	// 项目ID-必填
	ProjectId int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
}

var poolOpenApiPostFeeParam = sync.Pool{
	New: func() any {
		return new(OpenApiPostFeeParam)
	},
}

// GetOpenApiPostFeeParam() 从对象池中获取OpenApiPostFeeParam
func GetOpenApiPostFeeParam() *OpenApiPostFeeParam {
	return poolOpenApiPostFeeParam.Get().(*OpenApiPostFeeParam)
}

// ReleaseOpenApiPostFeeParam 释放OpenApiPostFeeParam
func ReleaseOpenApiPostFeeParam(v *OpenApiPostFeeParam) {
	v.Address = nil
	v.ProjectId = 0
	poolOpenApiPostFeeParam.Put(v)
}
