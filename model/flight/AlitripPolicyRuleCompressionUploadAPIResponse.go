package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripPolicyRuleCompressionUploadAPIResponse 大批量上传规则类型的单程/往返政策 API返回值
// alitrip.policy.rule.compression.upload
//
// 大批量上传规则类型的单程/往返政策
type AlitripPolicyRuleCompressionUploadAPIResponse struct {
	model.CommonResponse
	AlitripPolicyRuleCompressionUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripPolicyRuleCompressionUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripPolicyRuleCompressionUploadAPIResponseModel).Reset()
}

// AlitripPolicyRuleCompressionUploadAPIResponseModel is 大批量上传规则类型的单程/往返政策 成功返回结果
type AlitripPolicyRuleCompressionUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_policy_rule_compression_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripPolicyRuleCompressionUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripPolicyRuleCompressionUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripPolicyRuleCompressionUploadAPIResponse)
	},
}

// GetAlitripPolicyRuleCompressionUploadAPIResponse 从 sync.Pool 获取 AlitripPolicyRuleCompressionUploadAPIResponse
func GetAlitripPolicyRuleCompressionUploadAPIResponse() *AlitripPolicyRuleCompressionUploadAPIResponse {
	return poolAlitripPolicyRuleCompressionUploadAPIResponse.Get().(*AlitripPolicyRuleCompressionUploadAPIResponse)
}

// ReleaseAlitripPolicyRuleCompressionUploadAPIResponse 将 AlitripPolicyRuleCompressionUploadAPIResponse 保存到 sync.Pool
func ReleaseAlitripPolicyRuleCompressionUploadAPIResponse(v *AlitripPolicyRuleCompressionUploadAPIResponse) {
	v.Reset()
	poolAlitripPolicyRuleCompressionUploadAPIResponse.Put(v)
}
