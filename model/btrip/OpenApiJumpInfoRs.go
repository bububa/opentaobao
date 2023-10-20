package btrip

import (
	"sync"
)

// OpenApiJumpInfoRs 结构体
type OpenApiJumpInfoRs struct {
	// 跳转url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}

var poolOpenApiJumpInfoRs = sync.Pool{
	New: func() any {
		return new(OpenApiJumpInfoRs)
	},
}

// GetOpenApiJumpInfoRs() 从对象池中获取OpenApiJumpInfoRs
func GetOpenApiJumpInfoRs() *OpenApiJumpInfoRs {
	return poolOpenApiJumpInfoRs.Get().(*OpenApiJumpInfoRs)
}

// ReleaseOpenApiJumpInfoRs 释放OpenApiJumpInfoRs
func ReleaseOpenApiJumpInfoRs(v *OpenApiJumpInfoRs) {
	v.Url = ""
	poolOpenApiJumpInfoRs.Put(v)
}
