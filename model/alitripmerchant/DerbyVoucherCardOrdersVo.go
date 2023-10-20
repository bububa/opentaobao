package alitripmerchant

import (
	"sync"
)

// DerbyVoucherCardOrdersVo 结构体
type DerbyVoucherCardOrdersVo struct {
	// 商品数量
	ProductCount string `json:"product_count,omitempty" xml:"product_count,omitempty"`
	// 实付金额
	ProductPaidIn string `json:"product_paid_in,omitempty" xml:"product_paid_in,omitempty"`
	// 下单时间
	OrderTime string `json:"order_time,omitempty" xml:"order_time,omitempty"`
	// 商品类型
	ProductMinType string `json:"product_min_type,omitempty" xml:"product_min_type,omitempty"`
	// 订单状态
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 权益卡状态
	CardStatus string `json:"card_status,omitempty" xml:"card_status,omitempty"`
	// 订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 商品图片
	ProductPic string `json:"product_pic,omitempty" xml:"product_pic,omitempty"`
}

var poolDerbyVoucherCardOrdersVo = sync.Pool{
	New: func() any {
		return new(DerbyVoucherCardOrdersVo)
	},
}

// GetDerbyVoucherCardOrdersVo() 从对象池中获取DerbyVoucherCardOrdersVo
func GetDerbyVoucherCardOrdersVo() *DerbyVoucherCardOrdersVo {
	return poolDerbyVoucherCardOrdersVo.Get().(*DerbyVoucherCardOrdersVo)
}

// ReleaseDerbyVoucherCardOrdersVo 释放DerbyVoucherCardOrdersVo
func ReleaseDerbyVoucherCardOrdersVo(v *DerbyVoucherCardOrdersVo) {
	v.ProductCount = ""
	v.ProductPaidIn = ""
	v.OrderTime = ""
	v.ProductMinType = ""
	v.OrderStatus = ""
	v.CardStatus = ""
	v.OrderId = ""
	v.ProductPic = ""
	poolDerbyVoucherCardOrdersVo.Put(v)
}
