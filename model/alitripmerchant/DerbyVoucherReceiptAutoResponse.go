package alitripmerchant

import (
	"sync"
)

// DerbyVoucherReceiptAutoResponse 结构体
type DerbyVoucherReceiptAutoResponse struct {
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 税号
	TaxID string `json:"tax_i_d,omitempty" xml:"tax_i_d,omitempty"`
}

var poolDerbyVoucherReceiptAutoResponse = sync.Pool{
	New: func() any {
		return new(DerbyVoucherReceiptAutoResponse)
	},
}

// GetDerbyVoucherReceiptAutoResponse() 从对象池中获取DerbyVoucherReceiptAutoResponse
func GetDerbyVoucherReceiptAutoResponse() *DerbyVoucherReceiptAutoResponse {
	return poolDerbyVoucherReceiptAutoResponse.Get().(*DerbyVoucherReceiptAutoResponse)
}

// ReleaseDerbyVoucherReceiptAutoResponse 释放DerbyVoucherReceiptAutoResponse
func ReleaseDerbyVoucherReceiptAutoResponse(v *DerbyVoucherReceiptAutoResponse) {
	v.Name = ""
	v.TaxID = ""
	poolDerbyVoucherReceiptAutoResponse.Put(v)
}
