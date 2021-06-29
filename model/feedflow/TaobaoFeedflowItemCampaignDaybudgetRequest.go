package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取当日投放日预算总额 APIRequest
taobao.feedflow.item.campaign.daybudget

获取当日投放日预算总额
*/
type TaobaoFeedflowItemCampaignDaybudgetRequest struct {
    model.Params

}

func NewTaobaoFeedflowItemCampaignDaybudgetRequest() *TaobaoFeedflowItemCampaignDaybudgetRequest{
    return &TaobaoFeedflowItemCampaignDaybudgetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemCampaignDaybudgetRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.campaign.daybudget"
}

func (r TaobaoFeedflowItemCampaignDaybudgetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


