package wdk

import (
	"sync"
)

// OrderBalanceBillDo 结构体
type OrderBalanceBillDo struct {
	// 订单原价
	ParentTotalPrice string `json:"parent_total_price,omitempty" xml:"parent_total_price,omitempty"`
	// 订单付款时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 分账单创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 交易渠道 1：其他 2：app 3：pos
	OrderChannel string `json:"order_channel,omitempty" xml:"order_channel,omitempty"`
	// 父订单
	MordId string `json:"mord_id,omitempty" xml:"mord_id,omitempty"`
	// 门店名字
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 名店id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 子公司
	SubsidiaryName string `json:"subsidiary_name,omitempty" xml:"subsidiary_name,omitempty"`
	// 子公司code
	SubsidiaryCode string `json:"subsidiary_code,omitempty" xml:"subsidiary_code,omitempty"`
	// 集团
	MerchantName string `json:"merchant_name,omitempty" xml:"merchant_name,omitempty"`
	// 集团code
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 业务日期
	Thedate string `json:"thedate,omitempty" xml:"thedate,omitempty"`
	// 创建时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 订单技术服务费
	PayTechFee string `json:"pay_tech_fee,omitempty" xml:"pay_tech_fee,omitempty"`
	// 用户实付金额，减支付宝红包
	PayFee string `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
	// 订单付款金额
	ParentOriginalPrice string `json:"parent_original_price,omitempty" xml:"parent_original_price,omitempty"`
	// 总优惠金额
	DiscountFee string `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 支付优惠金额
	PayDiscountFee string `json:"pay_discount_fee,omitempty" xml:"pay_discount_fee,omitempty"`
	// 分账优惠金额
	DiscountFzFee string `json:"discount_fz_fee,omitempty" xml:"discount_fz_fee,omitempty"`
	// 订单优惠金额平台补贴商家费用
	PmtDiscountFee string `json:"pmt_discount_fee,omitempty" xml:"pmt_discount_fee,omitempty"`
	// 订单优惠金额
	ParentDiscountFee string `json:"parent_discount_fee,omitempty" xml:"parent_discount_fee,omitempty"`
	// 渠道来源 ：1 淘点点；2 美团；3 饿了么；4 淘宝TC；5 百度外卖；6 微信 ;7 三江；18 大润发飞牛优鲜；
	OrderFrom string `json:"order_from,omitempty" xml:"order_from,omitempty"`
	// 支付方式
	PayChannel string `json:"pay_channel,omitempty" xml:"pay_channel,omitempty"`
	// 单据类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 商家应收金额
	SellerReceiveFee string `json:"seller_receive_fee,omitempty" xml:"seller_receive_fee,omitempty"`
	// 优惠手续费
	DiscountChangeFee string `json:"discount_change_fee,omitempty" xml:"discount_change_fee,omitempty"`
	// 优惠技术服务费
	DiscountTechFee string `json:"discount_tech_fee,omitempty" xml:"discount_tech_fee,omitempty"`
	// 商家承担优惠金额
	DiscountSellerFee string `json:"discount_seller_fee,omitempty" xml:"discount_seller_fee,omitempty"`
	// 平台承担优惠金额
	DiscountPlatformFee string `json:"discount_platform_fee,omitempty" xml:"discount_platform_fee,omitempty"`
	// 分账金额
	FzFee string `json:"fz_fee,omitempty" xml:"fz_fee,omitempty"`
	// 配送费
	DeliveryFee string `json:"delivery_fee,omitempty" xml:"delivery_fee,omitempty"`
	// 订单手续费
	PayChangeFee string `json:"pay_change_fee,omitempty" xml:"pay_change_fee,omitempty"`
	// 原订单运费
	ParentPostFee string `json:"parent_post_fee,omitempty" xml:"parent_post_fee,omitempty"`
	// 用户应付金额（含运费）
	UserPayAmount string `json:"user_pay_amount,omitempty" xml:"user_pay_amount,omitempty"`
	// 支付宝收款
	PayByAliPay string `json:"pay_by_ali_pay,omitempty" xml:"pay_by_ali_pay,omitempty"`
	// 非支付宝收款
	PayByOther string `json:"pay_by_other,omitempty" xml:"pay_by_other,omitempty"`
	// 收入金额小计
	IncomeAmount string `json:"income_amount,omitempty" xml:"income_amount,omitempty"`
	// 打包时间
	PackageTime string `json:"package_time,omitempty" xml:"package_time,omitempty"`
	// 正向打包状态
	PackagedStatus string `json:"packaged_status,omitempty" xml:"packaged_status,omitempty"`
	// 主键id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 是否退款计算收入(1:是，0:否) 这是逆向的
	Income int64 `json:"income,omitempty" xml:"income,omitempty"`
}

var poolOrderBalanceBillDo = sync.Pool{
	New: func() any {
		return new(OrderBalanceBillDo)
	},
}

// GetOrderBalanceBillDo() 从对象池中获取OrderBalanceBillDo
func GetOrderBalanceBillDo() *OrderBalanceBillDo {
	return poolOrderBalanceBillDo.Get().(*OrderBalanceBillDo)
}

// ReleaseOrderBalanceBillDo 释放OrderBalanceBillDo
func ReleaseOrderBalanceBillDo(v *OrderBalanceBillDo) {
	v.ParentTotalPrice = ""
	v.PayTime = ""
	v.GmtCreate = ""
	v.OrderChannel = ""
	v.MordId = ""
	v.StoreName = ""
	v.StoreId = ""
	v.SubsidiaryName = ""
	v.SubsidiaryCode = ""
	v.MerchantName = ""
	v.MerchantCode = ""
	v.Thedate = ""
	v.GmtModified = ""
	v.PayTechFee = ""
	v.PayFee = ""
	v.ParentOriginalPrice = ""
	v.DiscountFee = ""
	v.PayDiscountFee = ""
	v.DiscountFzFee = ""
	v.PmtDiscountFee = ""
	v.ParentDiscountFee = ""
	v.OrderFrom = ""
	v.PayChannel = ""
	v.Type = ""
	v.SellerReceiveFee = ""
	v.DiscountChangeFee = ""
	v.DiscountTechFee = ""
	v.DiscountSellerFee = ""
	v.DiscountPlatformFee = ""
	v.FzFee = ""
	v.DeliveryFee = ""
	v.PayChangeFee = ""
	v.ParentPostFee = ""
	v.UserPayAmount = ""
	v.PayByAliPay = ""
	v.PayByOther = ""
	v.IncomeAmount = ""
	v.PackageTime = ""
	v.PackagedStatus = ""
	v.Id = 0
	v.Income = 0
	poolOrderBalanceBillDo.Put(v)
}
