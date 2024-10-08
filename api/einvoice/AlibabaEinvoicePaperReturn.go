package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoicePaperReturn 纸质发票结果回传
// alibaba.einvoice.paper.return
//
// 纸质发票结果回传
func AlibabaEinvoicePaperReturn(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoicePaperReturnAPIRequest, resp *einvoice.AlibabaEinvoicePaperReturnAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
