package nazca

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
变更认证回调 API返回值 
alibaba.nazca.auth.changeauthapply.callback

变更认证回调
*/
type AlibabaNazcaAuthChangeauthapplyCallbackAPIResponse struct {
    model.CommonResponse
    AlibabaNazcaAuthChangeauthapplyCallbackAPIResponseModel
}

// 变更认证回调 成功返回结果
type AlibabaNazcaAuthChangeauthapplyCallbackAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_nazca_auth_changeauthapply_callback_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ActionResult `json:"result,omitempty" xml:"result,omitempty"`
}
