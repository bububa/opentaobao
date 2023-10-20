package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdPlanAddAPIResponse 定向推广-新建计划 API返回值
// alibaba.scbp.target.ad.plan.add
//
// 定向推广-新建单条计划
type AlibabaScbpTargetAdPlanAddAPIResponse struct {
	model.CommonResponse
	AlibabaScbpTargetAdPlanAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpTargetAdPlanAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpTargetAdPlanAddAPIResponseModel).Reset()
}

// AlibabaScbpTargetAdPlanAddAPIResponseModel is 定向推广-新建计划 成功返回结果
type AlibabaScbpTargetAdPlanAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_target_ad_plan_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 计划ID
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpTargetAdPlanAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CampaignId = 0
}

var poolAlibabaScbpTargetAdPlanAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpTargetAdPlanAddAPIResponse)
	},
}

// GetAlibabaScbpTargetAdPlanAddAPIResponse 从 sync.Pool 获取 AlibabaScbpTargetAdPlanAddAPIResponse
func GetAlibabaScbpTargetAdPlanAddAPIResponse() *AlibabaScbpTargetAdPlanAddAPIResponse {
	return poolAlibabaScbpTargetAdPlanAddAPIResponse.Get().(*AlibabaScbpTargetAdPlanAddAPIResponse)
}

// ReleaseAlibabaScbpTargetAdPlanAddAPIResponse 将 AlibabaScbpTargetAdPlanAddAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpTargetAdPlanAddAPIResponse(v *AlibabaScbpTargetAdPlanAddAPIResponse) {
	v.Reset()
	poolAlibabaScbpTargetAdPlanAddAPIResponse.Put(v)
}
