package alitripreceipt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripreceipt"
)

/* AlitripReceiptSellerInvoiceQuery
飞猪发票查询
alitrip.receipt.seller.invoice.query

飞猪发票查询 */
func AlitripReceiptSellerInvoiceQuery(clt *core.SDKClient, req *alitripreceipt.AlitripReceiptSellerInvoiceQueryAPIRequest, session string) (*alitripreceipt.AlitripReceiptSellerInvoiceQueryAPIResponse, error) {
	var resp alitripreceipt.AlitripReceiptSellerInvoiceQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
