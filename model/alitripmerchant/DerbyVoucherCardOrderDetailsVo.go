package alitripmerchant

import (
	"sync"
)

// DerbyVoucherCardOrderDetailsVo 结构体
type DerbyVoucherCardOrderDetailsVo struct {
	// 状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 商品价格
	ProductPrice string `json:"product_price,omitempty" xml:"product_price,omitempty"`
	// 商品数量
	ProductCount string `json:"product_count,omitempty" xml:"product_count,omitempty"`
	// 实付金额
	ProductPaidIn string `json:"product_paid_in,omitempty" xml:"product_paid_in,omitempty"`
	// 下单时间
	OrderTime string `json:"order_time,omitempty" xml:"order_time,omitempty"`
	// 订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 商品类型
	ProductMinType string `json:"product_min_type,omitempty" xml:"product_min_type,omitempty"`
	// 会员卡号
	CardNumber string `json:"card_number,omitempty" xml:"card_number,omitempty"`
	// 权益卡号
	MemberVoucherCardID string `json:"member_voucher_card_i_d,omitempty" xml:"member_voucher_card_i_d,omitempty"`
	// 卡状态
	CardStatus string `json:"card_status,omitempty" xml:"card_status,omitempty"`
	// 商品图片
	ProductPic string `json:"product_pic,omitempty" xml:"product_pic,omitempty"`
	// 弹屏
	PopScreen bool `json:"pop_screen,omitempty" xml:"pop_screen,omitempty"`
}

var poolDerbyVoucherCardOrderDetailsVo = sync.Pool{
	New: func() any {
		return new(DerbyVoucherCardOrderDetailsVo)
	},
}

// GetDerbyVoucherCardOrderDetailsVo() 从对象池中获取DerbyVoucherCardOrderDetailsVo
func GetDerbyVoucherCardOrderDetailsVo() *DerbyVoucherCardOrderDetailsVo {
	return poolDerbyVoucherCardOrderDetailsVo.Get().(*DerbyVoucherCardOrderDetailsVo)
}

// ReleaseDerbyVoucherCardOrderDetailsVo 释放DerbyVoucherCardOrderDetailsVo
func ReleaseDerbyVoucherCardOrderDetailsVo(v *DerbyVoucherCardOrderDetailsVo) {
	v.Status = ""
	v.ProductPrice = ""
	v.ProductCount = ""
	v.ProductPaidIn = ""
	v.OrderTime = ""
	v.OrderId = ""
	v.ProductMinType = ""
	v.CardNumber = ""
	v.MemberVoucherCardID = ""
	v.CardStatus = ""
	v.ProductPic = ""
	v.PopScreen = false
	poolDerbyVoucherCardOrderDetailsVo.Put(v)
}
