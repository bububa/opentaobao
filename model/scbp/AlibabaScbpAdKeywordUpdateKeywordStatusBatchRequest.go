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
type AlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest struct {
    model.Params
    // 计划id
    campaignId   int64
    // 更新数据
    keywordUpdateQuery   *KeywordUpdateQuery
    // 用户信息
    topContext   *TopContextDto
}

// 初始化AlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest对象
func NewAlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest() *AlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest{
    return &AlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.update.keyword.status.batch"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest) GetCampaignId() int64 {
    return r.campaignId
}
// KeywordUpdateQuery Setter
// 更新数据
func (r *AlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest) SetKeywordUpdateQuery(keywordUpdateQuery *KeywordUpdateQuery) error {
    r.keywordUpdateQuery = keywordUpdateQuery
    r.Set("keyword_update_query", keywordUpdateQuery)
    return nil
}

// KeywordUpdateQuery Getter
func (r AlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest) GetKeywordUpdateQuery() *KeywordUpdateQuery {
    return r.keywordUpdateQuery
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest) GetTopContext() *TopContextDto {
    return r.topContext
}
