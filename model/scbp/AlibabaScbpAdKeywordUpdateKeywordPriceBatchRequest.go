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
type AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest struct {
    model.Params
    // 计划id
    _campaignId   int64
    // 请求参数
    _keywordUpdateQuery   *KeywordUpdateQuery
    // 用户信息
    _topContext   *TopContextDTO
}

// 初始化AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest对象
func NewAlibabaScbpAdKeywordUpdateKeywordPriceBatchRequest() *AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest{
    return &AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.update.keyword.price.batch"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest) GetCampaignId() int64 {
    return r._campaignId
}
// KeywordUpdateQuery Setter
// 请求参数
func (r *AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest) SetKeywordUpdateQuery(_keywordUpdateQuery *KeywordUpdateQuery) error {
    r._keywordUpdateQuery = _keywordUpdateQuery
    r.Set("keyword_update_query", _keywordUpdateQuery)
    return nil
}

// KeywordUpdateQuery Getter
func (r AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest) GetKeywordUpdateQuery() *KeywordUpdateQuery {
    return r._keywordUpdateQuery
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest) SetTopContext(_topContext *TopContextDTO) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIRequest) GetTopContext() *TopContextDTO {
    return r._topContext
}
