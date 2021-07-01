package nazca

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNazcaAuthIssueauthapplyCallbackAPIResponse
出证申请回调 API返回值
alibaba.nazca.auth.issueauthapply.callback

出证申请回调 */
type AlibabaNazcaAuthIssueauthapplyCallbackAPIResponse struct {
	model.CommonResponse
	AlibabaNazcaAuthIssueauthapplyCallbackAPIResponseModel
}

// AlibabaNazcaAuthIssueauthapplyCallbackAPIResponseModel is 出证申请回调 成功返回结果
type AlibabaNazcaAuthIssueauthapplyCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nazca_auth_issueauthapply_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ActionResult `json:"result,omitempty" xml:"result,omitempty"`
}
