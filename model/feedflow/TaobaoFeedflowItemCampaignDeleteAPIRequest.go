package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCampaignDeleteAPIRequest 删除计划 API请求
// taobao.feedflow.item.campaign.delete
//
// 删除计划
type TaobaoFeedflowItemCampaignDeleteAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
}

// NewTaobaoFeedflowItemCampaignDeleteRequest 初始化TaobaoFeedflowItemCampaignDeleteAPIRequest对象
func NewTaobaoFeedflowItemCampaignDeleteRequest() *TaobaoFeedflowItemCampaignDeleteAPIRequest {
	return &TaobaoFeedflowItemCampaignDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemCampaignDeleteAPIRequest) Reset() {
	r._campaignId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCampaignDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.campaign.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCampaignDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemCampaignDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *TaobaoFeedflowItemCampaignDeleteAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoFeedflowItemCampaignDeleteAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

var poolTaobaoFeedflowItemCampaignDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemCampaignDeleteRequest()
	},
}

// GetTaobaoFeedflowItemCampaignDeleteRequest 从 sync.Pool 获取 TaobaoFeedflowItemCampaignDeleteAPIRequest
func GetTaobaoFeedflowItemCampaignDeleteAPIRequest() *TaobaoFeedflowItemCampaignDeleteAPIRequest {
	return poolTaobaoFeedflowItemCampaignDeleteAPIRequest.Get().(*TaobaoFeedflowItemCampaignDeleteAPIRequest)
}

// ReleaseTaobaoFeedflowItemCampaignDeleteAPIRequest 将 TaobaoFeedflowItemCampaignDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemCampaignDeleteAPIRequest(v *TaobaoFeedflowItemCampaignDeleteAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemCampaignDeleteAPIRequest.Put(v)
}
