package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceCreatereq ERP开票请求接口
// alibaba.einvoice.createreq
//
// ERP发起开票请求
func AlibabaEinvoiceCreatereq(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceCreatereqAPIRequest, resp *einvoice.AlibabaEinvoiceCreatereqAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
