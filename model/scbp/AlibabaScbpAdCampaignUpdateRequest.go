package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改计划 APIRequest
alibaba.scbp.ad.campaign.update

修改计划
*/
type AlibabaScbpAdCampaignUpdateRequest struct {
    model.Params

    // 修改数据
    campaignOperation   *CampaignOperationDto 

    // 用户信息
    topContext   *TopContextDto 

}

func NewAlibabaScbpAdCampaignUpdateRequest() *AlibabaScbpAdCampaignUpdateRequest{
    return &AlibabaScbpAdCampaignUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdCampaignUpdateRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.campaign.update"
}

func (r AlibabaScbpAdCampaignUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdCampaignUpdateRequest) SetCampaignOperation(campaignOperation *CampaignOperationDto) error {
    r.campaignOperation = campaignOperation
    r.Set("campaign_operation", campaignOperation)
    return nil
}

func (r AlibabaScbpAdCampaignUpdateRequest) GetCampaignOperation() *CampaignOperationDto {
    return r.campaignOperation
}

func (r *AlibabaScbpAdCampaignUpdateRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdCampaignUpdateRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

