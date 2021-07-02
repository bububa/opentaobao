package feedflow

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCampaignAddAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.campaign.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCampaignAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Campaign Setter
// 计划信息
func (r *TaobaoFeedflowItemCampaignAddAPIRequest) SetCampaign(_campaign *CampaignDto) error {
	r._campaign = _campaign
	r.Set("campaign", _campaign)
	return nil
}

// Get Campaign Getter
func (r TaobaoFeedflowItemCampaignAddAPIRequest) GetCampaign() *CampaignDto {
	return r._campaign
}
