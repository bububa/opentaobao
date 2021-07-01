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
type AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest struct {
    model.Params
    // 用户信息
    _topContext   *TopContextDto
    // 请求参数
    _campaignKeywordQuery   *CampaignKeywordQuery
}

// 初始化AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest对象
func NewAlibabaScbpAdKeywordGetKeywordCountByQueryRequest() *AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest{
    return &AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.get.keyword.count.by.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest) SetTopContext(_topContext *TopContextDto) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest) GetTopContext() *TopContextDto {
    return r._topContext
}
// CampaignKeywordQuery Setter
// 请求参数
func (r *AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest) SetCampaignKeywordQuery(_campaignKeywordQuery *CampaignKeywordQuery) error {
    r._campaignKeywordQuery = _campaignKeywordQuery
    r.Set("campaign_keyword_query", _campaignKeywordQuery)
    return nil
}

// CampaignKeywordQuery Getter
func (r AlibabaScbpAdKeywordGetKeywordCountByQueryAPIRequest) GetCampaignKeywordQuery() *CampaignKeywordQuery {
    return r._campaignKeywordQuery
}
