package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除关键词 API请求
alibaba.scbp.ad.keyword.delete.keyword.batch

删除关键词
*/
type AlibabaScbpAdKeywordDeleteKeywordBatchRequest struct {
    model.Params
    // 计划id
    _campaignId   int64
    // 请求参数
    _keywordQuery   *KeywordQuery
    // 用户信息
    _topContext   *TopContextDTO
}

// 初始化AlibabaScbpAdKeywordDeleteKeywordBatchRequest对象
func NewAlibabaScbpAdKeywordDeleteKeywordBatchRequest() *AlibabaScbpAdKeywordDeleteKeywordBatchRequest{
    return &AlibabaScbpAdKeywordDeleteKeywordBatchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordDeleteKeywordBatchRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.delete.keyword.batch"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordDeleteKeywordBatchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdKeywordDeleteKeywordBatchRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdKeywordDeleteKeywordBatchRequest) GetCampaignId() int64 {
    return r._campaignId
}
// KeywordQuery Setter
// 请求参数
func (r *AlibabaScbpAdKeywordDeleteKeywordBatchRequest) SetKeywordQuery(_keywordQuery *KeywordQuery) error {
    r._keywordQuery = _keywordQuery
    r.Set("keyword_query", _keywordQuery)
    return nil
}

// KeywordQuery Getter
func (r AlibabaScbpAdKeywordDeleteKeywordBatchRequest) GetKeywordQuery() *KeywordQuery {
    return r._keywordQuery
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdKeywordDeleteKeywordBatchRequest) SetTopContext(_topContext *TopContextDTO) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdKeywordDeleteKeywordBatchRequest) GetTopContext() *TopContextDTO {
    return r._topContext
}
