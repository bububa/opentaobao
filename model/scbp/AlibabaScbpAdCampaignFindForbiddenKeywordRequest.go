package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询屏蔽词 APIRequest
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

func NewAlibabaScbpAdCampaignFindForbiddenKeywordRequest() *AlibabaScbpAdCampaignFindForbiddenKeywordRequest{
    return &AlibabaScbpAdCampaignFindForbiddenKeywordRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdCampaignFindForbiddenKeywordRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.campaign.find.forbidden.keyword"
}

func (r AlibabaScbpAdCampaignFindForbiddenKeywordRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdCampaignFindForbiddenKeywordRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r AlibabaScbpAdCampaignFindForbiddenKeywordRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *AlibabaScbpAdCampaignFindForbiddenKeywordRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdCampaignFindForbiddenKeywordRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

