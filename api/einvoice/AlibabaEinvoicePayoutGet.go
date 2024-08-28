package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoicePayoutGet 获取赔付计时列表数据
// alibaba.einvoice.payout.get
//
// 获取赔付计时列表数据
func AlibabaEinvoicePayoutGet(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoicePayoutGetAPIRequest, resp *einvoice.AlibabaEinvoicePayoutGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
