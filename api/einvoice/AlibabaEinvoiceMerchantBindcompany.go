package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceMerchantBindcompany 发票中台-跨平台绑定已入驻税号与商户
// alibaba.einvoice.merchant.bindcompany
//
// 税号在阿里发票平台入驻成功后，允许业务方通过本接口跨业务平台绑定入驻税号和业务平台商户，绑定成功后该商户可以使用该税号的盘进行开票。绑定成功后，可以使用同平台授权、取消授权税号适用商户接口来变更税号和商户关系。
func AlibabaEinvoiceMerchantBindcompany(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceMerchantBindcompanyAPIRequest, resp *einvoice.AlibabaEinvoiceMerchantBindcompanyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
