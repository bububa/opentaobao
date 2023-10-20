package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripInvoiceSettingAdd 发票设置
// alitrip.btrip.invoice.setting.add
//
// 发票设置
func AlitripBtripInvoiceSettingAdd(clt *core.SDKClient, req *btrip.AlitripBtripInvoiceSettingAddAPIRequest, resp *btrip.AlitripBtripInvoiceSettingAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
