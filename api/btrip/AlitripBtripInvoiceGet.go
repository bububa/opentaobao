package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripInvoiceGet 获取用户可用发票列表
// alitrip.btrip.invoice.get
//
// 差旅申请用户获取可用发票列表
func AlitripBtripInvoiceGet(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripInvoiceGetAPIRequest, resp *btrip.AlitripBtripInvoiceGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
