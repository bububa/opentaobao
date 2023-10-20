package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCampaignAddAPIRequest 信息流增加推广计划 API请求
// taobao.feedflow.item.campaign.add
//
// 信息流增加推广计划
type TaobaoFeedflowItemCampaignAddAPIRequest struct {
	model.Params
	// 计划信息
	_campaign *CampaignDto
}

// NewTaobaoFeedflowItemCampaignAddRequest 初始化TaobaoFeedflowItemCampaignAddAPIRequest对象
func NewTaobaoFeedflowItemCampaignAddRequest() *TaobaoFeedflowItemCampaignAddAPIRequest {
	return &TaobaoFeedflowItemCampaignAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemCampaignAddAPIRequest) Reset() {
	r._campaign = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCampaignAddAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.campaign.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCampaignAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemCampaignAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCampaign is Campaign Setter
// 计划信息
func (r *TaobaoFeedflowItemCampaignAddAPIRequest) SetCampaign(_campaign *CampaignDto) error {
	r._campaign = _campaign
	r.Set("campaign", _campaign)
	return nil
}

// GetCampaign Campaign Getter
func (r TaobaoFeedflowItemCampaignAddAPIRequest) GetCampaign() *CampaignDto {
	return r._campaign
}

var poolTaobaoFeedflowItemCampaignAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemCampaignAddRequest()
	},
}

// GetTaobaoFeedflowItemCampaignAddRequest 从 sync.Pool 获取 TaobaoFeedflowItemCampaignAddAPIRequest
func GetTaobaoFeedflowItemCampaignAddAPIRequest() *TaobaoFeedflowItemCampaignAddAPIRequest {
	return poolTaobaoFeedflowItemCampaignAddAPIRequest.Get().(*TaobaoFeedflowItemCampaignAddAPIRequest)
}

// ReleaseTaobaoFeedflowItemCampaignAddAPIRequest 将 TaobaoFeedflowItemCampaignAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemCampaignAddAPIRequest(v *TaobaoFeedflowItemCampaignAddAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemCampaignAddAPIRequest.Put(v)
}
