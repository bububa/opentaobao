package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

/* AlitripBtripInvoiceSettingAdd
发票设置
alitrip.btrip.invoice.setting.add

发票设置 */
func AlitripBtripInvoiceSettingAdd(clt *core.SDKClient, req *btrip.AlitripBtripInvoiceSettingAddAPIRequest, session string) (*btrip.AlitripBtripInvoiceSettingAddAPIResponse, error) {
	var resp btrip.AlitripBtripInvoiceSettingAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
