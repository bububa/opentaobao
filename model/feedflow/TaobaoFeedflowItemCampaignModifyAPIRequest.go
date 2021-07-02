package feedflow

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCampaignModifyAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.campaign.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCampaignModifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Campaign Setter
// 修改参数
func (r *TaobaoFeedflowItemCampaignModifyAPIRequest) SetCampaign(_campaign *CampaignDto) error {
	r._campaign = _campaign
	r.Set("campaign", _campaign)
	return nil
}

// Get Campaign Getter
func (r TaobaoFeedflowItemCampaignModifyAPIRequest) GetCampaign() *CampaignDto {
	return r._campaign
}
