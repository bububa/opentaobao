package nazca

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNazcaAuthAuthapplyCallbackAPIResponse 认证的统一回调接口 API返回值
// alibaba.nazca.auth.authapply.callback
//
// 认证的统一回调接口
type AlibabaNazcaAuthAuthapplyCallbackAPIResponse struct {
	model.CommonResponse
	AlibabaNazcaAuthAuthapplyCallbackAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaNazcaAuthAuthapplyCallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaNazcaAuthAuthapplyCallbackAPIResponseModel).Reset()
}

// AlibabaNazcaAuthAuthapplyCallbackAPIResponseModel is 认证的统一回调接口 成功返回结果
type AlibabaNazcaAuthAuthapplyCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nazca_auth_authapply_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ActionResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaNazcaAuthAuthapplyCallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaNazcaAuthAuthapplyCallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaNazcaAuthAuthapplyCallbackAPIResponse)
	},
}

// GetAlibabaNazcaAuthAuthapplyCallbackAPIResponse 从 sync.Pool 获取 AlibabaNazcaAuthAuthapplyCallbackAPIResponse
func GetAlibabaNazcaAuthAuthapplyCallbackAPIResponse() *AlibabaNazcaAuthAuthapplyCallbackAPIResponse {
	return poolAlibabaNazcaAuthAuthapplyCallbackAPIResponse.Get().(*AlibabaNazcaAuthAuthapplyCallbackAPIResponse)
}

// ReleaseAlibabaNazcaAuthAuthapplyCallbackAPIResponse 将 AlibabaNazcaAuthAuthapplyCallbackAPIResponse 保存到 sync.Pool
func ReleaseAlibabaNazcaAuthAuthapplyCallbackAPIResponse(v *AlibabaNazcaAuthAuthapplyCallbackAPIResponse) {
	v.Reset()
	poolAlibabaNazcaAuthAuthapplyCallbackAPIResponse.Put(v)
}
