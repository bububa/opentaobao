package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词添加 APIRequest
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

func NewAlibabaScbpAdKeywordCreateKeywordBatchRequest() *AlibabaScbpAdKeywordCreateKeywordBatchRequest{
    return &AlibabaScbpAdKeywordCreateKeywordBatchRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdKeywordCreateKeywordBatchRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.create.keyword.batch"
}

func (r AlibabaScbpAdKeywordCreateKeywordBatchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdKeywordCreateKeywordBatchRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r AlibabaScbpAdKeywordCreateKeywordBatchRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *AlibabaScbpAdKeywordCreateKeywordBatchRequest) SetKeywordQuery(keywordQuery *KeywordQuery) error {
    r.keywordQuery = keywordQuery
    r.Set("keyword_query", keywordQuery)
    return nil
}

func (r AlibabaScbpAdKeywordCreateKeywordBatchRequest) GetKeywordQuery() *KeywordQuery {
    return r.keywordQuery
}

func (r *AlibabaScbpAdKeywordCreateKeywordBatchRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdKeywordCreateKeywordBatchRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

