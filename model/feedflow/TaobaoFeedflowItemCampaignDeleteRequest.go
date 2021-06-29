package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除计划 APIRequest
taobao.feedflow.item.campaign.delete

删除计划
*/
type TaobaoFeedflowItemCampaignDeleteRequest struct {
    model.Params

    // 计划id
    campaignId   int64 

}

func NewTaobaoFeedflowItemCampaignDeleteRequest() *TaobaoFeedflowItemCampaignDeleteRequest{
    return &TaobaoFeedflowItemCampaignDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemCampaignDeleteRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.campaign.delete"
}

func (r TaobaoFeedflowItemCampaignDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemCampaignDeleteRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoFeedflowItemCampaignDeleteRequest) GetCampaignId() int64 {
    return r.campaignId
}

