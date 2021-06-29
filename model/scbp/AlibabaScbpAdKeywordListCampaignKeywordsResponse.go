package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取计划关键词 API返回值 
alibaba.scbp.ad.keyword.list.campaign.keywords

获取计划关键词
*/
type AlibabaScbpAdKeywordListCampaignKeywordsAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdKeywordListCampaignKeywordsResponse
}

// 获取计划关键词 成功返回结果
type AlibabaScbpAdKeywordListCampaignKeywordsResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_list_campaign_keywords_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回数据集合
    ResultList   []KeywordDto `json:"result_list,omitempty" xml:"result_list>keyword_dto,omitempty"`
}
