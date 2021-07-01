package alitripreceipt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripReceiptSellerInvoiceQueryAPIRequest
飞猪发票查询 API请求
alitrip.receipt.seller.invoice.query

飞猪发票查询 */
type AlitripReceiptSellerInvoiceQueryAPIRequest struct {
	model.Params
	// 入参对象
	_queryReceiptParam *QueryReceiptParam
}

// New
