package alsc

import (
	"sync"
)

// CustomerVoucherFullOpenReq 结构体
type CustomerVoucherFullOpenReq struct {
	// 优惠券状态 NORMAL，DELETED，ISUSED
	VoucherStatusList []string `json:"voucher_status_list,omitempty" xml:"voucher_status_list>string,omitempty"`
	// 品牌ID
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 顾客ID
	CustomerId string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
	// 外部品牌ID
	OutBrandId string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
	// 外部门店ID
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// SaaS门店ID
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 第几页,从1开始计数
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每页大小，默认20
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolCustomerVoucherFullOpenReq = sync.Pool{
	New: func() any {
		return new(CustomerVoucherFullOpenReq)
	},
}

// GetCustomerVoucherFullOpenReq() 从对象池中获取CustomerVoucherFullOpenReq
func GetCustomerVoucherFullOpenReq() *CustomerVoucherFullOpenReq {
	return poolCustomerVoucherFullOpenReq.Get().(*CustomerVoucherFullOpenReq)
}

// ReleaseCustomerVoucherFullOpenReq 释放CustomerVoucherFullOpenReq
func ReleaseCustomerVoucherFullOpenReq(v *CustomerVoucherFullOpenReq) {
	v.VoucherStatusList = v.VoucherStatusList[:0]
	v.BrandId = ""
	v.CustomerId = ""
	v.OutBrandId = ""
	v.OutShopId = ""
	v.ShopId = ""
	v.PageNo = 0
	v.PageSize = 0
	poolCustomerVoucherFullOpenReq.Put(v)
}
