package wdk

import (
	"sync"
)

// EleBillRequest 结构体
type EleBillRequest struct {
	// 查询页码,默认查询第一页,默认每页 20 条
	Page string `json:"page,omitempty" xml:"page,omitempty"`
	// 查询日期，时间戳格式(2019-06-10=1560124800)
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 渠道店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
}

var poolEleBillRequest = sync.Pool{
	New: func() any {
		return new(EleBillRequest)
	},
}

// GetEleBillRequest() 从对象池中获取EleBillRequest
func GetEleBillRequest() *EleBillRequest {
	return poolEleBillRequest.Get().(*EleBillRequest)
}

// ReleaseEleBillRequest 释放EleBillRequest
func ReleaseEleBillRequest(v *EleBillRequest) {
	v.Page = ""
	v.Date = ""
	v.ShopId = ""
	poolEleBillRequest.Put(v)
}
