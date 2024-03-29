package alscmerchant

import (
	"sync"
)

// MerchantTicketUseDetail 结构体
type MerchantTicketUseDetail struct {
	// 券核销流水号
	TicketTransId string `json:"ticket_trans_id,omitempty" xml:"ticket_trans_id,omitempty"`
	// 核销的券码
	TicketCode string `json:"ticket_code,omitempty" xml:"ticket_code,omitempty"`
	// 用户购买券的时候实际支付的金额，单位为元，精确到小数点后两位 /***** 20220922后对接的服务商，不支持该字段 ******
	BuyerPayAmount string `json:"buyer_pay_amount,omitempty" xml:"buyer_pay_amount,omitempty"`
	// 商家实收金额，单位为元，精确到小数点后两位 /***** 20220922后对接的服务商，不支持该字段 ******
	ReceiptAmount string `json:"receipt_amount,omitempty" xml:"receipt_amount,omitempty"`
	// 商家优惠金额，单位为元，精确到小数点后两位 /***** 20220922后对接的服务商，不支持该字段 ******
	DiscountAmount string `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	// 口碑补贴金额，单位为元，精确到小数点后两位 /***** 20220922后对接的服务商，不支持该字段 ******
	MerchantSubsidyAmount string `json:"merchant_subsidy_amount,omitempty" xml:"merchant_subsidy_amount,omitempty"`
	// 交易中可给用户开具发票的金额，单位为元，精确到小数点后两位 /***** 20220922后对接的服务商，不支持该字段 ******
	InvoiceAmount string `json:"invoice_amount,omitempty" xml:"invoice_amount,omitempty"`
	// 对应的凭证id
	TicketId string `json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
}

var poolMerchantTicketUseDetail = sync.Pool{
	New: func() any {
		return new(MerchantTicketUseDetail)
	},
}

// GetMerchantTicketUseDetail() 从对象池中获取MerchantTicketUseDetail
func GetMerchantTicketUseDetail() *MerchantTicketUseDetail {
	return poolMerchantTicketUseDetail.Get().(*MerchantTicketUseDetail)
}

// ReleaseMerchantTicketUseDetail 释放MerchantTicketUseDetail
func ReleaseMerchantTicketUseDetail(v *MerchantTicketUseDetail) {
	v.TicketTransId = ""
	v.TicketCode = ""
	v.BuyerPayAmount = ""
	v.ReceiptAmount = ""
	v.DiscountAmount = ""
	v.MerchantSubsidyAmount = ""
	v.InvoiceAmount = ""
	v.TicketId = ""
	poolMerchantTicketUseDetail.Put(v)
}
