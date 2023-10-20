package ascp

import (
	"sync"
)

// ModifyEstConTimeRequest 结构体
type ModifyEstConTimeRequest struct {
	// 修改后的发货时效，只允许往前改;1_+绝对时间为修改到当日的23:59:59，2_+天数为修改到付款时间加上天数
	EstConTime string `json:"est_con_time,omitempty" xml:"est_con_time,omitempty"`
	// 子单维度的商品id，若是组合商品需要传成分品的商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 交易子单编号
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
}

var poolModifyEstConTimeRequest = sync.Pool{
	New: func() any {
		return new(ModifyEstConTimeRequest)
	},
}

// GetModifyEstConTimeRequest() 从对象池中获取ModifyEstConTimeRequest
func GetModifyEstConTimeRequest() *ModifyEstConTimeRequest {
	return poolModifyEstConTimeRequest.Get().(*ModifyEstConTimeRequest)
}

// ReleaseModifyEstConTimeRequest 释放ModifyEstConTimeRequest
func ReleaseModifyEstConTimeRequest(v *ModifyEstConTimeRequest) {
	v.EstConTime = ""
	v.ItemId = 0
	v.SubOrderId = 0
	poolModifyEstConTimeRequest.Put(v)
}
