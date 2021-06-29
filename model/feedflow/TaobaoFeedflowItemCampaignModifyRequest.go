package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流修改计划 APIRequest
taobao.feedflow.item.campaign.modify

信息流修改计划
*/
type TaobaoFeedflowItemCampaignModifyRequest struct {
    model.Params

    // 修改参数
    campaign   *CampaignDto 

}

func NewTaobaoFeedflowItemCampaignModifyRequest() *TaobaoFeedflowItemCampaignModifyRequest{
    return &TaobaoFeedflowItemCampaignModifyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemCampaignModifyRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.campaign.modify"
}

func (r TaobaoFeedflowItemCampaignModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemCampaignModifyRequest) SetCampaign(campaign *CampaignDto) error {
    r.campaign = campaign
    r.Set("campaign", campaign)
    return nil
}

func (r TaobaoFeedflowItemCampaignModifyRequest) GetCampaign() *CampaignDto {
    return r.campaign
}

