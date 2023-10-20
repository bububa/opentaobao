package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripPolicyDomfareCompareAPIResponse 比价工具 API返回值
// alitrip.policy.domfare.compare
//
// 比价工具
type AlitripPolicyDomfareCompareAPIResponse struct {
	model.CommonResponse
	AlitripPolicyDomfareCompareAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripPolicyDomfareCompareAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripPolicyDomfareCompareAPIResponseModel).Reset()
}

// AlitripPolicyDomfareCompareAPIResponseModel is 比价工具 成功返回结果
type AlitripPolicyDomfareCompareAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_policy_domfare_compare_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	Result *AlitripPolicyDomfareCompareResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripPolicyDomfareCompareAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripPolicyDomfareCompareAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripPolicyDomfareCompareAPIResponse)
	},
}

// GetAlitripPolicyDomfareCompareAPIResponse 从 sync.Pool 获取 AlitripPolicyDomfareCompareAPIResponse
func GetAlitripPolicyDomfareCompareAPIResponse() *AlitripPolicyDomfareCompareAPIResponse {
	return poolAlitripPolicyDomfareCompareAPIResponse.Get().(*AlitripPolicyDomfareCompareAPIResponse)
}

// ReleaseAlitripPolicyDomfareCompareAPIResponse 将 AlitripPolicyDomfareCompareAPIResponse 保存到 sync.Pool
func ReleaseAlitripPolicyDomfareCompareAPIResponse(v *AlitripPolicyDomfareCompareAPIResponse) {
	v.Reset()
	poolAlitripPolicyDomfareCompareAPIResponse.Put(v)
}
