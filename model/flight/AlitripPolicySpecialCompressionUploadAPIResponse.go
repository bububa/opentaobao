package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripPolicySpecialCompressionUploadAPIResponse 大批量上传特殊类型的单程/往返政策 API返回值
// alitrip.policy.special.compression.upload
//
// 大批量上传特殊类型的单程/往返政策
type AlitripPolicySpecialCompressionUploadAPIResponse struct {
	model.CommonResponse
	AlitripPolicySpecialCompressionUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripPolicySpecialCompressionUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripPolicySpecialCompressionUploadAPIResponseModel).Reset()
}

// AlitripPolicySpecialCompressionUploadAPIResponseModel is 大批量上传特殊类型的单程/往返政策 成功返回结果
type AlitripPolicySpecialCompressionUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_policy_special_compression_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripPolicySpecialCompressionUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripPolicySpecialCompressionUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripPolicySpecialCompressionUploadAPIResponse)
	},
}

// GetAlitripPolicySpecialCompressionUploadAPIResponse 从 sync.Pool 获取 AlitripPolicySpecialCompressionUploadAPIResponse
func GetAlitripPolicySpecialCompressionUploadAPIResponse() *AlitripPolicySpecialCompressionUploadAPIResponse {
	return poolAlitripPolicySpecialCompressionUploadAPIResponse.Get().(*AlitripPolicySpecialCompressionUploadAPIResponse)
}

// ReleaseAlitripPolicySpecialCompressionUploadAPIResponse 将 AlitripPolicySpecialCompressionUploadAPIResponse 保存到 sync.Pool
func ReleaseAlitripPolicySpecialCompressionUploadAPIResponse(v *AlitripPolicySpecialCompressionUploadAPIResponse) {
	v.Reset()
	poolAlitripPolicySpecialCompressionUploadAPIResponse.Put(v)
}
