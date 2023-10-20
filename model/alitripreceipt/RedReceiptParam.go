package alitripreceipt

import (
	"sync"
)

// RedReceiptParam 结构体
type RedReceiptParam struct {
	// 发票备注
	ReceiptMemo string `json:"receipt_memo,omitempty" xml:"receipt_memo,omitempty"`
	// 接收人
	Receiver string `json:"receiver,omitempty" xml:"receiver,omitempty"`
	// 发票邮箱
	ReceiveMail string `json:"receive_mail,omitempty" xml:"receive_mail,omitempty"`
	// 原始发票号(冲红)
	OriginReceiptNumber string `json:"origin_receipt_number,omitempty" xml:"origin_receipt_number,omitempty"`
	// 企业名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 代理商订单号
	AgentOrderNo string `json:"agent_order_no,omitempty" xml:"agent_order_no,omitempty"`
	// 发票内容
	ReceiptContent string `json:"receipt_content,omitempty" xml:"receipt_content,omitempty"`
	// 接收人开户行名称
	ReceiveBankName string `json:"receive_bank_name,omitempty" xml:"receive_bank_name,omitempty"`
	// 接收人开户行账号
	ReceiveBankAccount string `json:"receive_bank_account,omitempty" xml:"receive_bank_account,omitempty"`
	// 扩展参数
	ExtMap string `json:"ext_map,omitempty" xml:"ext_map,omitempty"`
	// 发票抬头
	ReceiptTitle string `json:"receipt_title,omitempty" xml:"receipt_title,omitempty"`
	// 企业税号
	CompanyTaxNo string `json:"company_tax_no,omitempty" xml:"company_tax_no,omitempty"`
	// 接收人手机号
	ReceiveMobile string `json:"receive_mobile,omitempty" xml:"receive_mobile,omitempty"`
	// 发票抬头类型,0:企业;1:个人
	ReceiptTitleType int64 `json:"receipt_title_type,omitempty" xml:"receipt_title_type,omitempty"`
	// 代理商商家id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 原始订单号
	OriginTpOrderId int64 `json:"origin_tp_order_id,omitempty" xml:"origin_tp_order_id,omitempty"`
	// 发票金额(分)
	ReceiptAmount int64 `json:"receipt_amount,omitempty" xml:"receipt_amount,omitempty"`
}

var poolRedReceiptParam = sync.Pool{
	New: func() any {
		return new(RedReceiptParam)
	},
}

// GetRedReceiptParam() 从对象池中获取RedReceiptParam
func GetRedReceiptParam() *RedReceiptParam {
	return poolRedReceiptParam.Get().(*RedReceiptParam)
}

// ReleaseRedReceiptParam 释放RedReceiptParam
func ReleaseRedReceiptParam(v *RedReceiptParam) {
	v.ReceiptMemo = ""
	v.Receiver = ""
	v.ReceiveMail = ""
	v.OriginReceiptNumber = ""
	v.CompanyName = ""
	v.AgentOrderNo = ""
	v.ReceiptContent = ""
	v.ReceiveBankName = ""
	v.ReceiveBankAccount = ""
	v.ExtMap = ""
	v.ReceiptTitle = ""
	v.CompanyTaxNo = ""
	v.ReceiveMobile = ""
	v.ReceiptTitleType = 0
	v.AgentId = 0
	v.OriginTpOrderId = 0
	v.ReceiptAmount = 0
	poolRedReceiptParam.Put(v)
}
