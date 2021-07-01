package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发票中台-跨平台绑定已入驻税号与商户 API返回值 
alibaba.einvoice.merchant.bindcompany

税号在阿里发票平台入驻成功后，允许业务方通过本接口跨业务平台绑定入驻税号和业务平台商户，绑定成功后该商户可以使用该税号的盘进行开票。绑定成功后，可以使用同平台授权、取消授权税号适用商户接口来变更税号和商户关系。
*/
type AlibabaEinvoiceMerchantBindcompanyAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceMerchantBindcompanyAPIResponseModel
}

// 发票中台-跨平台绑定已入驻税号与商户 成功返回结果
type AlibabaEinvoiceMerchantBindcompanyAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_einvoice_merchant_bindcompany_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // token，此token用于税号适用门店新增和删除接口，需要业务方保存
    TaxToken   string `json:"tax_token,omitempty" xml:"tax_token,omitempty"`
}
