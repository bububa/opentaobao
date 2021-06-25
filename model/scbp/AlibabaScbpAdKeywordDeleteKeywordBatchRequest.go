package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除关键词 APIRequest
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

func NewAlibabaScbpAdKeywordDeleteKeywordBatchRequest() *AlibabaScbpAdKeywordDeleteKeywordBatchRequest{
    return &AlibabaScbpAdKeywordDeleteKeywordBatchRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdKeywordDeleteKeywordBatchRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.delete.keyword.batch"
}

func (r AlibabaScbpAdKeywordDeleteKeywordBatchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdKeywordDeleteKeywordBatchRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r AlibabaScbpAdKeywordDeleteKeywordBatchRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *AlibabaScbpAdKeywordDeleteKeywordBatchRequest) SetKeywordQuery(keywordQuery *KeywordQuery) error {
    r.keywordQuery = keywordQuery
    r.Set("keyword_query", keywordQuery)
    return nil
}

func (r AlibabaScbpAdKeywordDeleteKeywordBatchRequest) GetKeywordQuery() *KeywordQuery {
    return r.keywordQuery
}

func (r *AlibabaScbpAdKeywordDeleteKeywordBatchRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdKeywordDeleteKeywordBatchRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

