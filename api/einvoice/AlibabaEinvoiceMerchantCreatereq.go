package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceMerchantCreatereq 商家自研ERP开票请求接口
// alibaba.einvoice.merchant.createreq
//
// 商家自研ERP发起开票请求，无需授权，API只能使用商家入驻的税号进行开票
func AlibabaEinvoiceMerchantCreatereq(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceMerchantCreatereqAPIRequest, resp *einvoice.AlibabaEinvoiceMerchantCreatereqAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
