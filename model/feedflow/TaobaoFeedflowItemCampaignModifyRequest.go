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
type TaobaoFeedflowItemCampaignModifyRequest struct {
    model.Params
    // 修改参数
    campaign   *CampaignDto
}

// 初始化TaobaoFeedflowItemCampaignModifyRequest对象
func NewTaobaoFeedflowItemCampaignModifyRequest() *TaobaoFeedflowItemCampaignModifyRequest{
    return &TaobaoFeedflowItemCampaignModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCampaignModifyRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.campaign.modify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCampaignModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Campaign Setter
// 修改参数
func (r *TaobaoFeedflowItemCampaignModifyRequest) SetCampaign(campaign *CampaignDto) error {
    r.campaign = campaign
    r.Set("campaign", campaign)
    return nil
}

// Campaign Getter
func (r TaobaoFeedflowItemCampaignModifyRequest) GetCampaign() *CampaignDto {
    return r.campaign
}
