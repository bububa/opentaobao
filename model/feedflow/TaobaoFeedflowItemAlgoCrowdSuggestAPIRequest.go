package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest 单品人群建议出价 API请求
// taobao.feedflow.item.algo.crowd.suggest
//
// 给超级推荐的广告主查看建议出价
type TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest struct {
	model.Params
	// 人群列表
	_crowds []CrowdDto
	// 预估的宝贝id
	_itemId int64
	// 预估的计划id
	_campaignId int64
}

// NewTaobaoFeedflowItemAlgoCrowdSuggestRequest 初始化TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest对象
func NewTaobaoFeedflowItemAlgoCrowdSuggestRequest() *TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest {
	return &TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.algo.crowd.suggest"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCrowds is Crowds Setter
// 人群列表
func (r *TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest) SetCrowds(_crowds []CrowdDto) error {
	r._crowds = _crowds
	r.Set("crowds", _crowds)
	return nil
}

// GetCrowds Crowds Getter
func (r TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest) GetCrowds() []CrowdDto {
	return r._crowds
}

// SetItemId is ItemId Setter
// 预估的宝贝id
func (r *TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetCampaignId is CampaignId Setter
// 预估的计划id
func (r *TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
