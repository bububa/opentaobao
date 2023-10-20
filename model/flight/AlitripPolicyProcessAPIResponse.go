package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripPolicyProcessAPIResponse 政策进度查询 API返回值
// alitrip.policy.process
//
// 上传特殊类型的单程/往返政策
type AlitripPolicyProcessAPIResponse struct {
	model.CommonResponse
	AlitripPolicyProcessAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripPolicyProcessAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripPolicyProcessAPIResponseModel).Reset()
}

// AlitripPolicyProcessAPIResponseModel is 政策进度查询 成功返回结果
type AlitripPolicyProcessAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_policy_process_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlitripPolicyProcessResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripPolicyProcessAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripPolicyProcessAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripPolicyProcessAPIResponse)
	},
}

// GetAlitripPolicyProcessAPIResponse 从 sync.Pool 获取 AlitripPolicyProcessAPIResponse
func GetAlitripPolicyProcessAPIResponse() *AlitripPolicyProcessAPIResponse {
	return poolAlitripPolicyProcessAPIResponse.Get().(*AlitripPolicyProcessAPIResponse)
}

// ReleaseAlitripPolicyProcessAPIResponse 将 AlitripPolicyProcessAPIResponse 保存到 sync.Pool
func ReleaseAlitripPolicyProcessAPIResponse(v *AlitripPolicyProcessAPIResponse) {
	v.Reset()
	poolAlitripPolicyProcessAPIResponse.Put(v)
}
