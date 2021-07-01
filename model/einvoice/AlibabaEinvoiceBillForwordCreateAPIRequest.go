package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceBillForwordCreateAPIRequest
扫码开票结算单同步前开发票 API请求
alibaba.einvoice.bill.forword.create

扫码开票结算单同步前开发票，会将数据同步到结算单中 */
type AlibabaEinvoiceBillForwordCreateAPIRequest struct {
	model.Params
	// 发票商品明细
	_invoiceItems []BillItemDo
	// 结算单订单ID
	_orderId string
	// 收款方税号
	_payeeRegisterNo string
	// 调用平台，用以区分不同的订单ID，不填默认为default
	_platform string
	// 开票流水号，若不填则系统默认生成
	_seriNo string
	// 店铺名称
	_shopName string
	// 开票金额，和明细累计总金额需相同
	_sumPrice string
	// 付款方地址
	_payerAddress string
	// 付款方银行账户
	_payerBankaccount string
	// 付款方开票邮件通知邮箱
	_payerEmail string
	// 付款方发票抬头
	_payerName string
	// 企业电话
	_payerPhone string
	// 付款方税号
	_payerRegisterNo string
	// 个人电话，接收发票通知
	_phoneNumber string
	// 企业或个人抬头发票，0=个人，1=企业。默认=1
	_businessType int64
}

// New
