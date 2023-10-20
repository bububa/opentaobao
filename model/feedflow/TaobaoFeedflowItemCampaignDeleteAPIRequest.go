package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemcampaigndeleteAPIRequest 删除计划 API请求
// taobao.feedflow.item.campaign.delete
//
// 删除计划
type TaobaofeedflowitemcampaigndeleteAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
}

// NewTaobaofeedflowitemcampaigndeleteRequest 初始化TaobaofeedflowitemcampaigndeleteAPIRequest对象
func NewTaobaofeedflowitemcampaigndeleteRequest() *TaobaofeedflowitemcampaigndeleteAPIRequest {
	return &TaobaofeedflowitemcampaigndeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemcampaigndeleteAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.campaign.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemcampaigndeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemcampaigndeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *TaobaofeedflowitemcampaigndeleteAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaofeedflowitemcampaigndeleteAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
