package alitripmerchant

import (
	"sync"
)

// DerbyVoucherCardPurchasableVo 结构体
type DerbyVoucherCardPurchasableVo struct {
	// 订单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 时间戳
	TimeStamp string `json:"time_stamp,omitempty" xml:"time_stamp,omitempty"`
	// 随机字符串
	NonceStr string `json:"nonce_str,omitempty" xml:"nonce_str,omitempty"`
	// 支付认证
	Packag string `json:"packag,omitempty" xml:"packag,omitempty"`
	// 签名
	PaySign string `json:"pay_sign,omitempty" xml:"pay_sign,omitempty"`
	// 签名方式
	SignType string `json:"sign_type,omitempty" xml:"sign_type,omitempty"`
	// 权益卡名称（备用字段，暂时无用）
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 权益卡商品编码（下单要用）
	VoucherCardCode string `json:"voucher_card_code,omitempty" xml:"voucher_card_code,omitempty"`
	// 权益卡图片
	CardImage string `json:"card_image,omitempty" xml:"card_image,omitempty"`
	// 权益卡规格 DerbyVoucherCardOrderProductMinTypeEnum code
	VoucherCardCategory string `json:"voucher_card_category,omitempty" xml:"voucher_card_category,omitempty"`
	// 权益卡规格 DerbyVoucherCardOrderProductMinTypeEnum desc
	VoucherCardCategoryName string `json:"voucher_card_category_name,omitempty" xml:"voucher_card_category_name,omitempty"`
	// 商品名称 accor_paid_membership_card雅高付费会员卡 code
	ProductType string `json:"product_type,omitempty" xml:"product_type,omitempty"`
	// 商品名称 accor_paid_membership_card雅高付费会员卡 desc
	ProductTypeName string `json:"product_type_name,omitempty" xml:"product_type_name,omitempty"`
	// 优惠的原因
	PriceDesc string `json:"price_desc,omitempty" xml:"price_desc,omitempty"`
	// 权益卡原价
	Price float64 `json:"price,omitempty" xml:"price,omitempty"`
	// 权益卡售卖价 （为空则选原价）
	SalePrice float64 `json:"sale_price,omitempty" xml:"sale_price,omitempty"`
	// 权益到期时间搓
	CountDownTime int64 `json:"count_down_time,omitempty" xml:"count_down_time,omitempty"`
	// 商品数量
	ProductCount int64 `json:"product_count,omitempty" xml:"product_count,omitempty"`
	// 是否需要注册 true需要 false不需要
	NeedRegister bool `json:"need_register,omitempty" xml:"need_register,omitempty"`
}

var poolDerbyVoucherCardPurchasableVo = sync.Pool{
	New: func() any {
		return new(DerbyVoucherCardPurchasableVo)
	},
}

// GetDerbyVoucherCardPurchasableVo() 从对象池中获取DerbyVoucherCardPurchasableVo
func GetDerbyVoucherCardPurchasableVo() *DerbyVoucherCardPurchasableVo {
	return poolDerbyVoucherCardPurchasableVo.Get().(*DerbyVoucherCardPurchasableVo)
}

// ReleaseDerbyVoucherCardPurchasableVo 释放DerbyVoucherCardPurchasableVo
func ReleaseDerbyVoucherCardPurchasableVo(v *DerbyVoucherCardPurchasableVo) {
	v.OrderId = ""
	v.TimeStamp = ""
	v.NonceStr = ""
	v.Packag = ""
	v.PaySign = ""
	v.SignType = ""
	v.Name = ""
	v.VoucherCardCode = ""
	v.CardImage = ""
	v.VoucherCardCategory = ""
	v.VoucherCardCategoryName = ""
	v.ProductType = ""
	v.ProductTypeName = ""
	v.PriceDesc = ""
	v.Price = 0
	v.SalePrice = 0
	v.CountDownTime = 0
	v.ProductCount = 0
	v.NeedRegister = false
	poolDerbyVoucherCardPurchasableVo.Put(v)
}
