package nazca

import (
    "github.com/bububa/opentaobao/model"
)

/* 
出证申请回调 APIResponse
alibaba.nazca.auth.issueauthapply.callback

出证申请回调
*/
type AlibabaNazcaAuthIssueauthapplyCallbackAPIResponse struct {
    model.CommonResponse
    Response *AlibabaNazcaAuthIssueauthapplyCallbackResponse `json:"alibaba_nazca_auth_issueauthapply_callback_response,omitempty"`
}

type AlibabaNazcaAuthIssueauthapplyCallbackResponse struct {

    // result
    Result   *ActionResult `json:"result,omitempty"`

}
