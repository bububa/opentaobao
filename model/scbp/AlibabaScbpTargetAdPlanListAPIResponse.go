package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdPlanListAPIResponse 定向推广-查询定向推广计划列表并返回计划基础信息 API返回值
// alibaba.scbp.target.ad.plan.list
//
// 定向推广-查询定向推广计划列表并返回计划基础信息
type AlibabaScbpTargetAdPlanListAPIResponse struct {
	model.CommonResponse
	AlibabaScbpTargetAdPlanListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpTargetAdPlanListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpTargetAdPlanListAPIResponseModel).Reset()
}

// AlibabaScbpTargetAdPlanListAPIResponseModel is 定向推广-查询定向推广计划列表并返回计划基础信息 成功返回结果
type AlibabaScbpTargetAdPlanListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_target_ad_plan_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 定向推广计划列表
	QuickCampaignList []TopP4pBasicQuickCampaignView `json:"quick_campaign_list,omitempty" xml:"quick_campaign_list>top_p4p_basic_quick_campaign_view,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// 总个数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpTargetAdPlanListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.QuickCampaignList = m.QuickCampaignList[:0]
	m.TotalPage = 0
	m.TotalNum = 0
}

var poolAlibabaScbpTargetAdPlanListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpTargetAdPlanListAPIResponse)
	},
}

// GetAlibabaScbpTargetAdPlanListAPIResponse 从 sync.Pool 获取 AlibabaScbpTargetAdPlanListAPIResponse
func GetAlibabaScbpTargetAdPlanListAPIResponse() *AlibabaScbpTargetAdPlanListAPIResponse {
	return poolAlibabaScbpTargetAdPlanListAPIResponse.Get().(*AlibabaScbpTargetAdPlanListAPIResponse)
}

// ReleaseAlibabaScbpTargetAdPlanListAPIResponse 将 AlibabaScbpTargetAdPlanListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpTargetAdPlanListAPIResponse(v *AlibabaScbpTargetAdPlanListAPIResponse) {
	v.Reset()
	poolAlibabaScbpTargetAdPlanListAPIResponse.Put(v)
}
