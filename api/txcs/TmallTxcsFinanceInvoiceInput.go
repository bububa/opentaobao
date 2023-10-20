package txcs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/txcs"
)

// TmallTxcsFinanceInvoiceInput 供应商发票录入
// tmall.txcs.finance.invoice.input
//
// 提供天猫超市外部合作商家财务：供应商发票录入
func TmallTxcsFinanceInvoiceInput(clt *core.SDKClient, req *txcs.TmallTxcsFinanceInvoiceInputAPIRequest, resp *txcs.TmallTxcsFinanceInvoiceInputAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
