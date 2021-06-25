package nazca

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据token获取变更认证申请信息 APIResponse
alibaba.nazca.token.changeauthapply.get

根据token获取变更认证申请信息
*/
type AlibabaNazcaTokenChangeauthapplyGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaNazcaTokenChangeauthapplyGetResponse `json:"alibaba_nazca_token_changeauthapply_get_response,omitempty"`
}

type AlibabaNazcaTokenChangeauthapplyGetResponse struct {

    // result
    Result   *ActionResult `json:"result,omitempty"`

}
