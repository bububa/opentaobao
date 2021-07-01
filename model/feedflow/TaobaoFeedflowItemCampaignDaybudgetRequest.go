package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取当日投放日预算总额 API请求
taobao.feedflow.item.campaign.daybudget

获取当日投放日预算总额
*/
type TaobaoFeedflowItemCampaignDaybudgetAPIRequest struct {
    model.Params
}

// 初始化TaobaoFeedflowItemCampaignDaybudgetAPIRequest对象
func NewTaobaoFeedflowItemCampaignDaybudgetRequest() *TaobaoFeedflowItemCampaignDaybudgetAPIRequest{
    return &TaobaoFeedflowItemCampaignDaybudgetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCampaignDaybudgetAPIRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.campaign.daybudget"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCampaignDaybudgetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
