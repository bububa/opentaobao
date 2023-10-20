package nazca

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNazcaTokenChangeauthapplyGetAPIResponse 根据token获取变更认证申请信息 API返回值
// alibaba.nazca.token.changeauthapply.get
//
// 根据token获取变更认证申请信息
type AlibabaNazcaTokenChangeauthapplyGetAPIResponse struct {
	model.CommonResponse
	AlibabaNazcaTokenChangeauthapplyGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaNazcaTokenChangeauthapplyGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaNazcaTokenChangeauthapplyGetAPIResponseModel).Reset()
}

// AlibabaNazcaTokenChangeauthapplyGetAPIResponseModel is 根据token获取变更认证申请信息 成功返回结果
type AlibabaNazcaTokenChangeauthapplyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nazca_token_changeauthapply_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ActionResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaNazcaTokenChangeauthapplyGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaNazcaTokenChangeauthapplyGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaNazcaTokenChangeauthapplyGetAPIResponse)
	},
}

// GetAlibabaNazcaTokenChangeauthapplyGetAPIResponse 从 sync.Pool 获取 AlibabaNazcaTokenChangeauthapplyGetAPIResponse
func GetAlibabaNazcaTokenChangeauthapplyGetAPIResponse() *AlibabaNazcaTokenChangeauthapplyGetAPIResponse {
	return poolAlibabaNazcaTokenChangeauthapplyGetAPIResponse.Get().(*AlibabaNazcaTokenChangeauthapplyGetAPIResponse)
}

// ReleaseAlibabaNazcaTokenChangeauthapplyGetAPIResponse 将 AlibabaNazcaTokenChangeauthapplyGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaNazcaTokenChangeauthapplyGetAPIResponse(v *AlibabaNazcaTokenChangeauthapplyGetAPIResponse) {
	v.Reset()
	poolAlibabaNazcaTokenChangeauthapplyGetAPIResponse.Put(v)
}
