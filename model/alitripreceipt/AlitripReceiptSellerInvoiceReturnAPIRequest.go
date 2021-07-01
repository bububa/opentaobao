package alitripreceipt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripReceiptSellerInvoiceReturnAPIRequest
飞猪发票商家回调接口 API请求
alitrip.receipt.seller.invoice.return

飞猪发票回调接口 */
type AlitripReceiptSellerInvoiceReturnAPIRequest struct {
	model.Params
	// 入参对象
	_receiptDo *ReceiptDo
}

// New
