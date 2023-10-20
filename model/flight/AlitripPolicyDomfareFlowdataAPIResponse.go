package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripPolicyDomfareFlowdataAPIResponse 比价工具流量详情 API返回值
// alitrip.policy.domfare.flowdata
//
// 比价工具流量详情
type AlitripPolicyDomfareFlowdataAPIResponse struct {
	model.CommonResponse
	AlitripPolicyDomfareFlowdataAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripPolicyDomfareFlowdataAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripPolicyDomfareFlowdataAPIResponseModel).Reset()
}

// AlitripPolicyDomfareFlowdataAPIResponseModel is 比价工具流量详情 成功返回结果
type AlitripPolicyDomfareFlowdataAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_policy_domfare_flowdata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	Result *AlitripPolicyDomfareFlowdataResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripPolicyDomfareFlowdataAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripPolicyDomfareFlowdataAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripPolicyDomfareFlowdataAPIResponse)
	},
}

// GetAlitripPolicyDomfareFlowdataAPIResponse 从 sync.Pool 获取 AlitripPolicyDomfareFlowdataAPIResponse
func GetAlitripPolicyDomfareFlowdataAPIResponse() *AlitripPolicyDomfareFlowdataAPIResponse {
	return poolAlitripPolicyDomfareFlowdataAPIResponse.Get().(*AlitripPolicyDomfareFlowdataAPIResponse)
}

// ReleaseAlitripPolicyDomfareFlowdataAPIResponse 将 AlitripPolicyDomfareFlowdataAPIResponse 保存到 sync.Pool
func ReleaseAlitripPolicyDomfareFlowdataAPIResponse(v *AlitripPolicyDomfareFlowdataAPIResponse) {
	v.Reset()
	poolAlitripPolicyDomfareFlowdataAPIResponse.Put(v)
}
