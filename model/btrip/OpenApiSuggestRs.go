package btrip

import (
	"sync"
)

// OpenApiSuggestRs 结构体
type OpenApiSuggestRs struct {
	// 城市列表
	Cities []CityVo `json:"cities,omitempty" xml:"cities>city_vo,omitempty"`
}

var poolOpenApiSuggestRs = sync.Pool{
	New: func() any {
		return new(OpenApiSuggestRs)
	},
}

// GetOpenApiSuggestRs() 从对象池中获取OpenApiSuggestRs
func GetOpenApiSuggestRs() *OpenApiSuggestRs {
	return poolOpenApiSuggestRs.Get().(*OpenApiSuggestRs)
}

// ReleaseOpenApiSuggestRs 释放OpenApiSuggestRs
func ReleaseOpenApiSuggestRs(v *OpenApiSuggestRs) {
	v.Cities = v.Cities[:0]
	poolOpenApiSuggestRs.Put(v)
}
