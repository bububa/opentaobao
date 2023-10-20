package alitripreceipt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripreceipt"
)

// AlitripReceiptSellerInvoiceRed 飞猪发票冲红
// alitrip.receipt.seller.invoice.red
//
// 飞猪发票创建
func AlitripReceiptSellerInvoiceRed(clt *core.SDKClient, req *alitripreceipt.AlitripReceiptSellerInvoiceRedAPIRequest, resp *alitripreceipt.AlitripReceiptSellerInvoiceRedAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
