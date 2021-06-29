package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改关键词价格 API请求
alibaba.scbp.ad.keyword.update.keyword.price.batch

修改关键词价格
*/
type AlibabaScbpAdKeywordUpdateKeywordPriceBatchRequest struct {
    model.Params
    // 计划id
    _campaignId   int64
    // 请求参数
    _keywordUpdateQuery   *KeywordUpdateQuery
    // 用户信息
    _topContext   *TopContextDto
}

// 初始化AlibabaScbpAdKeywordUpdateKeywordPriceBatchRequest对象
func NewAlibabaScbpAdKeywordUpdateKeywordPriceBatchRequest() *AlibabaScbpAdKeywordUpdateKeywordPriceBatchRequest{
    return &AlibabaScbpAdKeywordUpdateKeywordPriceBatchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordUpdateKeywordPriceBatchRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.update.keyword.price.batch"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordUpdateKeywordPriceBatchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdKeywordUpdateKeywordPriceBatchRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdKeywordUpdateKeywordPriceBatchRequest) GetCampaignId() int64 {
    return r._campaignId
}
// KeywordUpdateQuery Setter
// 请求参数
func (r *AlibabaScbpAdKeywordUpdateKeywordPriceBatchRequest) SetKeywordUpdateQuery(_keywordUpdateQuery *KeywordUpdateQuery) error {
    r._keywordUpdateQuery = _keywordUpdateQuery
    r.Set("keyword_update_query", _keywordUpdateQuery)
    return nil
}

// KeywordUpdateQuery Getter
func (r AlibabaScbpAdKeywordUpdateKeywordPriceBatchRequest) GetKeywordUpdateQuery() *KeywordUpdateQuery {
    return r._keywordUpdateQuery
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdKeywordUpdateKeywordPriceBatchRequest) SetTopContext(_topContext *TopContextDto) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdKeywordUpdateKeywordPriceBatchRequest) GetTopContext() *TopContextDto {
    return r._topContext
}
