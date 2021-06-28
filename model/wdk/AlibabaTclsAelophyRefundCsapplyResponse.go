package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家代客售后提交逆向申请 APIResponse
alibaba.tcls.aelophy.refund.csapply

商家代客售后提交逆向申请
*/
type AlibabaTclsAelophyRefundCsapplyAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaTclsAelophyRefundCsapplyResponse `json:"alibaba_tcls_aelophy_refund_csapply_response,omitempty"` 
    AlibabaTclsAelophyRefundCsapplyResponse
}

/* model for simplify = false
type AlibabaTclsAelophyRefundCsapplyResponse struct {

    // 根据站点名称查询产品
    
    ApiResult  *struct {
        AlibabaTclsAelophyRefundCsapplyApiResult  *AlibabaTclsAelophyRefundCsapplyApiResult `json:"alibaba_tcls_aelophy_refund_csapply_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type AlibabaTclsAelophyRefundCsapplyResponse struct {

    // 根据站点名称查询产品
    ApiResult   *AlibabaTclsAelophyRefundCsapplyApiResult `json:"api_result,omitempty"`

}
