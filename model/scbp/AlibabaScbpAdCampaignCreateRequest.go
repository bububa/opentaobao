package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建计划 APIRequest
alibaba.scbp.ad.campaign.create

创建计划
*/
type AlibabaScbpAdCampaignCreateRequest struct {
    model.Params

    // 返回数据
    campaignOperation   *CampaignOperationDto 

    // 用户信息
    topContext   *TopContextDto 

}

func NewAlibabaScbpAdCampaignCreateRequest() *AlibabaScbpAdCampaignCreateRequest{
    return &AlibabaScbpAdCampaignCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdCampaignCreateRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.campaign.create"
}

func (r AlibabaScbpAdCampaignCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdCampaignCreateRequest) SetCampaignOperation(campaignOperation *CampaignOperationDto) error {
    r.campaignOperation = campaignOperation
    r.Set("campaign_operation", campaignOperation)
    return nil
}

func (r AlibabaScbpAdCampaignCreateRequest) GetCampaignOperation() *CampaignOperationDto {
    return r.campaignOperation
}

func (r *AlibabaScbpAdCampaignCreateRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdCampaignCreateRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

