package einvoice

import (
	"sync"
)

// Apply 结构体
type Apply struct {
	// 发票明细
	InvoiceItems []InvoiceItem `json:"invoice_items,omitempty" xml:"invoice_items>invoice_item,omitempty"`
	// 电商平台代码,TB,TM,ALIPAY,JD
	PlatformCode string `json:"platform_code,omitempty" xml:"platform_code,omitempty"`
	// 买家备注
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 买家抬头
	PayerName string `json:"payer_name,omitempty" xml:"payer_name,omitempty"`
	// 电商平台对应的订单号
	PlatformTid string `json:"platform_tid,omitempty" xml:"platform_tid,omitempty"`
	// 买家税号
	PayerRegisterNo string `json:"payer_register_no,omitempty" xml:"payer_register_no,omitempty"`
	// 开票申请的触发类型，buyer_payed=卖家已付款，sent_goods=卖家已发货，buyer_confirm=买家确认收货，refund_seller_confirm=卖家同意退款，invoice_supply=买家申请补开发票，invoice_change=买家申请改抬头，change_paper=电换纸
	TriggerStatus string `json:"trigger_status,omitempty" xml:"trigger_status,omitempty"`
	// 发票(开票)类型，蓝票blue,红票red，默认blue
	InvoiceType string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	// 开票金额
	InvoiceAmount string `json:"invoice_amount,omitempty" xml:"invoice_amount,omitempty"`
	// 不含税总金额
	SumPrice string `json:"sum_price,omitempty" xml:"sum_price,omitempty"`
	// 总税额
	SumTax string `json:"sum_tax,omitempty" xml:"sum_tax,omitempty"`
	// 购买方联系电话
	PayerPhone string `json:"payer_phone,omitempty" xml:"payer_phone,omitempty"`
	// 购买方地址
	PayerAddress string `json:"payer_address,omitempty" xml:"payer_address,omitempty"`
	// 购买方开户行账号
	PayerBankaccount string `json:"payer_bankaccount,omitempty" xml:"payer_bankaccount,omitempty"`
	// 购买方开户银行
	PayerBank string `json:"payer_bank,omitempty" xml:"payer_bank,omitempty"`
	// 发票申请修改时间
	GmtModifiedStr string `json:"gmt_modified_str,omitempty" xml:"gmt_modified_str,omitempty"`
	// 扩展属性
	ExtendProps string `json:"extend_props,omitempty" xml:"extend_props,omitempty"`
	// 发票申请创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 开票申请状态，0=已拒绝，1=申请中，2=已同意
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 发票种类，0=电子发票，1=纸质发票，2=专票，3=电子专用发票，4=全电普通发票，5=全电专用发票
	InvoiceKind int64 `json:"invoice_kind,omitempty" xml:"invoice_kind,omitempty"`
	// 抬头类型，0=个人，1=企业
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
}

var poolApply = sync.Pool{
	New: func() any {
		return new(Apply)
	},
}

// GetApply() 从对象池中获取Apply
func GetApply() *Apply {
	return poolApply.Get().(*Apply)
}

// ReleaseApply 释放Apply
func ReleaseApply(v *Apply) {
	v.InvoiceItems = v.InvoiceItems[:0]
	v.PlatformCode = ""
	v.Memo = ""
	v.PayerName = ""
	v.PlatformTid = ""
	v.PayerRegisterNo = ""
	v.TriggerStatus = ""
	v.InvoiceType = ""
	v.InvoiceAmount = ""
	v.SumPrice = ""
	v.SumTax = ""
	v.PayerPhone = ""
	v.PayerAddress = ""
	v.PayerBankaccount = ""
	v.PayerBank = ""
	v.GmtModifiedStr = ""
	v.ExtendProps = ""
	v.GmtCreate = ""
	v.Status = 0
	v.InvoiceKind = 0
	v.BusinessType = 0
	poolApply.Put(v)
}
