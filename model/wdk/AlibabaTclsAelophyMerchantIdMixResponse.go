package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家用户id混淆 APIResponse
alibaba.tcls.aelophy.merchant.id.mix

商家用户id混淆
*/
type AlibabaTclsAelophyMerchantIdMixAPIResponse struct {
    model.CommonResponse
    Response *AlibabaTclsAelophyMerchantIdMixResponse `json:"alibaba_tcls_aelophy_merchant_id_mix_response,omitempty"`
}

type AlibabaTclsAelophyMerchantIdMixResponse struct {

    // 接口返回model
    ApiResult   *AlibabaTclsAelophyMerchantIdMixApiResult `json:"api_result,omitempty"`

}
