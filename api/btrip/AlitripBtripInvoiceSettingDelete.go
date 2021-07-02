package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripInvoiceSettingDelete 发票删除
// alitrip.btrip.invoice.setting.delete
//
// 发票删除
func AlitripBtripInvoiceSettingDelete(clt *core.SDKClient, req *btrip.AlitripBtripInvoiceSettingDeleteAPIRequest, session string) (*btrip.AlitripBtripInvoiceSettingDeleteAPIResponse, error) {
	var resp btrip.AlitripBtripInvoiceSettingDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
