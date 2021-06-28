package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
saas 售后逆向 商户发起逆向取货 APIResponse
alibaba.tcls.aelophy.refund.fetchgoods

saas 售后逆向 商户发起逆向取货
*/
type AlibabaTclsAelophyRefundFetchgoodsAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaTclsAelophyRefundFetchgoodsResponse `json:"alibaba_tcls_aelophy_refund_fetchgoods_response,omitempty"` 
    AlibabaTclsAelophyRefundFetchgoodsResponse
}

/* model for simplify = false
type AlibabaTclsAelophyRefundFetchgoodsResponse struct {

    // 出参
    
    Result  *struct {
        AlibabaTclsAelophyRefundFetchgoodsApiResult  *AlibabaTclsAelophyRefundFetchgoodsApiResult `json:"alibaba_tcls_aelophy_refund_fetchgoods_api_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaTclsAelophyRefundFetchgoodsResponse struct {

    // 出参
    Result   *AlibabaTclsAelophyRefundFetchgoodsApiResult `json:"result,omitempty"`

}
