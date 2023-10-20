package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemcampaignmodifyAPIRequest 信息流修改计划 API请求
// taobao.feedflow.item.campaign.modify
//
// 信息流修改计划
type TaobaofeedflowitemcampaignmodifyAPIRequest struct {
	model.Params
	// 修改参数
	_campaign *CampaignDto
}

// NewTaobaofeedflowitemcampaignmodifyRequest 初始化TaobaofeedflowitemcampaignmodifyAPIRequest对象
func NewTaobaofeedflowitemcampaignmodifyRequest() *TaobaofeedflowitemcampaignmodifyAPIRequest {
	return &TaobaofeedflowitemcampaignmodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemcampaignmodifyAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.campaign.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemcampaignmodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemcampaignmodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCampaign is Campaign Setter
// 修改参数
func (r *TaobaofeedflowitemcampaignmodifyAPIRequest) SetCampaign(_campaign *CampaignDto) error {
	r._campaign = _campaign
	r.Set("campaign", _campaign)
	return nil
}

// GetCampaign Campaign Getter
func (r TaobaofeedflowitemcampaignmodifyAPIRequest) GetCampaign() *CampaignDto {
	return r._campaign
}
