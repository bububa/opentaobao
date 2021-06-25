package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询关键词前五名排价 APIRequest
alibaba.scbp.ad.keyword.query.keyword.rank.price

查询关键词前五名排价
*/
type AlibabaScbpAdKeywordQueryKeywordRankPriceRequest struct {
    model.Params

    // 计划id
    campaignId   int64 

    // 请求参数
    keywordQuery   *KeywordQuery 

    // 用户信息
    topContext   *TopContextDto 

}

func NewAlibabaScbpAdKeywordQueryKeywordRankPriceRequest() *AlibabaScbpAdKeywordQueryKeywordRankPriceRequest{
    return &AlibabaScbpAdKeywordQueryKeywordRankPriceRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdKeywordQueryKeywordRankPriceRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.query.keyword.rank.price"
}

func (r AlibabaScbpAdKeywordQueryKeywordRankPriceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdKeywordQueryKeywordRankPriceRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r AlibabaScbpAdKeywordQueryKeywordRankPriceRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *AlibabaScbpAdKeywordQueryKeywordRankPriceRequest) SetKeywordQuery(keywordQuery *KeywordQuery) error {
    r.keywordQuery = keywordQuery
    r.Set("keyword_query", keywordQuery)
    return nil
}

func (r AlibabaScbpAdKeywordQueryKeywordRankPriceRequest) GetKeywordQuery() *KeywordQuery {
    return r.keywordQuery
}

func (r *AlibabaScbpAdKeywordQueryKeywordRankPriceRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdKeywordQueryKeywordRankPriceRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

