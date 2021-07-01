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
type TaobaoFeedflowItemTargetValidlistAPIRequest struct {
    model.Params
    // 计划id
    _campaignId   int64
}

// 初始化TaobaoFeedflowItemTargetValidlistAPIRequest对象
func NewTaobaoFeedflowItemTargetValidlistRequest() *TaobaoFeedflowItemTargetValidlistAPIRequest{
    return &TaobaoFeedflowItemTargetValidlistAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemTargetValidlistAPIRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.target.validlist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemTargetValidlistAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *TaobaoFeedflowItemTargetValidlistAPIRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoFeedflowItemTargetValidlistAPIRequest) GetCampaignId() int64 {
    return r._campaignId
}
