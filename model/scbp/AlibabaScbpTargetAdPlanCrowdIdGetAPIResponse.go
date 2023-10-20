package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdPlanCrowdIdGetAPIResponse 定向推广-人群标签ID获取(店铺老客、优选人群) API返回值
// alibaba.scbp.target.ad.plan.crowd.id.get
//
// 定向推广-人群标签ID获取(店铺老客、优选人群)
type AlibabaScbpTargetAdPlanCrowdIdGetAPIResponse struct {
	model.CommonResponse
	AlibabaScbpTargetAdPlanCrowdIdGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpTargetAdPlanCrowdIdGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpTargetAdPlanCrowdIdGetAPIResponseModel).Reset()
}

// AlibabaScbpTargetAdPlanCrowdIdGetAPIResponseModel is 定向推广-人群标签ID获取(店铺老客、优选人群) 成功返回结果
type AlibabaScbpTargetAdPlanCrowdIdGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_target_ad_plan_crowd_id_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果list
	ResultList []CrowdView `json:"result_list,omitempty" xml:"result_list>crowd_view,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpTargetAdPlanCrowdIdGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
}

var poolAlibabaScbpTargetAdPlanCrowdIdGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpTargetAdPlanCrowdIdGetAPIResponse)
	},
}

// GetAlibabaScbpTargetAdPlanCrowdIdGetAPIResponse 从 sync.Pool 获取 AlibabaScbpTargetAdPlanCrowdIdGetAPIResponse
func GetAlibabaScbpTargetAdPlanCrowdIdGetAPIResponse() *AlibabaScbpTargetAdPlanCrowdIdGetAPIResponse {
	return poolAlibabaScbpTargetAdPlanCrowdIdGetAPIResponse.Get().(*AlibabaScbpTargetAdPlanCrowdIdGetAPIResponse)
}

// ReleaseAlibabaScbpTargetAdPlanCrowdIdGetAPIResponse 将 AlibabaScbpTargetAdPlanCrowdIdGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpTargetAdPlanCrowdIdGetAPIResponse(v *AlibabaScbpTargetAdPlanCrowdIdGetAPIResponse) {
	v.Reset()
	poolAlibabaScbpTargetAdPlanCrowdIdGetAPIResponse.Put(v)
}
