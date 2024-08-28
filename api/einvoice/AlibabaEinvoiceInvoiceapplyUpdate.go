package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceInvoiceapplyUpdate 商家开票申请单状态回传
// alibaba.einvoice.invoiceapply.update
//
// 开票服务商更新商家开票申请单状态
func AlibabaEinvoiceInvoiceapplyUpdate(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceInvoiceapplyUpdateAPIRequest, resp *einvoice.AlibabaEinvoiceInvoiceapplyUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
