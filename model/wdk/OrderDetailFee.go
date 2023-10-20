package wdk

import (
	"sync"
)

// OrderDetailFee 结构体
type OrderDetailFee struct {
	// 众包呼单费
	ZhongbaoCallFee string `json:"zhongbao_call_fee,omitempty" xml:"zhongbao_call_fee,omitempty"`
	// 冷链加工费
	ColdBoxFee string `json:"cold_box_fee,omitempty" xml:"cold_box_fee,omitempty"`
	// 用户实付
	UserFee string `json:"user_fee,omitempty" xml:"user_fee,omitempty"`
	// 实收佣金
	Commission string `json:"commission,omitempty" xml:"commission,omitempty"`
	// 代理商补贴
	AgentRate string `json:"agent_rate,omitempty" xml:"agent_rate,omitempty"`
	// 平台补贴
	PlatformRate string `json:"platform_rate,omitempty" xml:"platform_rate,omitempty"`
	// 商户补贴
	ShopRate string `json:"shop_rate,omitempty" xml:"shop_rate,omitempty"`
	// 配送费
	SendFee string `json:"send_fee,omitempty" xml:"send_fee,omitempty"`
	// 餐盒费
	PackageFee string `json:"package_fee,omitempty" xml:"package_fee,omitempty"`
	// 商品金额
	ProductFee string `json:"product_fee,omitempty" xml:"product_fee,omitempty"`
	// 渠道费
	PayChannelFee string `json:"pay_channel_fee,omitempty" xml:"pay_channel_fee,omitempty"`
	// 渠道费率
	PayChannelRate string `json:"pay_channel_rate,omitempty" xml:"pay_channel_rate,omitempty"`
}

var poolOrderDetailFee = sync.Pool{
	New: func() any {
		return new(OrderDetailFee)
	},
}

// GetOrderDetailFee() 从对象池中获取OrderDetailFee
func GetOrderDetailFee() *OrderDetailFee {
	return poolOrderDetailFee.Get().(*OrderDetailFee)
}

// ReleaseOrderDetailFee 释放OrderDetailFee
func ReleaseOrderDetailFee(v *OrderDetailFee) {
	v.ZhongbaoCallFee = ""
	v.ColdBoxFee = ""
	v.UserFee = ""
	v.Commission = ""
	v.AgentRate = ""
	v.PlatformRate = ""
	v.ShopRate = ""
	v.SendFee = ""
	v.PackageFee = ""
	v.ProductFee = ""
	v.PayChannelFee = ""
	v.PayChannelRate = ""
	poolOrderDetailFee.Put(v)
}
