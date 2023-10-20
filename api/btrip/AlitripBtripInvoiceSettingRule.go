package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripInvoiceSettingRule 发票规则设置
// alitrip.btrip.invoice.setting.rule
//
// 发票规则设置
func AlitripBtripInvoiceSettingRule(clt *core.SDKClient, req *btrip.AlitripBtripInvoiceSettingRuleAPIRequest, resp *btrip.AlitripBtripInvoiceSettingRuleAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
