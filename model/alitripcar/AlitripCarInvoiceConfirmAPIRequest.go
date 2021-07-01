package alitripcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripCarInvoiceConfirmAPIRequest
发票确认接口 API请求
alitrip.car.invoice.confirm

飞猪发票回调接口 */
type AlitripCarInvoiceConfirmAPIRequest struct {
	model.Params
	// 入参对象
	_receiptDo *ReceiptDo
}

// New
