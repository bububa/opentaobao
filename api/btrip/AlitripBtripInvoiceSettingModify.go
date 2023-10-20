package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripInvoiceSettingModify 发票变更
// alitrip.btrip.invoice.setting.modify
//
// 发票变更
func AlitripBtripInvoiceSettingModify(clt *core.SDKClient, req *btrip.AlitripBtripInvoiceSettingModifyAPIRequest, resp *btrip.AlitripBtripInvoiceSettingModifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
