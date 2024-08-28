package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripInvoiceSettingModify 发票变更
// alitrip.btrip.invoice.setting.modify
//
// 发票变更
func AlitripBtripInvoiceSettingModify(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripInvoiceSettingModifyAPIRequest, resp *btrip.AlitripBtripInvoiceSettingModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
