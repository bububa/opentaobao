package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询计划 APIRequest
alibaba.scbp.ad.campaign.find.campaign.page

分页查询计划
*/
type AlibabaScbpAdCampaignFindCampaignPageRequest struct {
    model.Params

    // 请求实体类
    campaignQuery   *CampaignQueryDto 

    // 用户信息
    topContext   *TopContextDto 

}

func NewAlibabaScbpAdCampaignFindCampaignPageRequest() *AlibabaScbpAdCampaignFindCampaignPageRequest{
    return &AlibabaScbpAdCampaignFindCampaignPageRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdCampaignFindCampaignPageRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.campaign.find.campaign.page"
}

func (r AlibabaScbpAdCampaignFindCampaignPageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdCampaignFindCampaignPageRequest) SetCampaignQuery(campaignQuery *CampaignQueryDto) error {
    r.campaignQuery = campaignQuery
    r.Set("campaign_query", campaignQuery)
    return nil
}

func (r AlibabaScbpAdCampaignFindCampaignPageRequest) GetCampaignQuery() *CampaignQueryDto {
    return r.campaignQuery
}

func (r *AlibabaScbpAdCampaignFindCampaignPageRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdCampaignFindCampaignPageRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

