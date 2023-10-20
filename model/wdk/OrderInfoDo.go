package wdk

import (
	"sync"
)

// OrderInfoDo 结构体
type OrderInfoDo struct {
	// 小票付款渠道
	PayChannels []ReceiptPayChannelDo `json:"pay_channels,omitempty" xml:"pay_channels>receipt_pay_channel_do,omitempty"`
	// 商品明细
	SubOrders []SubOrderInfoDo `json:"sub_orders,omitempty" xml:"sub_orders>sub_order_info_do,omitempty"`
	// 来源，商家级别，当前取值：RENRENLE / SANJIANG
	OrderFrom string `json:"order_from,omitempty" xml:"order_from,omitempty"`
	// 原交易时间
	OriginalTrdTime string `json:"original_trd_time,omitempty" xml:"original_trd_time,omitempty"`
	// 原流水号
	OriginalSerialNum string `json:"original_serial_num,omitempty" xml:"original_serial_num,omitempty"`
	// 原款机号
	OriginalPosNo string `json:"original_pos_no,omitempty" xml:"original_pos_no,omitempty"`
	// 会员卡号
	MemberCardNum string `json:"member_card_num,omitempty" xml:"member_card_num,omitempty"`
	// 收银员姓名
	OpName string `json:"op_name,omitempty" xml:"op_name,omitempty"`
	// 收银员工号
	OpNum string `json:"op_num,omitempty" xml:"op_num,omitempty"`
	// 成交时间
	TrdTime string `json:"trd_time,omitempty" xml:"trd_time,omitempty"`
	// 流水号
	SerialNum string `json:"serial_num,omitempty" xml:"serial_num,omitempty"`
	// 款机号
	PosNo string `json:"pos_no,omitempty" xml:"pos_no,omitempty"`
	// 渠道店id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 阿里userID
	AliUserId string `json:"ali_user_id,omitempty" xml:"ali_user_id,omitempty"`
	// 折扣优惠金额
	DiscountAmt int64 `json:"discount_amt,omitempty" xml:"discount_amt,omitempty"`
	// 损溢金额
	OverflowAmt int64 `json:"overflow_amt,omitempty" xml:"overflow_amt,omitempty"`
	// 会员优惠
	MemberDiscount int64 `json:"member_discount,omitempty" xml:"member_discount,omitempty"`
	// 找零金额
	ChangeAmt int64 `json:"change_amt,omitempty" xml:"change_amt,omitempty"`
	// 实收金额
	ActualAmt int64 `json:"actual_amt,omitempty" xml:"actual_amt,omitempty"`
	// 应收金额
	AskAmt int64 `json:"ask_amt,omitempty" xml:"ask_amt,omitempty"`
	// 交易类型
	TrdType int64 `json:"trd_type,omitempty" xml:"trd_type,omitempty"`
}

var poolOrderInfoDo = sync.Pool{
	New: func() any {
		return new(OrderInfoDo)
	},
}

// GetOrderInfoDo() 从对象池中获取OrderInfoDo
func GetOrderInfoDo() *OrderInfoDo {
	return poolOrderInfoDo.Get().(*OrderInfoDo)
}

// ReleaseOrderInfoDo 释放OrderInfoDo
func ReleaseOrderInfoDo(v *OrderInfoDo) {
	v.PayChannels = v.PayChannels[:0]
	v.SubOrders = v.SubOrders[:0]
	v.OrderFrom = ""
	v.OriginalTrdTime = ""
	v.OriginalSerialNum = ""
	v.OriginalPosNo = ""
	v.MemberCardNum = ""
	v.OpName = ""
	v.OpNum = ""
	v.TrdTime = ""
	v.SerialNum = ""
	v.PosNo = ""
	v.StoreId = ""
	v.AliUserId = ""
	v.DiscountAmt = 0
	v.OverflowAmt = 0
	v.MemberDiscount = 0
	v.ChangeAmt = 0
	v.ActualAmt = 0
	v.AskAmt = 0
	v.TrdType = 0
	poolOrderInfoDo.Put(v)
}
