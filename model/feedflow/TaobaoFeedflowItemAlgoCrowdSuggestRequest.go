package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单品人群建议出价 API请求
taobao.feedflow.item.algo.crowd.suggest

给超级推荐的广告主查看建议出价
*/
type TaobaoFeedflowItemAlgoCrowdSuggestRequest struct {
    model.Params
    // 人群列表
    _crowds   []CrowdDto
    // 预估的宝贝id
    _itemId   int64
    // 预估的计划id
    _campaignId   int64
}

// 初始化TaobaoFeedflowItemAlgoCrowdSuggestRequest对象
func NewTaobaoFeedflowItemAlgoCrowdSuggestRequest() *TaobaoFeedflowItemAlgoCrowdSuggestRequest{
    return &TaobaoFeedflowItemAlgoCrowdSuggestRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAlgoCrowdSuggestRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.algo.crowd.suggest"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAlgoCrowdSuggestRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Crowds Setter
// 人群列表
func (r *TaobaoFeedflowItemAlgoCrowdSuggestRequest) SetCrowds(_crowds []CrowdDto) error {
    r._crowds = _crowds
    r.Set("crowds", _crowds)
    return nil
}

// Crowds Getter
func (r TaobaoFeedflowItemAlgoCrowdSuggestRequest) GetCrowds() []CrowdDto {
    return r._crowds
}
// ItemId Setter
// 预估的宝贝id
func (r *TaobaoFeedflowItemAlgoCrowdSuggestRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoFeedflowItemAlgoCrowdSuggestRequest) GetItemId() int64 {
    return r._itemId
}
// CampaignId Setter
// 预估的计划id
func (r *TaobaoFeedflowItemAlgoCrowdSuggestRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoFeedflowItemAlgoCrowdSuggestRequest) GetCampaignId() int64 {
    return r._campaignId
}
