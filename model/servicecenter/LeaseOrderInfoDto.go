package servicecenter

import (
	"sync"
)

// LeaseOrderInfoDto 结构体
type LeaseOrderInfoDto struct {
	// 网商申请号
	ApplyNo string `json:"apply_no,omitempty" xml:"apply_no,omitempty"`
	// 单状态（请关注此状态，如果为TRADE_CLOSED_BY_TAOBAO状态，则不要对此订单进行发货，切记啊！）。可选值: TRADE_NO_CREATE_PAY(没有创建支付宝交易) WAIT_BUYER_PAY(等待买家付款) WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款) WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货) TRADE_BUYER_SIGNED(买家已签收,货到付款专用) TRADE_FINISHED(交易成功) TRADE_CLOSED(付款以后用户退款成功，交易自动关闭) TRADE_CLOSED_BY_TAOBAO(付款以前，卖家或买家主动关闭交易)PAY_PENDING(国际信用卡支付付款确认中)
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 门店自定义编码
	StoreOutId string `json:"store_out_id,omitempty" xml:"store_out_id,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolLeaseOrderInfoDto = sync.Pool{
	New: func() any {
		return new(LeaseOrderInfoDto)
	},
}

// GetLeaseOrderInfoDto() 从对象池中获取LeaseOrderInfoDto
func GetLeaseOrderInfoDto() *LeaseOrderInfoDto {
	return poolLeaseOrderInfoDto.Get().(*LeaseOrderInfoDto)
}

// ReleaseLeaseOrderInfoDto 释放LeaseOrderInfoDto
func ReleaseLeaseOrderInfoDto(v *LeaseOrderInfoDto) {
	v.ApplyNo = ""
	v.Status = ""
	v.StoreOutId = ""
	v.OrderId = 0
	poolLeaseOrderInfoDto.Put(v)
}
