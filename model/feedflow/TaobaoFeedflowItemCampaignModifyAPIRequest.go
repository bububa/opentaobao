package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCampaignModifyAPIRequest 信息流修改计划 API请求
// taobao.feedflow.item.campaign.modify
//
// 信息流修改计划
type TaobaoFeedflowItemCampaignModifyAPIRequest struct {
	model.Params
	// 修改参数
	_campaign *CampaignDto
}

// NewTaobaoFeedflowItemCampaignModifyRequest 初始化TaobaoFeedflowItemCampaignModifyAPIRequest对象
func NewTaobaoFeedflowItemCampaignModifyRequest() *TaobaoFeedflowItemCampaignModifyAPIRequest {
	return &TaobaoFeedflowItemCampaignModifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemCampaignModifyAPIRequest) Reset() {
	r._campaign = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCampaignModifyAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.campaign.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCampaignModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemCampaignModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCampaign is Campaign Setter
// 修改参数
func (r *TaobaoFeedflowItemCampaignModifyAPIRequest) SetCampaign(_campaign *CampaignDto) error {
	r._campaign = _campaign
	r.Set("campaign", _campaign)
	return nil
}

// GetCampaign Campaign Getter
func (r TaobaoFeedflowItemCampaignModifyAPIRequest) GetCampaign() *CampaignDto {
	return r._campaign
}

var poolTaobaoFeedflowItemCampaignModifyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemCampaignModifyRequest()
	},
}

// GetTaobaoFeedflowItemCampaignModifyRequest 从 sync.Pool 获取 TaobaoFeedflowItemCampaignModifyAPIRequest
func GetTaobaoFeedflowItemCampaignModifyAPIRequest() *TaobaoFeedflowItemCampaignModifyAPIRequest {
	return poolTaobaoFeedflowItemCampaignModifyAPIRequest.Get().(*TaobaoFeedflowItemCampaignModifyAPIRequest)
}

// ReleaseTaobaoFeedflowItemCampaignModifyAPIRequest 将 TaobaoFeedflowItemCampaignModifyAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemCampaignModifyAPIRequest(v *TaobaoFeedflowItemCampaignModifyAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemCampaignModifyAPIRequest.Put(v)
}
