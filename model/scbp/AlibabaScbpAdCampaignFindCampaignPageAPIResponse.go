package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdCampaignFindCampaignPageAPIResponse 分页查询计划 API返回值
// alibaba.scbp.ad.campaign.find.campaign.page
//
// 分页查询计划
type AlibabaScbpAdCampaignFindCampaignPageAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdCampaignFindCampaignPageAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdCampaignFindCampaignPageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdCampaignFindCampaignPageAPIResponseModel).Reset()
}

// AlibabaScbpAdCampaignFindCampaignPageAPIResponseModel is 分页查询计划 成功返回结果
type AlibabaScbpAdCampaignFindCampaignPageAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_campaign_find_campaign_page_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回数据
	ResultList []CampaignDto `json:"result_list,omitempty" xml:"result_list>campaign_dto,omitempty"`
	// 总数量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdCampaignFindCampaignPageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
	m.TotalCount = 0
}

var poolAlibabaScbpAdCampaignFindCampaignPageAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdCampaignFindCampaignPageAPIResponse)
	},
}

// GetAlibabaScbpAdCampaignFindCampaignPageAPIResponse 从 sync.Pool 获取 AlibabaScbpAdCampaignFindCampaignPageAPIResponse
func GetAlibabaScbpAdCampaignFindCampaignPageAPIResponse() *AlibabaScbpAdCampaignFindCampaignPageAPIResponse {
	return poolAlibabaScbpAdCampaignFindCampaignPageAPIResponse.Get().(*AlibabaScbpAdCampaignFindCampaignPageAPIResponse)
}

// ReleaseAlibabaScbpAdCampaignFindCampaignPageAPIResponse 将 AlibabaScbpAdCampaignFindCampaignPageAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdCampaignFindCampaignPageAPIResponse(v *AlibabaScbpAdCampaignFindCampaignPageAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdCampaignFindCampaignPageAPIResponse.Put(v)
}
