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
type AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest struct {
    model.Params
    // 计划id
    _campaignId   int64
    // 请求参数
    _keywordQuery   *KeywordQuery
    // 用户信息
    _topContext   *TopContextDto
}

// 初始化AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest对象
func NewAlibabaScbpAdKeywordDeleteKeywordBatchRequest() *AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest{
    return &AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.delete.keyword.batch"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest) GetCampaignId() int64 {
    return r._campaignId
}
// KeywordQuery Setter
// 请求参数
func (r *AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest) SetKeywordQuery(_keywordQuery *KeywordQuery) error {
    r._keywordQuery = _keywordQuery
    r.Set("keyword_query", _keywordQuery)
    return nil
}

// KeywordQuery Getter
func (r AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest) GetKeywordQuery() *KeywordQuery {
    return r._keywordQuery
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest) SetTopContext(_topContext *TopContextDto) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest) GetTopContext() *TopContextDto {
    return r._topContext
}
