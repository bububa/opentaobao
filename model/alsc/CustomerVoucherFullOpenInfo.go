package alsc

import (
	"sync"
)

// CustomerVoucherFullOpenInfo 结构体
type CustomerVoucherFullOpenInfo struct {
	// 券失效时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 获取时间
	GmtCreated string `json:"gmt_created,omitempty" xml:"gmt_created,omitempty"`
	// 券生效时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 优惠券状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 券名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 优惠券实例ID
	VoucherId string `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
	// 优惠券类型，CASH：代金券，DISCOUNT：折扣券，GIFT：礼品券
	VoucherType string `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
	// 优惠率
	DiscountRate string `json:"discount_rate,omitempty" xml:"discount_rate,omitempty"`
	// 优惠券模版ID
	VoucherTemplateId string `json:"voucher_template_id,omitempty" xml:"voucher_template_id,omitempty"`
	// 优惠券金额，单位：分
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 顾客ID
	CustomerId int64 `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
	// 门店ID
	ShopId int64 `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 满足金额阀值，单位：分
	StartFee int64 `json:"start_fee,omitempty" xml:"start_fee,omitempty"`
	// 券实例有效点数
	GiftPoint int64 `json:"gift_point,omitempty" xml:"gift_point,omitempty"`
}

var poolCustomerVoucherFullOpenInfo = sync.Pool{
	New: func() any {
		return new(CustomerVoucherFullOpenInfo)
	},
}

// GetCustomerVoucherFullOpenInfo() 从对象池中获取CustomerVoucherFullOpenInfo
func GetCustomerVoucherFullOpenInfo() *CustomerVoucherFullOpenInfo {
	return poolCustomerVoucherFullOpenInfo.Get().(*CustomerVoucherFullOpenInfo)
}

// ReleaseCustomerVoucherFullOpenInfo 释放CustomerVoucherFullOpenInfo
func ReleaseCustomerVoucherFullOpenInfo(v *CustomerVoucherFullOpenInfo) {
	v.EndTime = ""
	v.GmtCreated = ""
	v.StartTime = ""
	v.Status = ""
	v.Title = ""
	v.VoucherId = ""
	v.VoucherType = ""
	v.DiscountRate = ""
	v.VoucherTemplateId = ""
	v.Amount = 0
	v.CustomerId = 0
	v.ShopId = 0
	v.StartFee = 0
	v.GiftPoint = 0
	poolCustomerVoucherFullOpenInfo.Put(v)
}
