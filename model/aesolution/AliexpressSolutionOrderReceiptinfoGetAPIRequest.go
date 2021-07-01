package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSolutionOrderReceiptinfoGetAPIRequest
Get Order Receipt Info API请求
aliexpress.solution.order.receiptinfo.get

Get Order Receipt Info, Support multi stores requirements for Turkey sellers. */
type AliexpressSolutionOrderReceiptinfoGetAPIRequest struct {
	model.Params
	// query param
	_param1 *SingleOrderQuery
}

// New
