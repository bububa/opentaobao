package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过计划id查询计划 APIRequest
taobao.feedflow.item.campaign.get

通过计划id查询计划
*/
type TaobaoFeedflowItemCampaignGetRequest struct {
    model.Params

    // 计划id
    campaginId   int64 

}

func NewTaobaoFeedflowItemCampaignGetRequest() *TaobaoFeedflowItemCampaignGetRequest{
    return &TaobaoFeedflowItemCampaignGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemCampaignGetRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.campaign.get"
}

func (r TaobaoFeedflowItemCampaignGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemCampaignGetRequest) SetCampaginId(campaginId int64) error {
    r.campaginId = campaginId
    r.Set("campagin_id", campaginId)
    return nil
}

func (r TaobaoFeedflowItemCampaignGetRequest) GetCampaginId() int64 {
    return r.campaginId
}

