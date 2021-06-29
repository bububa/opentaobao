package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
计划关键词数目 API请求
alibaba.scbp.ad.keyword.get.keyword.count.by.query

计划关键词数目
*/
type AlibabaScbpAdKeywordGetKeywordCountByQueryRequest struct {
    model.Params
    // 用户信息
    topContext   *TopContextDto
    // 请求参数
    campaignKeywordQuery   *CampaignKeywordQuery
}

// 初始化AlibabaScbpAdKeywordGetKeywordCountByQueryRequest对象
func NewAlibabaScbpAdKeywordGetKeywordCountByQueryRequest() *AlibabaScbpAdKeywordGetKeywordCountByQueryRequest{
    return &AlibabaScbpAdKeywordGetKeywordCountByQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordGetKeywordCountByQueryRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.get.keyword.count.by.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordGetKeywordCountByQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdKeywordGetKeywordCountByQueryRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdKeywordGetKeywordCountByQueryRequest) GetTopContext() *TopContextDto {
    return r.topContext
}
// CampaignKeywordQuery Setter
// 请求参数
func (r *AlibabaScbpAdKeywordGetKeywordCountByQueryRequest) SetCampaignKeywordQuery(campaignKeywordQuery *CampaignKeywordQuery) error {
    r.campaignKeywordQuery = campaignKeywordQuery
    r.Set("campaign_keyword_query", campaignKeywordQuery)
    return nil
}

// CampaignKeywordQuery Getter
func (r AlibabaScbpAdKeywordGetKeywordCountByQueryRequest) GetCampaignKeywordQuery() *CampaignKeywordQuery {
    return r.campaignKeywordQuery
}
