package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripInvoiceSettingDelete 发票删除
// alitrip.btrip.invoice.setting.delete
//
// 发票删除
func AlitripBtripInvoiceSettingDelete(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripInvoiceSettingDeleteAPIRequest, resp *btrip.AlitripBtripInvoiceSettingDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
