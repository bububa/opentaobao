package refund

import (
	"sync"
)

// CancelGoodsDto 结构体
type CancelGoodsDto struct {
	// 操作时间
	OperateTime string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 状态SUCCESS、FAIL
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 商家商品ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 子订单ID
	Oid int64 `json:"oid,omitempty" xml:"oid,omitempty"`
	// 退款单ID
	RefundId int64 `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	// 退款金额 单位 分
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 商品ID
	AuctionId int64 `json:"auction_id,omitempty" xml:"auction_id,omitempty"`
	// 商品数量
	AuctionNum int64 `json:"auction_num,omitempty" xml:"auction_num,omitempty"`
	// 主订单ID
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}

var poolCancelGoodsDto = sync.Pool{
	New: func() any {
		return new(CancelGoodsDto)
	},
}

// GetCancelGoodsDto() 从对象池中获取CancelGoodsDto
func GetCancelGoodsDto() *CancelGoodsDto {
	return poolCancelGoodsDto.Get().(*CancelGoodsDto)
}

// ReleaseCancelGoodsDto 释放CancelGoodsDto
func ReleaseCancelGoodsDto(v *CancelGoodsDto) {
	v.OperateTime = ""
	v.Status = ""
	v.OuterId = ""
	v.Msg = ""
	v.Oid = 0
	v.RefundId = 0
	v.RefundFee = 0
	v.AuctionId = 0
	v.AuctionNum = 0
	v.Tid = 0
	poolCancelGoodsDto.Put(v)
}
