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
type TaobaoFeedflowItemCampaignAddAPIRequest struct {
    model.Params
    // 计划信息
    _campaign   *CampaignDTO
}

// 初始化TaobaoFeedflowItemCampaignAddAPIRequest对象
func NewTaobaoFeedflowItemCampaignAddRequest() *TaobaoFeedflowItemCampaignAddAPIRequest{
    return &TaobaoFeedflowItemCampaignAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCampaignAddAPIRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.campaign.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCampaignAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Campaign Setter
// 计划信息
func (r *TaobaoFeedflowItemCampaignAddAPIRequest) SetCampaign(_campaign *CampaignDTO) error {
    r._campaign = _campaign
    r.Set("campaign", _campaign)
    return nil
}

// Campaign Getter
func (r TaobaoFeedflowItemCampaignAddAPIRequest) GetCampaign() *CampaignDTO {
    return r._campaign
}
