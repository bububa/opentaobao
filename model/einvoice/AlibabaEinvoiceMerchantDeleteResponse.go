package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发票中台-同平台取消授权税号适用商户 APIResponse
alibaba.einvoice.merchant.delete

税号授权给同平台下其他商户使用后，可以使用此接口取消授权，被取消授权的商户失去开票能力
*/
type AlibabaEinvoiceMerchantDeleteAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceMerchantDeleteResponse
}

type AlibabaEinvoiceMerchantDeleteResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_merchant_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 删除成功的业务平台商户ID
    
    MerchantUserId   string `json:"merchant_user_id,omitempty" xml:"merchant_user_id,omitempty"`

    
}
