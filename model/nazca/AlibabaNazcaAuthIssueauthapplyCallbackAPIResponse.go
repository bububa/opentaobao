package nazca

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabanazcaauthissueauthapplycallbackAPIResponse 出证申请回调 API返回值
// alibaba.nazca.auth.issueauthapply.callback
//
// 出证申请回调
type AlibabanazcaauthissueauthapplycallbackAPIResponse struct {
	model.CommonResponse
	AlibabanazcaauthissueauthapplycallbackAPIResponseModel
}

// AlibabanazcaauthissueauthapplycallbackAPIResponseModel is 出证申请回调 成功返回结果
type AlibabanazcaauthissueauthapplycallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nazca_auth_issueauthapply_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ActionResult `json:"result,omitempty" xml:"result,omitempty"`
}
