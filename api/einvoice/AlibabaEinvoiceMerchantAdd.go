package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceMerchantAdd 发票中台-同平台授权税号适用商户
// alibaba.einvoice.merchant.add
//
// 适用于以下场景：
// 业务税号入驻成功后，需要将税号授权给同平台下其他商户，使得其他商户也具备开票能力
func AlibabaEinvoiceMerchantAdd(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceMerchantAddAPIRequest, resp *einvoice.AlibabaEinvoiceMerchantAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
