package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流修改计划 API请求
taobao.feedflow.item.campaign.modify

信息流修改计划
*/
type TaobaoFeedflowItemCampaignModifyAPIRequest struct {
    model.Params
    // 修改参数
    _campaign   *CampaignDTO
}

// 初始化TaobaoFeedflowItemCampaignModifyAPIRequest对象
func NewTaobaoFeedflowItemCampaignModifyRequest() *TaobaoFeedflowItemCampaignModifyAPIRequest{
    return &TaobaoFeedflowItemCampaignModifyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCampaignModifyAPIRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.campaign.modify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCampaignModifyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Campaign Setter
// 修改参数
func (r *TaobaoFeedflowItemCampaignModifyAPIRequest) SetCampaign(_campaign *CampaignDTO) error {
    r._campaign = _campaign
    r.Set("campaign", _campaign)
    return nil
}

// Campaign Getter
func (r TaobaoFeedflowItemCampaignModifyAPIRequest) GetCampaign() *CampaignDTO {
    return r._campaign
}
