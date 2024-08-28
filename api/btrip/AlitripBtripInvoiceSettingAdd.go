package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripInvoiceSettingAdd 发票设置
// alitrip.btrip.invoice.setting.add
//
// 发票设置
func AlitripBtripInvoiceSettingAdd(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripInvoiceSettingAddAPIRequest, resp *btrip.AlitripBtripInvoiceSettingAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
