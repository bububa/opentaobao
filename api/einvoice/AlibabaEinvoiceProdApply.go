package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceProdApply 提交发票申请
// alibaba.einvoice.prod.apply
//
// 提交开票申请，如果商户授权自动开票则自动转开票，否则等待商户审核。
func AlibabaEinvoiceProdApply(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceProdApplyAPIRequest, resp *einvoice.AlibabaEinvoiceProdApplyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
