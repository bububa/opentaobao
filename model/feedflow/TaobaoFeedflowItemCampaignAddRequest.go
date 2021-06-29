package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流增加推广计划 API请求
taobao.feedflow.item.campaign.add

信息流增加推广计划
*/
type TaobaoFeedflowItemCampaignAddRequest struct {
    model.Params
    // 计划信息
    _campaign   *CampaignDto
}

// 初始化TaobaoFeedflowItemCampaignAddRequest对象
func NewTaobaoFeedflowItemCampaignAddRequest() *TaobaoFeedflowItemCampaignAddRequest{
    return &TaobaoFeedflowItemCampaignAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCampaignAddRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.campaign.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCampaignAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Campaign Setter
// 计划信息
func (r *TaobaoFeedflowItemCampaignAddRequest) SetCampaign(_campaign *CampaignDto) error {
    r._campaign = _campaign
    r.Set("campaign", _campaign)
    return nil
}

// Campaign Getter
func (r TaobaoFeedflowItemCampaignAddRequest) GetCampaign() *CampaignDto {
    return r._campaign
}
