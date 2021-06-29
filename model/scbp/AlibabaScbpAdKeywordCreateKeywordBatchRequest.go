package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词添加 API请求
alibaba.scbp.ad.keyword.create.keyword.batch

关键词添加
*/
type AlibabaScbpAdKeywordCreateKeywordBatchRequest struct {
    model.Params
    // 计划id
    campaignId   int64
    // 请求参数
    keywordQuery   *KeywordQuery
    // 用户信息
    topContext   *TopContextDto
}

// 初始化AlibabaScbpAdKeywordCreateKeywordBatchRequest对象
func NewAlibabaScbpAdKeywordCreateKeywordBatchRequest() *AlibabaScbpAdKeywordCreateKeywordBatchRequest{
    return &AlibabaScbpAdKeywordCreateKeywordBatchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordCreateKeywordBatchRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.create.keyword.batch"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordCreateKeywordBatchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdKeywordCreateKeywordBatchRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdKeywordCreateKeywordBatchRequest) GetCampaignId() int64 {
    return r.campaignId
}
// KeywordQuery Setter
// 请求参数
func (r *AlibabaScbpAdKeywordCreateKeywordBatchRequest) SetKeywordQuery(keywordQuery *KeywordQuery) error {
    r.keywordQuery = keywordQuery
    r.Set("keyword_query", keywordQuery)
    return nil
}

// KeywordQuery Getter
func (r AlibabaScbpAdKeywordCreateKeywordBatchRequest) GetKeywordQuery() *KeywordQuery {
    return r.keywordQuery
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdKeywordCreateKeywordBatchRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdKeywordCreateKeywordBatchRequest) GetTopContext() *TopContextDto {
    return r.topContext
}
