package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripInvoiceSettingModifyAPIRequest
发票变更 API请求
alitrip.btrip.invoice.setting.modify

发票变更 */
type AlitripBtripInvoiceSettingModifyAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenInvoiceModifyAndNewRq
}

// New
