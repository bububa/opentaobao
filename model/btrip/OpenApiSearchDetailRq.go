package btrip

import (
	"sync"
)

// OpenApiSearchDetailRq 结构体
type OpenApiSearchDetailRq struct {
	// 第三方企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 第三方用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 1、老版本2、isv对外版本
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}

var poolOpenApiSearchDetailRq = sync.Pool{
	New: func() any {
		return new(OpenApiSearchDetailRq)
	},
}

// GetOpenApiSearchDetailRq() 从对象池中获取OpenApiSearchDetailRq
func GetOpenApiSearchDetailRq() *OpenApiSearchDetailRq {
	return poolOpenApiSearchDetailRq.Get().(*OpenApiSearchDetailRq)
}

// ReleaseOpenApiSearchDetailRq 释放OpenApiSearchDetailRq
func ReleaseOpenApiSearchDetailRq(v *OpenApiSearchDetailRq) {
	v.CorpId = ""
	v.UserId = ""
	v.OrderId = 0
	v.Version = 0
	poolOpenApiSearchDetailRq.Put(v)
}
