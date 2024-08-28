package alitripreceipt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripreceipt"
)

// AlitripReceiptSellerInvoiceRed 飞猪发票冲红
// alitrip.receipt.seller.invoice.red
//
// 飞猪发票创建
func AlitripReceiptSellerInvoiceRed(ctx context.Context, clt *core.SDKClient, req *alitripreceipt.AlitripReceiptSellerInvoiceRedAPIRequest, resp *alitripreceipt.AlitripReceiptSellerInvoiceRedAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
