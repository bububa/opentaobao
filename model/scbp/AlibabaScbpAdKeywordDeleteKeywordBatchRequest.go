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
    campaignId   int64
    // 请求参数
    keywordQuery   *KeywordQuery
    // 用户信息
    topContext   *TopContextDto
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
func (r *AlibabaScbpAdKeywordDeleteKeywordBatchRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdKeywordDeleteKeywordBatchRequest) GetCampaignId() int64 {
    return r.campaignId
}
// KeywordQuery Setter
// 请求参数
func (r *AlibabaScbpAdKeywordDeleteKeywordBatchRequest) SetKeywordQuery(keywordQuery *KeywordQuery) error {
    r.keywordQuery = keywordQuery
    r.Set("keyword_query", keywordQuery)
    return nil
}

// KeywordQuery Getter
func (r AlibabaScbpAdKeywordDeleteKeywordBatchRequest) GetKeywordQuery() *KeywordQuery {
    return r.keywordQuery
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdKeywordDeleteKeywordBatchRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdKeywordDeleteKeywordBatchRequest) GetTopContext() *TopContextDto {
    return r.topContext
}
