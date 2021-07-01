package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除计划 API请求
taobao.feedflow.item.campaign.delete

删除计划
*/
type TaobaoFeedflowItemCampaignDeleteAPIRequest struct {
    model.Params
    // 计划id
    _campaignId   int64
}

// 初始化TaobaoFeedflowItemCampaignDeleteAPIRequest对象
func NewTaobaoFeedflowItemCampaignDeleteRequest() *TaobaoFeedflowItemCampaignDeleteAPIRequest{
    return &TaobaoFeedflowItemCampaignDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCampaignDeleteAPIRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.campaign.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCampaignDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *TaobaoFeedflowItemCampaignDeleteAPIRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoFeedflowItemCampaignDeleteAPIRequest) GetCampaignId() int64 {
    return r._campaignId
}
