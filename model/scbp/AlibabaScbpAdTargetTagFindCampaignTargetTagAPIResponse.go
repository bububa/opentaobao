package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdTargetTagFindCampaignTargetTagAPIResponse 查询标签数据 API返回值
// alibaba.scbp.ad.target.tag.find.campaign.target.tag
//
// 查询标签数据
type AlibabaScbpAdTargetTagFindCampaignTargetTagAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdTargetTagFindCampaignTargetTagAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdTargetTagFindCampaignTargetTagAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdTargetTagFindCampaignTargetTagAPIResponseModel).Reset()
}

// AlibabaScbpAdTargetTagFindCampaignTargetTagAPIResponseModel is 查询标签数据 成功返回结果
type AlibabaScbpAdTargetTagFindCampaignTargetTagAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_target_tag_find_campaign_target_tag_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回实体
	ResultList []AdsTargetingTagDto `json:"result_list,omitempty" xml:"result_list>ads_targeting_tag_dto,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdTargetTagFindCampaignTargetTagAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
}

var poolAlibabaScbpAdTargetTagFindCampaignTargetTagAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdTargetTagFindCampaignTargetTagAPIResponse)
	},
}

// GetAlibabaScbpAdTargetTagFindCampaignTargetTagAPIResponse 从 sync.Pool 获取 AlibabaScbpAdTargetTagFindCampaignTargetTagAPIResponse
func GetAlibabaScbpAdTargetTagFindCampaignTargetTagAPIResponse() *AlibabaScbpAdTargetTagFindCampaignTargetTagAPIResponse {
	return poolAlibabaScbpAdTargetTagFindCampaignTargetTagAPIResponse.Get().(*AlibabaScbpAdTargetTagFindCampaignTargetTagAPIResponse)
}

// ReleaseAlibabaScbpAdTargetTagFindCampaignTargetTagAPIResponse 将 AlibabaScbpAdTargetTagFindCampaignTargetTagAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdTargetTagFindCampaignTargetTagAPIResponse(v *AlibabaScbpAdTargetTagFindCampaignTargetTagAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdTargetTagFindCampaignTargetTagAPIResponse.Put(v)
}
