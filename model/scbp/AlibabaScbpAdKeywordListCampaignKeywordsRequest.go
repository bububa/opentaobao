package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取计划关键词 API请求
alibaba.scbp.ad.keyword.list.campaign.keywords

获取计划关键词
*/
type AlibabaScbpAdKeywordListCampaignKeywordsRequest struct {
    model.Params
    // 计划id
    _campaignId   int64
    // 搜索条件
    _campaignKeywordQuery   *CampaignKeywordQuery
    // 用户信息
    _topContext   *TopContextDTO
}

// 初始化AlibabaScbpAdKeywordListCampaignKeywordsRequest对象
func NewAlibabaScbpAdKeywordListCampaignKeywordsRequest() *AlibabaScbpAdKeywordListCampaignKeywordsRequest{
    return &AlibabaScbpAdKeywordListCampaignKeywordsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordListCampaignKeywordsRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.list.campaign.keywords"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordListCampaignKeywordsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdKeywordListCampaignKeywordsRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdKeywordListCampaignKeywordsRequest) GetCampaignId() int64 {
    return r._campaignId
}
// CampaignKeywordQuery Setter
// 搜索条件
func (r *AlibabaScbpAdKeywordListCampaignKeywordsRequest) SetCampaignKeywordQuery(_campaignKeywordQuery *CampaignKeywordQuery) error {
    r._campaignKeywordQuery = _campaignKeywordQuery
    r.Set("campaign_keyword_query", _campaignKeywordQuery)
    return nil
}

// CampaignKeywordQuery Getter
func (r AlibabaScbpAdKeywordListCampaignKeywordsRequest) GetCampaignKeywordQuery() *CampaignKeywordQuery {
    return r._campaignKeywordQuery
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdKeywordListCampaignKeywordsRequest) SetTopContext(_topContext *TopContextDTO) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdKeywordListCampaignKeywordsRequest) GetTopContext() *TopContextDTO {
    return r._topContext
}
