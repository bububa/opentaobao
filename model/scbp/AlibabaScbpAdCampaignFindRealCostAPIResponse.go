package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdCampaignFindRealCostAPIResponse 批量查询计划消耗数据 API返回值
// alibaba.scbp.ad.campaign.find.real.cost
//
// 批量查询计划消耗数据
type AlibabaScbpAdCampaignFindRealCostAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdCampaignFindRealCostAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdCampaignFindRealCostAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdCampaignFindRealCostAPIResponseModel).Reset()
}

// AlibabaScbpAdCampaignFindRealCostAPIResponseModel is 批量查询计划消耗数据 成功返回结果
type AlibabaScbpAdCampaignFindRealCostAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_campaign_find_real_cost_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回数据结果，json数据，key是campaignId,value是消耗数据信息
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdCampaignFindRealCostAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaScbpAdCampaignFindRealCostAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdCampaignFindRealCostAPIResponse)
	},
}

// GetAlibabaScbpAdCampaignFindRealCostAPIResponse 从 sync.Pool 获取 AlibabaScbpAdCampaignFindRealCostAPIResponse
func GetAlibabaScbpAdCampaignFindRealCostAPIResponse() *AlibabaScbpAdCampaignFindRealCostAPIResponse {
	return poolAlibabaScbpAdCampaignFindRealCostAPIResponse.Get().(*AlibabaScbpAdCampaignFindRealCostAPIResponse)
}

// ReleaseAlibabaScbpAdCampaignFindRealCostAPIResponse 将 AlibabaScbpAdCampaignFindRealCostAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdCampaignFindRealCostAPIResponse(v *AlibabaScbpAdCampaignFindRealCostAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdCampaignFindRealCostAPIResponse.Put(v)
}
