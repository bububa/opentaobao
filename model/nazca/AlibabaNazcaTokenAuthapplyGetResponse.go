package nazca

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据token获取认证申请信息 APIResponse
alibaba.nazca.token.authapply.get

根据token获取认证申请信息
*/
type AlibabaNazcaTokenAuthapplyGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaNazcaTokenAuthapplyGetResponse `json:"alibaba_nazca_token_authapply_get_response,omitempty"`
}

type AlibabaNazcaTokenAuthapplyGetResponse struct {

    // result
    Result   *ActionResult `json:"result,omitempty"`

}
