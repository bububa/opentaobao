package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取有权限的定向列表 API请求
taobao.feedflow.item.target.validlist

获取有权限的定向列表
*/
type TaobaoFeedflowItemTargetValidlistRequest struct {
    model.Params
    // 计划id
    campaignId   int64
}

// 初始化TaobaoFeedflowItemTargetValidlistRequest对象
func NewTaobaoFeedflowItemTargetValidlistRequest() *TaobaoFeedflowItemTargetValidlistRequest{
    return &TaobaoFeedflowItemTargetValidlistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemTargetValidlistRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.target.validlist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemTargetValidlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *TaobaoFeedflowItemTargetValidlistRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoFeedflowItemTargetValidlistRequest) GetCampaignId() int64 {
    return r.campaignId
}
