package nazca

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
变更认证回调 APIResponse
alibaba.nazca.auth.changeauthapply.callback

变更认证回调
*/
type AlibabaNazcaAuthChangeauthapplyCallbackAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_nazca_auth_changeauthapply_callback_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *ActionResult `json:"result,omitempty" xml:"