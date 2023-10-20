package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicemerchantbindcompany 发票中台-跨平台绑定已入驻税号与商户
// alibaba.einvoice.merchant.bindcompany
//
// 税号在阿里发票平台入驻成功后，允许业务方通过本接口跨业务平台绑定入驻税号和业务平台商户，绑定成功后该商户可以使用该税号的盘进行开票。绑定成功后，可以使用同平台授权、取消授权税号适用商户接口来变更税号和商户关系。
func Alibabaeinvoicemerchantbindcompany(clt *core.SDKClient, req *einvoice.AlibabaeinvoicemerchantbindcompanyAPIRequest, session string) (*einvoice.AlibabaeinvoicemerchantbindcompanyAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicemerchantbindcompanyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
