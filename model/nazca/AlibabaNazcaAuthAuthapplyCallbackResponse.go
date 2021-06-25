package nazca

import (
    "github.com/bububa/opentaobao/model"
)

/* 
认证的统一回调接口 APIResponse
alibaba.nazca.auth.authapply.callback

认证的统一回调接口
*/
type AlibabaNazcaAuthAuthapplyCallbackAPIResponse struct {
    model.CommonResponse
    Response *AlibabaNazcaAuthAuthapplyCallbackResponse `json:"alibaba_nazca_auth_authapply_callback_response,omitempty"`
}

type AlibabaNazcaAuthAuthapplyCallbackResponse struct {

    // result
    Result   *ActionResult `json:"result,omitempty"`

}
