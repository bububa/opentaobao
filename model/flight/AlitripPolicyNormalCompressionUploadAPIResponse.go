package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripPolicyNormalCompressionUploadAPIResponse 大批量上传普通类型的单程/往返政策 API返回值
// alitrip.policy.normal.compression.upload
//
// 大批量上传普通类型的单程/往返政策
type AlitripPolicyNormalCompressionUploadAPIResponse struct {
	model.CommonResponse
	AlitripPolicyNormalCompressionUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripPolicyNormalCompressionUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripPolicyNormalCompressionUploadAPIResponseModel).Reset()
}

// AlitripPolicyNormalCompressionUploadAPIResponseModel is 大批量上传普通类型的单程/往返政策 成功返回结果
type AlitripPolicyNormalCompressionUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_policy_normal_compression_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripPolicyNormalCompressionUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripPolicyNormalCompressionUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripPolicyNormalCompressionUploadAPIResponse)
	},
}

// GetAlitripPolicyNormalCompressionUploadAPIResponse 从 sync.Pool 获取 AlitripPolicyNormalCompressionUploadAPIResponse
func GetAlitripPolicyNormalCompressionUploadAPIResponse() *AlitripPolicyNormalCompressionUploadAPIResponse {
	return poolAlitripPolicyNormalCompressionUploadAPIResponse.Get().(*AlitripPolicyNormalCompressionUploadAPIResponse)
}

// ReleaseAlitripPolicyNormalCompressionUploadAPIResponse 将 AlitripPolicyNormalCompressionUploadAPIResponse 保存到 sync.Pool
func ReleaseAlitripPolicyNormalCompressionUploadAPIResponse(v *AlitripPolicyNormalCompressionUploadAPIResponse) {
	v.Reset()
	poolAlitripPolicyNormalCompressionUploadAPIResponse.Put(v)
}
