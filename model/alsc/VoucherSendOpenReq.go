package alsc

import (
	"sync"
)

// VoucherSendOpenReq 结构体
type VoucherSendOpenReq struct {
	// 外部品牌ID
	OutBrandId string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
	// 请求id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 品牌id
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 外部门店id
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 用户id
	CustomerId string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
	// 门店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 操作人id
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 券模板id
	VoucherId string `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
	// 发券数量
	VoucherNum int64 `json:"voucher_num,omitempty" xml:"voucher_num,omitempty"`
}

var poolVoucherSendOpenReq = sync.Pool{
	New: func() any {
		return new(VoucherSendOpenReq)
	},
}

// GetVoucherSendOpenReq() 从对象池中获取VoucherSendOpenReq
func GetVoucherSendOpenReq() *VoucherSendOpenReq {
	return poolVoucherSendOpenReq.Get().(*VoucherSendOpenReq)
}

// ReleaseVoucherSendOpenReq 释放VoucherSendOpenReq
func ReleaseVoucherSendOpenReq(v *VoucherSendOpenReq) {
	v.OutBrandId = ""
	v.RequestId = ""
	v.BrandId = ""
	v.OutShopId = ""
	v.CustomerId = ""
	v.ShopId = ""
	v.OperatorId = ""
	v.VoucherId = ""
	v.VoucherNum = 0
	poolVoucherSendOpenReq.Put(v)
}
