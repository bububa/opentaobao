package einvoice

import (
	"sync"
)

// ResultList 结构体
type ResultList struct {
	// 开票明细列表
	InvoiceItems []InvoiceItems `json:"invoice_items,omitempty" xml:"invoice_items>invoice_items,omitempty"`
	// 付款方税号
	PayeeRegisterNo string `json:"payee_register_no,omitempty" xml:"payee_register_no,omitempty"`
	// 付款方平台
	Platform string `json:"platform,omitempty" xml:"platform,omitempty"`
	// 订单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 开票金额
	SumPrice string `json:"sum_price,omitempty" xml:"sum_price,omitempty"`
	// seriNo
	SeriNo string `json:"seri_no,omitempty" xml:"seri_no,omitempty"`
	// invoiceStatus
	InvoiceStatus int64 `json:"invoice_status,omitempty" xml:"invoice_status,omitempty"`
}

var poolResultList = sync.Pool{
	New: func() any {
		return new(ResultList)
	},
}

// GetResultList() 从对象池中获取ResultList
func GetResultList() *ResultList {
	return poolResultList.Get().(*ResultList)
}

// ReleaseResultList 释放ResultList
func ReleaseResultList(v *ResultList) {
	v.InvoiceItems = v.InvoiceItems[:0]
	v.PayeeRegisterNo = ""
	v.Platform = ""
	v.OrderId = ""
	v.SumPrice = ""
	v.SeriNo = ""
	v.InvoiceStatus = 0
	poolResultList.Put(v)
}
