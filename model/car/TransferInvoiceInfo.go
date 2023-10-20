package car

import (
	"sync"
)

// TransferInvoiceInfo 结构体
type TransferInvoiceInfo struct {
	// 发票寄送邮箱（用于电子发票）
	EInvoiceSendEmail string `json:"e_invoice_send_email,omitempty" xml:"e_invoice_send_email,omitempty"`
	// 企业（公司）税号
	TaxNo string `json:"tax_no,omitempty" xml:"tax_no,omitempty"`
	// 发票抬头
	InvoiceTitle string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	// 开票金额(元)
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 发票类型 1企业 2个人
	InvoiceType int64 `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	// -1:未开具1:开具中;2:开具完成;3:已发送
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolTransferInvoiceInfo = sync.Pool{
	New: func() any {
		return new(TransferInvoiceInfo)
	},
}

// GetTransferInvoiceInfo() 从对象池中获取TransferInvoiceInfo
func GetTransferInvoiceInfo() *TransferInvoiceInfo {
	return poolTransferInvoiceInfo.Get().(*TransferInvoiceInfo)
}

// ReleaseTransferInvoiceInfo 释放TransferInvoiceInfo
func ReleaseTransferInvoiceInfo(v *TransferInvoiceInfo) {
	v.EInvoiceSendEmail = ""
	v.TaxNo = ""
	v.InvoiceTitle = ""
	v.Amount = ""
	v.InvoiceType = 0
	v.Status = 0
	poolTransferInvoiceInfo.Put(v)
}
