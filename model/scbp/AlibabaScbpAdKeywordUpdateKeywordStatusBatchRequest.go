package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改关键词状态 APIRequest
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

func NewAlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest() *AlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest{
    return &AlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.update.keyword.status.batch"
}

func (r AlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r AlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *AlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest) SetKeywordUpdateQuery(keywordUpdateQuery *KeywordUpdateQuery) error {
    r.keywordUpdateQuery = keywordUpdateQuery
    r.Set("keyword_update_query", keywordUpdateQuery)
    return nil
}

func (r AlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest) GetKeywordUpdateQuery() *KeywordUpdateQuery {
    return r.keywordUpdateQuery
}

func (r *AlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdKeywordUpdateKeywordStatusBatchRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

