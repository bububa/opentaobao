package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripInvoiceSettingAddAPIRequest
发票设置 API请求
alitrip.btrip.invoice.setting.add

发票设置 */
type AlitripBtripInvoiceSettingAddAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenInvoiceModifyAndNewRq
}

// New
