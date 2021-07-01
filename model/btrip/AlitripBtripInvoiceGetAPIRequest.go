package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripInvoiceGetAPIRequest
获取用户可用发票列表 API请求
alitrip.btrip.invoice.get

差旅申请用户获取可用发票列表 */
type AlitripBtripInvoiceGetAPIRequest struct {
	model.Params
	// 企业id
	_corpId string
	// 用户id
	_userId string
}

// New
