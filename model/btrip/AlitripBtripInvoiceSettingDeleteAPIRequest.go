package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripInvoiceSettingDeleteAPIRequest
发票删除 API请求
alitrip.btrip.invoice.setting.delete

发票删除 */
type AlitripBtripInvoiceSettingDeleteAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenInvoiceDeleteRq
}

// New
