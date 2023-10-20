package tblogistics

import (
	"sync"
)

// ConsignStatusRequest 结构体
type ConsignStatusRequest struct {
	// 子订单id（组合品不需要传系统会自动计算）
	SubTid string `json:"sub_tid,omitempty" xml:"sub_tid,omitempty"`
	// 子订单是否部分发货，true：部分发货；false：全部发货；周期购、分销订单不支持部分发货
	IsPartConsign bool `json:"is_part_consign,omitempty" xml:"is_part_consign,omitempty"`
}

var poolConsignStatusRequest = sync.Pool{
	New: func() any {
		return new(ConsignStatusRequest)
	},
}

// GetConsignStatusRequest() 从对象池中获取ConsignStatusRequest
func GetConsignStatusRequest() *ConsignStatusRequest {
	return poolConsignStatusRequest.Get().(*ConsignStatusRequest)
}

// ReleaseConsignStatusRequest 释放ConsignStatusRequest
func ReleaseConsignStatusRequest(v *ConsignStatusRequest) {
	v.SubTid = ""
	v.IsPartConsign = false
	poolConsignStatusRequest.Put(v)
}
