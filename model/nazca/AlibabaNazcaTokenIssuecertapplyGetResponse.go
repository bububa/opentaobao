package nazca

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据token获取出证申请信息 APIResponse
alibaba.nazca.token.issuecertapply.get

根据token获取出证申请信息
*/
type AlibabaNazcaTokenIssuecertapplyGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaNazcaTokenIssuecertapplyGetResponse `json:"alibaba_nazca_token_issuecertapply_get_response,omitempty"` 
    AlibabaNazcaTokenIssuecertapplyGetResponse
}

/* model for simplify = false
type AlibabaNazcaTokenIssuecertapplyGetResponse struct {

    // result
    
    Result  *struct {
        ActionResult  *ActionResult `json:"action_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaNazcaTokenIssuecertapplyGetResponse struct {

    // result
    Result   *ActionResult `json:"result,omitempty"`

}
