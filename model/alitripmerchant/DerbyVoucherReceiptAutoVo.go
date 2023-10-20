package alitripmerchant

import (
	"sync"
)

// DerbyVoucherReceiptAutoVo 结构体
type DerbyVoucherReceiptAutoVo struct {
	// 详情
	Data []DerbyVoucherReceiptAutoResponse `json:"data,omitempty" xml:"data>derby_voucher_receipt_auto_response,omitempty"`
	// 发票抬头
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 税号
	TaxID string `json:"tax_i_d,omitempty" xml:"tax_i_d,omitempty"`
}

var poolDerbyVoucherReceiptAutoVo = sync.Pool{
	New: func() any {
		return new(DerbyVoucherReceiptAutoVo)
	},
}

// GetDerbyVoucherReceiptAutoVo() 从对象池中获取DerbyVoucherReceiptAutoVo
func GetDerbyVoucherReceiptAutoVo() *DerbyVoucherReceiptAutoVo {
	return poolDerbyVoucherReceiptAutoVo.Get().(*DerbyVoucherReceiptAutoVo)
}

// ReleaseDerbyVoucherReceiptAutoVo 释放DerbyVoucherReceiptAutoVo
func ReleaseDerbyVoucherReceiptAutoVo(v *DerbyVoucherReceiptAutoVo) {
	v.Data = v.Data[:0]
	v.Name = ""
	v.TaxID = ""
	poolDerbyVoucherReceiptAutoVo.Put(v)
}
