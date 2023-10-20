package alsc

import (
	"sync"
)

// InvoiceResultFeedbackTopReq 结构体
type InvoiceResultFeedbackTopReq struct {
	// 实际开票金额
	RealAmount string `json:"real_amount,omitempty" xml:"real_amount,omitempty"`
	// 口碑开票申请id
	InvoiceApplyId string `json:"invoice_apply_id,omitempty" xml:"invoice_apply_id,omitempty"`
	// 运单号
	ExpressNumber string `json:"express_number,omitempty" xml:"express_number,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误原因
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 发票的票号
	InvoiceId string `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	// 电子发票地址
	EmailInvoiceImage string `json:"email_invoice_image,omitempty" xml:"email_invoice_image,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolInvoiceResultFeedbackTopReq = sync.Pool{
	New: func() any {
		return new(InvoiceResultFeedbackTopReq)
	},
}

// GetInvoiceResultFeedbackTopReq() 从对象池中获取InvoiceResultFeedbackTopReq
func GetInvoiceResultFeedbackTopReq() *InvoiceResultFeedbackTopReq {
	return poolInvoiceResultFeedbackTopReq.Get().(*InvoiceResultFeedbackTopReq)
}

// ReleaseInvoiceResultFeedbackTopReq 释放InvoiceResultFeedbackTopReq
func ReleaseInvoiceResultFeedbackTopReq(v *InvoiceResultFeedbackTopReq) {
	v.RealAmount = ""
	v.InvoiceApplyId = ""
	v.ExpressNumber = ""
	v.ErrCode = ""
	v.ErrMsg = ""
	v.InvoiceId = ""
	v.EmailInvoiceImage = ""
	v.Success = false
	poolInvoiceResultFeedbackTopReq.Put(v)
}
