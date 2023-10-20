package tmallcar

import (
	"sync"
)

// CarSubOrderDetailDto 结构体
type CarSubOrderDetailDto struct {
	// 商家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 买家昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 预约门店id
	EtShopId string `json:"et_shop_id,omitempty" xml:"et_shop_id,omitempty"`
	// 商家自定义参数
	SellerDefinitionField string `json:"seller_definition_field,omitempty" xml:"seller_definition_field,omitempty"`
	// 预约门店名称
	EtShopName string `json:"et_shop_name,omitempty" xml:"et_shop_name,omitempty"`
	// 核销门店名称
	ConsumeShopName string `json:"consume_shop_name,omitempty" xml:"consume_shop_name,omitempty"`
	// 核销门店id
	ConsumeShopId string `json:"consume_shop_id,omitempty" xml:"consume_shop_id,omitempty"`
	// 核销时间
	ConsumeTime string `json:"consume_time,omitempty" xml:"consume_time,omitempty"`
	// 选装选配json
	OptionList string `json:"option_list,omitempty" xml:"option_list,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 实付金额，单位分
	ActualPaidFee int64 `json:"actual_paid_fee,omitempty" xml:"actual_paid_fee,omitempty"`
	// 退款状态：1买家申请退款，等待卖家同意；2卖家已同意退款，等待买家退货；3买家已退货，等待卖家确认收货；4退款关闭；5退款成功；6卖家拒绝退款；9没有申请退款
	RefundStatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 支付状态：1未付款；2已付款；4已退款(交易关闭)；6交易成功；8未支付订单关闭
	PayStatus int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
}

var poolCarSubOrderDetailDto = sync.Pool{
	New: func() any {
		return new(CarSubOrderDetailDto)
	},
}

// GetCarSubOrderDetailDto() 从对象池中获取CarSubOrderDetailDto
func GetCarSubOrderDetailDto() *CarSubOrderDetailDto {
	return poolCarSubOrderDetailDto.Get().(*CarSubOrderDetailDto)
}

// ReleaseCarSubOrderDetailDto 释放CarSubOrderDetailDto
func ReleaseCarSubOrderDetailDto(v *CarSubOrderDetailDto) {
	v.SellerNick = ""
	v.BuyerNick = ""
	v.EtShopId = ""
	v.SellerDefinitionField = ""
	v.EtShopName = ""
	v.ConsumeShopName = ""
	v.ConsumeShopId = ""
	v.ConsumeTime = ""
	v.OptionList = ""
	v.OrderId = 0
	v.ActualPaidFee = 0
	v.RefundStatus = 0
	v.PayStatus = 0
	poolCarSubOrderDetailDto.Put(v)
}
