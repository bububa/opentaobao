package nazca

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNazcaAuthChangeauthapplyCallbackAPIResponse 变更认证回调 API返回值
// alibaba.nazca.auth.changeauthapply.callback
//
// 变更认证回调
type AlibabaNazcaAuthChangeauthapplyCallbackAPIResponse struct {
	model.CommonResponse
	AlibabaNazcaAuthChangeauthapplyCallbackAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaNazcaAuthChangeauthapplyCallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaNazcaAuthChangeauthapplyCallbackAPIResponseModel).Reset()
}

// AlibabaNazcaAuthChangeauthapplyCallbackAPIResponseModel is 变更认证回调 成功返回结果
type AlibabaNazcaAuthChangeauthapplyCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nazca_auth_changeauthapply_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ActionResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaNazcaAuthChangeauthapplyCallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaNazcaAuthChangeauthapplyCallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaNazcaAuthChangeauthapplyCallbackAPIResponse)
	},
}

// GetAlibabaNazcaAuthChangeauthapplyCallbackAPIResponse 从 sync.Pool 获取 AlibabaNazcaAuthChangeauthapplyCallbackAPIResponse
func GetAlibabaNazcaAuthChangeauthapplyCallbackAPIResponse() *AlibabaNazcaAuthChangeauthapplyCallbackAPIResponse {
	return poolAlibabaNazcaAuthChangeauthapplyCallbackAPIResponse.Get().(*AlibabaNazcaAuthChangeauthapplyCallbackAPIResponse)
}

// ReleaseAlibabaNazcaAuthChangeauthapplyCallbackAPIResponse 将 AlibabaNazcaAuthChangeauthapplyCallbackAPIResponse 保存到 sync.Pool
func ReleaseAlibabaNazcaAuthChangeauthapplyCallbackAPIResponse(v *AlibabaNazcaAuthChangeauthapplyCallbackAPIResponse) {
	v.Reset()
	poolAlibabaNazcaAuthChangeauthapplyCallbackAPIResponse.Put(v)
}
