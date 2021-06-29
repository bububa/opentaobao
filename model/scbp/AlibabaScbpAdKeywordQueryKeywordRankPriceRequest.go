package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询关键词前五名排价 API请求
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

// 初始化AlibabaScbpAdKeywordQueryKeywordRankPriceRequest对象
func NewAlibabaScbpAdKeywordQueryKeywordRankPriceRequest() *AlibabaScbpAdKeywordQueryKeywordRankPriceRequest{
    return &AlibabaScbpAdKeywordQueryKeywordRankPriceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordQueryKeywordRankPriceRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.query.keyword.rank.price"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordQueryKeywordRankPriceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdKeywordQueryKeywordRankPriceRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdKeywordQueryKeywordRankPriceRequest) GetCampaignId() int64 {
    return r.campaignId
}
// KeywordQuery Setter
// 请求参数
func (r *AlibabaScbpAdKeywordQueryKeywordRankPriceRequest) SetKeywordQuery(keywordQuery *KeywordQuery) error {
    r.keywordQuery = keywordQuery
    r.Set("keyword_query", keywordQuery)
    return nil
}

// KeywordQuery Getter
func (r AlibabaScbpAdKeywordQueryKeywordRankPriceRequest) GetKeywordQuery() *KeywordQuery {
    return r.keywordQuery
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdKeywordQueryKeywordRankPriceRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdKeywordQueryKeywordRankPriceRequest) GetTopContext() *TopContextDto {
    return r.topContext
}
