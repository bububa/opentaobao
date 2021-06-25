package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取计划关键词 APIResponse
alibaba.scbp.ad.keyword.list.campaign.keywords

获取计划关键词
*/
type AlibabaScbpAdKeywordListCampaignKeywordsAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAdKeywordListCampaignKeywordsResponse `json:"alibaba_scbp_ad_keyword_list_campaign_keywords_response,omitempty"`
}

type AlibabaScbpAdKeywordListCampaignKeywordsResponse struct {

    // 返回数据集合
    ResultList   []KeywordDto `json:"result_list,omitempty"`

}
