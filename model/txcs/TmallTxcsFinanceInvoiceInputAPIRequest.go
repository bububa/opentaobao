package txcs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallTxcsFinanceInvoiceInputAPIRequest
供应商发票录入 API请求
tmall.txcs.finance.invoice.input

提供天猫超市外部合作商家财务：供应商发票录入 */
type TmallTxcsFinanceInvoiceInputAPIRequest struct {
	model.Params
	// 门店ID
	_ouCode string
	// 发票内容
	_invoiceInputDTO1 []InvoiceInputDto
}

// New
