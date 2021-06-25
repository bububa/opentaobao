package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取计划关键词 APIRequest
alibaba.scbp.ad.keyword.list.campaign.keywords

获取计划关键词
*/
type AlibabaScbpAdKeywordListCampaignKeywordsRequest struct {
    model.Params

    // 计划id
    campaignId   int64 

    // 搜索条件
    campaignKeywordQuery   *CampaignKeywordQuery 

    // 用户信息
    topContext   *TopContextDto 

}

func NewAlibabaScbpAdKeywordListCampaignKeywordsRequest() *AlibabaScbpAdKeywordListCampaignKeywordsRequest{
    return &AlibabaScbpAdKeywordListCampaignKeywordsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdKeywordListCampaignKeywordsRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.list.campaign.keywords"
}

func (r AlibabaScbpAdKeywordListCampaignKeywordsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdKeywordListCampaignKeywordsRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r AlibabaScbpAdKeywordListCampaignKeywordsRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *AlibabaScbpAdKeywordListCampaignKeywordsRequest) SetCampaignKeywordQuery(campaignKeywordQuery *CampaignKeywordQuery) error {
    r.campaignKeywordQuery = campaignKeywordQuery
    r.Set("campaign_keyword_query", campaignKeywordQuery)
    return nil
}

func (r AlibabaScbpAdKeywordListCampaignKeywordsRequest) GetCampaignKeywordQuery() *CampaignKeywordQuery {
    return r.campaignKeywordQuery
}

func (r *AlibabaScbpAdKeywordListCampaignKeywordsRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdKeywordListCampaignKeywordsRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

