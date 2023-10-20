package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdCampaignFindForbiddenKeywordAPIResponse 查询屏蔽词 API返回值
// alibaba.scbp.ad.campaign.find.forbidden.keyword
//
// 查询屏蔽词
type AlibabaScbpAdCampaignFindForbiddenKeywordAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdCampaignFindForbiddenKeywordAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdCampaignFindForbiddenKeywordAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdCampaignFindForbiddenKeywordAPIResponseModel).Reset()
}

// AlibabaScbpAdCampaignFindForbiddenKeywordAPIResponseModel is 查询屏蔽词 成功返回结果
type AlibabaScbpAdCampaignFindForbiddenKeywordAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_campaign_find_forbidden_keyword_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回数据
	Result *AlibabaScbpAdCampaignFindForbiddenKeywordResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdCampaignFindForbiddenKeywordAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaScbpAdCampaignFindForbiddenKeywordAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdCampaignFindForbiddenKeywordAPIResponse)
	},
}

// GetAlibabaScbpAdCampaignFindForbiddenKeywordAPIResponse 从 sync.Pool 获取 AlibabaScbpAdCampaignFindForbiddenKeywordAPIResponse
func GetAlibabaScbpAdCampaignFindForbiddenKeywordAPIResponse() *AlibabaScbpAdCampaignFindForbiddenKeywordAPIResponse {
	return poolAlibabaScbpAdCampaignFindForbiddenKeywordAPIResponse.Get().(*AlibabaScbpAdCampaignFindForbiddenKeywordAPIResponse)
}

// ReleaseAlibabaScbpAdCampaignFindForbiddenKeywordAPIResponse 将 AlibabaScbpAdCampaignFindForbiddenKeywordAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdCampaignFindForbiddenKeywordAPIResponse(v *AlibabaScbpAdCampaignFindForbiddenKeywordAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdCampaignFindForbiddenKeywordAPIResponse.Put(v)
}
