package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemTargetValidlistAPIRequest 获取有权限的定向列表 API请求
// taobao.feedflow.item.target.validlist
//
// 获取有权限的定向列表
type TaobaoFeedflowItemTargetValidlistAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
}

// NewTaobaoFeedflowItemTargetValidlistRequest 初始化TaobaoFeedflowItemTargetValidlistAPIRequest对象
func NewTaobaoFeedflowItemTargetValidlistRequest() *TaobaoFeedflowItemTargetValidlistAPIRequest {
	return &TaobaoFeedflowItemTargetValidlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemTargetValidlistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.target.validlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemTargetValidlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemTargetValidlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *TaobaoFeedflowItemTargetValidlistAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoFeedflowItemTargetValidlistAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
