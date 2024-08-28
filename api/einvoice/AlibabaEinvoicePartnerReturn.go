package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoicePartnerReturn 开票商回传开票结果
// alibaba.einvoice.partner.return
//
// 开票商返回开票结果数据
func AlibabaEinvoicePartnerReturn(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoicePartnerReturnAPIRequest, resp *einvoice.AlibabaEinvoicePartnerReturnAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
