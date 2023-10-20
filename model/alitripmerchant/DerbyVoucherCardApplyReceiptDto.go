package alitripmerchant

import (
	"sync"
)

// DerbyVoucherCardApplyReceiptDto 结构体
type DerbyVoucherCardApplyReceiptDto struct {
	// 发票类型(个人/公司)
	ReceiptType string `json:"receipt_type,omitempty" xml:"receipt_type,omitempty"`
	// 张三
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 13365231234
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 1456@163.com
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 84531204051
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 321312312
	TaxId string `json:"tax_id,omitempty" xml:"tax_id,omitempty"`
	// 2131
	TotalRate string `json:"total_rate,omitempty" xml:"total_rate,omitempty"`
}

var poolDerbyVoucherCardApplyReceiptDto = sync.Pool{
	New: func() any {
		return new(DerbyVoucherCardApplyReceiptDto)
	},
}

// GetDerbyVoucherCardApplyReceiptDto() 从对象池中获取DerbyVoucherCardApplyReceiptDto
func GetDerbyVoucherCardApplyReceiptDto() *DerbyVoucherCardApplyReceiptDto {
	return poolDerbyVoucherCardApplyReceiptDto.Get().(*DerbyVoucherCardApplyReceiptDto)
}

// ReleaseDerbyVoucherCardApplyReceiptDto 释放DerbyVoucherCardApplyReceiptDto
func ReleaseDerbyVoucherCardApplyReceiptDto(v *DerbyVoucherCardApplyReceiptDto) {
	v.ReceiptType = ""
	v.Name = ""
	v.Phone = ""
	v.Email = ""
	v.OrderId = ""
	v.TaxId = ""
	v.TotalRate = ""
	poolDerbyVoucherCardApplyReceiptDto.Put(v)
}
