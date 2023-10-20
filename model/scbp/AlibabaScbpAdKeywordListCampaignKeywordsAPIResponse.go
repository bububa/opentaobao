package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordListCampaignKeywordsAPIResponse 获取计划关键词 API返回值
// alibaba.scbp.ad.keyword.list.campaign.keywords
//
// 获取计划关键词
type AlibabaScbpAdKeywordListCampaignKeywordsAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdKeywordListCampaignKeywordsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordListCampaignKeywordsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdKeywordListCampaignKeywordsAPIResponseModel).Reset()
}

// AlibabaScbpAdKeywordListCampaignKeywordsAPIResponseModel is 获取计划关键词 成功返回结果
type AlibabaScbpAdKeywordListCampaignKeywordsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_list_campaign_keywords_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回数据集合
	ResultList []KeywordDto `json:"result_list,omitempty" xml:"result_list>keyword_dto,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordListCampaignKeywordsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
}

var poolAlibabaScbpAdKeywordListCampaignKeywordsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdKeywordListCampaignKeywordsAPIResponse)
	},
}

// GetAlibabaScbpAdKeywordListCampaignKeywordsAPIResponse 从 sync.Pool 获取 AlibabaScbpAdKeywordListCampaignKeywordsAPIResponse
func GetAlibabaScbpAdKeywordListCampaignKeywordsAPIResponse() *AlibabaScbpAdKeywordListCampaignKeywordsAPIResponse {
	return poolAlibabaScbpAdKeywordListCampaignKeywordsAPIResponse.Get().(*AlibabaScbpAdKeywordListCampaignKeywordsAPIResponse)
}

// ReleaseAlibabaScbpAdKeywordListCampaignKeywordsAPIResponse 将 AlibabaScbpAdKeywordListCampaignKeywordsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdKeywordListCampaignKeywordsAPIResponse(v *AlibabaScbpAdKeywordListCampaignKeywordsAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdKeywordListCampaignKeywordsAPIResponse.Put(v)
}
