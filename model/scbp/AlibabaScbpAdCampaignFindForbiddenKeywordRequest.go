package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询屏蔽词 API请求
alibaba.scbp.ad.campaign.find.forbidden.keyword

查询屏蔽词
*/
type AlibabaScbpAdCampaignFindForbiddenKeywordRequest struct {
    model.Params
    // 计划id
    campaignId   int64
    // 用户信息
    topContext   *TopContextDto
}

// 初始化AlibabaScbpAdCampaignFindForbiddenKeywordRequest对象
func NewAlibabaScbpAdCampaignFindForbiddenKeywordRequest() *AlibabaScbpAdCampaignFindForbiddenKeywordRequest{
    return &AlibabaScbpAdCampaignFindForbiddenKeywordRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCampaignFindForbiddenKeywordRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.campaign.find.forbidden.keyword"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCampaignFindForbiddenKeywordRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdCampaignFindForbiddenKeywordRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdCampaignFindForbiddenKeywordRequest) GetCampaignId() int64 {
    return r.campaignId
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignFindForbiddenKeywordRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdCampaignFindForbiddenKeywordRequest) GetTopContext() *TopContextDto {
    return r.topContext
}
