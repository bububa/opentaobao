package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripInvoiceSettingDelete 发票删除
// alitrip.btrip.invoice.setting.delete
//
// 发票删除
func AlitripBtripInvoiceSettingDelete(clt *core.SDKClient, req *btrip.AlitripBtripInvoiceSettingDeleteAPIRequest, resp *btrip.AlitripBtripInvoiceSettingDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
