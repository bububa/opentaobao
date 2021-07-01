package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripInvoiceSearchAPIRequest
根据发票抬头搜索发票 API请求
alitrip.btrip.invoice.search

用户根据发票抬头搜索发票信息 */
type AlitripBtripInvoiceSearchAPIRequest struct {
	model.Params
	// 企业id
	_corpId string
	// 用户id
	_userId string
	// 发票抬头
	_title string
}

// New
