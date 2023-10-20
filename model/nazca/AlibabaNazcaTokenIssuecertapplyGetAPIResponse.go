package nazca

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNazcaTokenIssuecertapplyGetAPIResponse 根据token获取出证申请信息 API返回值
// alibaba.nazca.token.issuecertapply.get
//
// 根据token获取出证申请信息
type AlibabaNazcaTokenIssuecertapplyGetAPIResponse struct {
	model.CommonResponse
	AlibabaNazcaTokenIssuecertapplyGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaNazcaTokenIssuecertapplyGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaNazcaTokenIssuecertapplyGetAPIResponseModel).Reset()
}

// AlibabaNazcaTokenIssuecertapplyGetAPIResponseModel is 根据token获取出证申请信息 成功返回结果
type AlibabaNazcaTokenIssuecertapplyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nazca_token_issuecertapply_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ActionResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaNazcaTokenIssuecertapplyGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaNazcaTokenIssuecertapplyGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaNazcaTokenIssuecertapplyGetAPIResponse)
	},
}

// GetAlibabaNazcaTokenIssuecertapplyGetAPIResponse 从 sync.Pool 获取 AlibabaNazcaTokenIssuecertapplyGetAPIResponse
func GetAlibabaNazcaTokenIssuecertapplyGetAPIResponse() *AlibabaNazcaTokenIssuecertapplyGetAPIResponse {
	return poolAlibabaNazcaTokenIssuecertapplyGetAPIResponse.Get().(*AlibabaNazcaTokenIssuecertapplyGetAPIResponse)
}

// ReleaseAlibabaNazcaTokenIssuecertapplyGetAPIResponse 将 AlibabaNazcaTokenIssuecertapplyGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaNazcaTokenIssuecertapplyGetAPIResponse(v *AlibabaNazcaTokenIssuecertapplyGetAPIResponse) {
	v.Reset()
	poolAlibabaNazcaTokenIssuecertapplyGetAPIResponse.Put(v)
}
