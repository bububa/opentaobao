package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdPlanTagGetAPIResponse 定向推广-获取计划的定向溢价数据 API返回值
// alibaba.scbp.target.ad.plan.tag.get
//
// 定向推广-获取计划的定向溢价数据
type AlibabaScbpTargetAdPlanTagGetAPIResponse struct {
	model.CommonResponse
	AlibabaScbpTargetAdPlanTagGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpTargetAdPlanTagGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpTargetAdPlanTagGetAPIResponseModel).Reset()
}

// AlibabaScbpTargetAdPlanTagGetAPIResponseModel is 定向推广-获取计划的定向溢价数据 成功返回结果
type AlibabaScbpTargetAdPlanTagGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_target_ad_plan_tag_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// TopP4pCampaignTargetingTagView
	Result *TopP4pCampaignTargetingTagView `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpTargetAdPlanTagGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaScbpTargetAdPlanTagGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpTargetAdPlanTagGetAPIResponse)
	},
}

// GetAlibabaScbpTargetAdPlanTagGetAPIResponse 从 sync.Pool 获取 AlibabaScbpTargetAdPlanTagGetAPIResponse
func GetAlibabaScbpTargetAdPlanTagGetAPIResponse() *AlibabaScbpTargetAdPlanTagGetAPIResponse {
	return poolAlibabaScbpTargetAdPlanTagGetAPIResponse.Get().(*AlibabaScbpTargetAdPlanTagGetAPIResponse)
}

// ReleaseAlibabaScbpTargetAdPlanTagGetAPIResponse 将 AlibabaScbpTargetAdPlanTagGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpTargetAdPlanTagGetAPIResponse(v *AlibabaScbpTargetAdPlanTagGetAPIResponse) {
	v.Reset()
	poolAlibabaScbpTargetAdPlanTagGetAPIResponse.Put(v)
}
