package alitripmerchant

import (
	"sync"
)

// DerbyVoucherReceiptQueryVo 结构体
type DerbyVoucherReceiptQueryVo struct {
	// 邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 发票状态
	ReceiptStatus string `json:"receipt_status,omitempty" xml:"receipt_status,omitempty"`
	// 发票pdf url
	ReceiptURL string `json:"receipt_u_r_l,omitempty" xml:"receipt_u_r_l,omitempty"`
	// 订单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 税号
	TaxID string `json:"tax_i_d,omitempty" xml:"tax_i_d,omitempty"`
	// 总金额
	TotalRate string `json:"total_rate,omitempty" xml:"total_rate,omitempty"`
	// 发票类型(公司/个人)
	ReceiptType string `json:"receipt_type,omitempty" xml:"receipt_type,omitempty"`
	// 手机号
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
}

var poolDerbyVoucherReceiptQueryVo = sync.Pool{
	New: func() any {
		return new(DerbyVoucherReceiptQueryVo)
	},
}

// GetDerbyVoucherReceiptQueryVo() 从对象池中获取DerbyVoucherReceiptQueryVo
func GetDerbyVoucherReceiptQueryVo() *DerbyVoucherReceiptQueryVo {
	return poolDerbyVoucherReceiptQueryVo.Get().(*DerbyVoucherReceiptQueryVo)
}

// ReleaseDerbyVoucherReceiptQueryVo 释放DerbyVoucherReceiptQueryVo
func ReleaseDerbyVoucherReceiptQueryVo(v *DerbyVoucherReceiptQueryVo) {
	v.Email = ""
	v.Name = ""
	v.ReceiptStatus = ""
	v.ReceiptURL = ""
	v.OrderId = ""
	v.TaxID = ""
	v.TotalRate = ""
	v.ReceiptType = ""
	v.Phone = ""
	poolDerbyVoucherReceiptQueryVo.Put(v)
}
