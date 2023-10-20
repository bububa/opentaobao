package ascp

import (
	"sync"
)

// FuturePlanItemResult 结构体
type FuturePlanItemResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 外部交易号(主)
	ExtOrderId string `json:"ext_order_id,omitempty" xml:"ext_order_id,omitempty"`
	// 外部交易号(子)
	ExtSubOrderId string `json:"ext_sub_order_id,omitempty" xml:"ext_sub_order_id,omitempty"`
	// 仓code
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 期货库存id
	ChannelInvId string `json:"channel_inv_id,omitempty" xml:"channel_inv_id,omitempty"`
	// 货主id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 货品
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 业务结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolFuturePlanItemResult = sync.Pool{
	New: func() any {
		return new(FuturePlanItemResult)
	},
}

// GetFuturePlanItemResult() 从对象池中获取FuturePlanItemResult
func GetFuturePlanItemResult() *FuturePlanItemResult {
	return poolFuturePlanItemResult.Get().(*FuturePlanItemResult)
}

// ReleaseFuturePlanItemResult 释放FuturePlanItemResult
func ReleaseFuturePlanItemResult(v *FuturePlanItemResult) {
	v.ErrorCode = ""
	v.ErrorMessage = ""
	v.ExtOrderId = ""
	v.ExtSubOrderId = ""
	v.StoreCode = ""
	v.ChannelInvId = ""
	v.UserId = 0
	v.ScItemId = 0
	v.Success = false
	poolFuturePlanItemResult.Put(v)
}
