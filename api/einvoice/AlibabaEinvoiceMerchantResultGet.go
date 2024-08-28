package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceMerchantResultGet 商家自研ERP开票结果获取
// alibaba.einvoice.merchant.result.get
//
// 商家自研ERP开票结果获取
func AlibabaEinvoiceMerchantResultGet(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceMerchantResultGetAPIRequest, resp *einvoice.AlibabaEinvoiceMerchantResultGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
