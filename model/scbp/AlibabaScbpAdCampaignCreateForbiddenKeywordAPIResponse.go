package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponse 创建屏蔽词 API返回值
// alibaba.scbp.ad.campaign.create.forbidden.keyword
//
// 创建屏蔽词
type AlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponseModel).Reset()
}

// AlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponseModel is 创建屏蔽词 成功返回结果
type AlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_campaign_create_forbidden_keyword_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponse)
	},
}

// GetAlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponse 从 sync.Pool 获取 AlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponse
func GetAlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponse() *AlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponse {
	return poolAlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponse.Get().(*AlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponse)
}

// ReleaseAlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponse 将 AlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponse(v *AlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponse.Put(v)
}
