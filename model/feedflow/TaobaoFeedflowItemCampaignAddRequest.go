package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流增加推广计划 APIRequest
taobao.feedflow.item.campaign.add

信息流增加推广计划
*/
type TaobaoFeedflowItemCampaignAddRequest struct {
    model.Params

    // 计划信息
    campaign   *CampaignDto 

}

func NewTaobaoFeedflowItemCampaignAddRequest() *TaobaoFeedflowItemCampaignAddRequest{
    return &TaobaoFeedflowItemCampaignAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemCampaignAddRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.campaign.add"
}

func (r TaobaoFeedflowItemCampaignAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemCampaignAddRequest) SetCampaign(campaign *CampaignDto) error {
    r.campaign = campaign
    r.Set("campaign", campaign)
    return nil
}

func (r TaobaoFeedflowItemCampaignAddRequest) GetCampaign() *CampaignDto {
    return r.campaign
}

