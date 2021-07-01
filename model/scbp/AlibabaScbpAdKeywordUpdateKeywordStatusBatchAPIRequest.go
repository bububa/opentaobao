package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改关键词状态 API请求
alibaba.scbp.ad.keyword.update.keyword.status.batch

修改关键词状态
*/
type AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIRequest struct {
    model.Params
    // 计划id
    _campaignId   int64
    // 更新数据
    _keywordUpdateQuery   *KeywordUpdateQuery
    // 用户信息
    _topContext   *TopContextDto
}

// 初始化AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIRequest对象
func NewAlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest() *AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIRequest{
    return &AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.update.keyword.status.batch"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIRequest) GetCampaignId() int64 {
    return r._campaignId
}
// KeywordUpdateQuery Setter
// 更新数据
func (r *AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIRequest) SetKeywordUpdateQuery(_keywordUpdateQuery *KeywordUpdateQuery) error {
    r._keywordUpdateQuery = _keywordUpdateQuery
    r.Set("keyword_update_query", _keywordUpdateQuery)
    return nil
}

// KeywordUpdateQuery Getter
func (r AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIRequest) GetKeywordUpdateQuery() *KeywordUpdateQuery {
    return r._keywordUpdateQuery
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIRequest) SetTopContext(_topContext *TopContextDto) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIRequest) GetTopContext() *TopContextDto {
    return r._topContext
}
