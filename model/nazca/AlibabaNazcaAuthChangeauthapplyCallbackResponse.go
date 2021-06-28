package nazca

import (
    "github.com/bububa/opentaobao/model"
)

/* 
变更认证回调 APIResponse
alibaba.nazca.auth.changeauthapply.callback

变更认证回调
*/
type AlibabaNazcaAuthChangeauthapplyCallbackAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaNazcaAuthChangeauthapplyCallbackResponse `json:"alibaba_nazca_auth_changeauthapply_callback_response,omitempty"` 
    AlibabaNazcaAuthChangeauthapplyCallbackResponse
}

/* model for simplify = false
type AlibabaNazcaAuthChangeauthapplyCallbackResponse struct {

    // result
    
    Result  *struct {
        ActionResult  *ActionResult `json:"action_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaNazcaAuthChangeauthapplyCallbackResponse struct {

    // result
    Result   *ActionResult `json:"result,omitempty"`

}
