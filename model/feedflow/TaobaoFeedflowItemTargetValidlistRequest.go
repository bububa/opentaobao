package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取有权限的定向列表 APIRequest
taobao.feedflow.item.target.validlist

获取有权限的定向列表
*/
type TaobaoFeedflowItemTargetValidlistRequest struct {
    model.Params

    // 计划id
    campaignId   int64 

}

func NewTaobaoFeedflowItemTargetValidlistRequest() *TaobaoFeedflowItemTargetValidlistRequest{
    return &TaobaoFeedflowItemTargetValidlistRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemTargetValidlistRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.target.validlist"
}

func (r TaobaoFeedflowItemTargetValidlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemTargetValidlistRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoFeedflowItemTargetValidlistRequest) GetCampaignId() int64 {
    return r.campaignId
}

