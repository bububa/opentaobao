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
type AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest struct {
    model.Params
    // 计划id
    _campaignId   int64
    // 请求参数
    _keywordQuery   *KeywordQuery
    // 用户信息
    _topContext   *TopContextDto
}

// 初始化AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest对象
func NewAlibabaScbpAdKeywordQueryKeywordRankPriceRequest() *AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest{
    return &AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.query.keyword.rank.price"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) GetCampaignId() int64 {
    return r._campaignId
}
// KeywordQuery Setter
// 请求参数
func (r *AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) SetKeywordQuery(_keywordQuery *KeywordQuery) error {
    r._keywordQuery = _keywordQuery
    r.Set("keyword_query", _keywordQuery)
    return nil
}

// KeywordQuery Getter
func (r AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) GetKeywordQuery() *KeywordQuery {
    return r._keywordQuery
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) SetTopContext(_topContext *TopContextDto) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdKeywordQueryKeywordRankPriceAPIRequest) GetTopContext() *TopContextDto {
    return r._topContext
}
