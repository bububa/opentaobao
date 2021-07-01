package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripInvoiceSettingRuleAPIRequest
发票规则设置 API请求
alitrip.btrip.invoice.setting.rule

发票规则设置 */
type AlitripBtripInvoiceSettingRuleAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenInvoiceRuleRq
}

// New
