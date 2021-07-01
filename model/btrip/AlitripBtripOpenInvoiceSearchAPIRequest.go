package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripOpenInvoiceSearchAPIRequest
差旅申请用户搜索可用发票列表 API请求
alitrip.btrip.open.invoice.search

差旅申请用户搜索可用发票列表 */
type AlitripBtripOpenInvoiceSearchAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenInvoiceRq
}

// New
