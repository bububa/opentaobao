package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripPolicySpecialUploadAPIResponse 特殊政策上传 API返回值
// alitrip.policy.special.upload
//
// 上传特殊类型的单程/往返政策
type AlitripPolicySpecialUploadAPIResponse struct {
	model.CommonResponse
	AlitripPolicySpecialUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripPolicySpecialUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripPolicySpecialUploadAPIResponseModel).Reset()
}

// AlitripPolicySpecialUploadAPIResponseModel is 特殊政策上传 成功返回结果
type AlitripPolicySpecialUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_policy_special_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlitripPolicySpecialUploadResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripPolicySpecialUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripPolicySpecialUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripPolicySpecialUploadAPIResponse)
	},
}

// GetAlitripPolicySpecialUploadAPIResponse 从 sync.Pool 获取 AlitripPolicySpecialUploadAPIResponse
func GetAlitripPolicySpecialUploadAPIResponse() *AlitripPolicySpecialUploadAPIResponse {
	return poolAlitripPolicySpecialUploadAPIResponse.Get().(*AlitripPolicySpecialUploadAPIResponse)
}

// ReleaseAlitripPolicySpecialUploadAPIResponse 将 AlitripPolicySpecialUploadAPIResponse 保存到 sync.Pool
func ReleaseAlitripPolicySpecialUploadAPIResponse(v *AlitripPolicySpecialUploadAPIResponse) {
	v.Reset()
	poolAlitripPolicySpecialUploadAPIResponse.Put(v)
}
