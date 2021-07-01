package alitripreceipt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripReceiptSellerInvoiceRedAPIRequest
飞猪发票冲红 API请求
alitrip.receipt.seller.invoice.red

飞猪发票创建 */
type AlitripReceiptSellerInvoiceRedAPIRequest struct {
	model.Params
	// 入参对象
	_redReceiptParam *RedReceiptParam
}

// New
