package feedflow

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest) Reset() {
	r._crowds = r._crowds[:0]
	r._itemId = 0
	r._campaignId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.algo.crowd.suggest"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoFeedflowItemAlgoCrowdSuggestAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemAlgoCrowdSuggestRequest()
	},
}

// GetTaobaoFeedflowItemAlgoCrowdSuggestRequest 从 sync.Pool 获取 TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest
func GetTaobaoFeedflowItemAlgoCrowdSuggestAPIRequest() *TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest {
	return poolTaobaoFeedflowItemAlgoCrowdSuggestAPIRequest.Get().(*TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest)
}

// ReleaseTaobaoFeedflowItemAlgoCrowdSuggestAPIRequest 将 TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemAlgoCrowdSuggestAPIRequest(v *TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemAlgoCrowdSuggestAPIRequest.Put(v)
}
