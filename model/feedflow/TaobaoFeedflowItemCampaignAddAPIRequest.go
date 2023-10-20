package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemcampaignaddAPIRequest 信息流增加推广计划 API请求
// taobao.feedflow.item.campaign.add
//
// 信息流增加推广计划
type TaobaofeedflowitemcampaignaddAPIRequest struct {
	model.Params
	// 计划信息
	_campaign *CampaignDto
}

// NewTaobaofeedflowitemcampaignaddRequest 初始化TaobaofeedflowitemcampaignaddAPIRequest对象
func NewTaobaofeedflowitemcampaignaddRequest() *TaobaofeedflowitemcampaignaddAPIRequest {
	return &TaobaofeedflowitemcampaignaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemcampaignaddAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.campaign.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemcampaignaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemcampaignaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCampaign is Campaign Setter
// 计划信息
func (r *TaobaofeedflowitemcampaignaddAPIRequest) SetCampaign(_campaign *CampaignDto) error {
	r._campaign = _campaign
	r.Set("campaign", _campaign)
	return nil
}

// GetCampaign Campaign Getter
func (r TaobaofeedflowitemcampaignaddAPIRequest) GetCampaign() *CampaignDto {
	return r._campaign
}
