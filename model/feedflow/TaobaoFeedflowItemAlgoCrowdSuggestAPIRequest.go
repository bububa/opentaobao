package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemalgocrowdsuggestAPIRequest 单品人群建议出价 API请求
// taobao.feedflow.item.algo.crowd.suggest
//
// 给超级推荐的广告主查看建议出价
type TaobaofeedflowitemalgocrowdsuggestAPIRequest struct {
	model.Params
	// 人群列表
	_crowds []CrowdDto
	// 预估的宝贝id
	_itemId int64
	// 预估的计划id
	_campaignId int64
}

// NewTaobaofeedflowitemalgocrowdsuggestRequest 初始化TaobaofeedflowitemalgocrowdsuggestAPIRequest对象
func NewTaobaofeedflowitemalgocrowdsuggestRequest() *TaobaofeedflowitemalgocrowdsuggestAPIRequest {
	return &TaobaofeedflowitemalgocrowdsuggestAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemalgocrowdsuggestAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.algo.crowd.suggest"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemalgocrowdsuggestAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemalgocrowdsuggestAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCrowds is Crowds Setter
// 人群列表
func (r *TaobaofeedflowitemalgocrowdsuggestAPIRequest) SetCrowds(_crowds []CrowdDto) error {
	r._crowds = _crowds
	r.Set("crowds", _crowds)
	return nil
}

// GetCrowds Crowds Getter
func (r TaobaofeedflowitemalgocrowdsuggestAPIRequest) GetCrowds() []CrowdDto {
	return r._crowds
}

// SetItemId is ItemId Setter
// 预估的宝贝id
func (r *TaobaofeedflowitemalgocrowdsuggestAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaofeedflowitemalgocrowdsuggestAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetCampaignId is CampaignId Setter
// 预估的计划id
func (r *TaobaofeedflowitemalgocrowdsuggestAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaofeedflowitemalgocrowdsuggestAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
