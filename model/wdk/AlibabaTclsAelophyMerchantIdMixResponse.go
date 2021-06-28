package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家用户id混淆 APIResponse
alibaba.tcls.aelophy.merchant.id.mix

商家用户id混淆
*/
type AlibabaTclsAelophyMerchantIdMixAPIResponse struct {
    model.CommonResponse
    AlibabaTclsAelophyMerchantIdMixResponse
}

type AlibabaTclsAelophyMerchantIdMixResponse struct {
    XMLName xml.Name `xml:"alibaba_tcls_aelophy_merchant_id_mix_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    ApiResult   *AlibabaTclsAelophyMerchantIdMixApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`

    
}
