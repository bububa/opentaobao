package wdk

import (
	"sync"
)

// PosOrderCreateRequest 结构体
type PosOrderCreateRequest struct {
	// 子订单列表
	SubOrderDOList []PosSubOrderDo `json:"sub_order_d_o_list,omitempty" xml:"sub_order_d_o_list>pos_sub_order_do,omitempty"`
	// 支付方式
	PayChannelList []PosPayChannel `json:"pay_channel_list,omitempty" xml:"pay_channel_list>pos_pay_channel,omitempty"`
	// 支付时间，必填
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 外部主订单号，必填
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 经营店code，必填
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 渠道店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 会员卡号
	MemberCardNum string `json:"member_card_num,omitempty" xml:"member_card_num,omitempty"`
	// 兼容老接口的数据
	OldData int64 `json:"old_data,omitempty" xml:"old_data,omitempty"`
}

var poolPosOrderCreateRequest = sync.Pool{
	New: func() any {
		return new(PosOrderCreateRequest)
	},
}

// GetPosOrderCreateRequest() 从对象池中获取PosOrderCreateRequest
func GetPosOrderCreateRequest() *PosOrderCreateRequest {
	return poolPosOrderCreateRequest.Get().(*PosOrderCreateRequest)
}

// ReleasePosOrderCreateRequest 释放PosOrderCreateRequest
func ReleasePosOrderCreateRequest(v *PosOrderCreateRequest) {
	v.SubOrderDOList = v.SubOrderDOList[:0]
	v.PayChannelList = v.PayChannelList[:0]
	v.PayTime = ""
	v.OutOrderId = ""
	v.StoreId = ""
	v.ShopId = ""
	v.MemberCardNum = ""
	v.OldData = 0
	poolPosOrderCreateRequest.Put(v)
}
