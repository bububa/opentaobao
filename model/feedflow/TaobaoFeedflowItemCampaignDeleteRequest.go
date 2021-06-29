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
type TaobaoFeedflowItemCampaignDeleteRequest struct {
    model.Params
    // 计划id
    campaignId   int64
}

// 初始化TaobaoFeedflowItemCampaignDeleteRequest对象
func NewTaobaoFeedflowItemCampaignDeleteRequest() *TaobaoFeedflowItemCampaignDeleteRequest{
    return &TaobaoFeedflowItemCampaignDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCampaignDeleteRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.campaign.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCampaignDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *TaobaoFeedflowItemCampaignDeleteRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoFeedflowItemCampaignDeleteRequest) GetCampaignId() int64 {
    return r.campaignId
}
