package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemtargetvalidlistAPIRequest 获取有权限的定向列表 API请求
// taobao.feedflow.item.target.validlist
//
// 获取有权限的定向列表
type TaobaofeedflowitemtargetvalidlistAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
}

// NewTaobaofeedflowitemtargetvalidlistRequest 初始化TaobaofeedflowitemtargetvalidlistAPIRequest对象
func NewTaobaofeedflowitemtargetvalidlistRequest() *TaobaofeedflowitemtargetvalidlistAPIRequest {
	return &TaobaofeedflowitemtargetvalidlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemtargetvalidlistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.target.validlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemtargetvalidlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemtargetvalidlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *TaobaofeedflowitemtargetvalidlistAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaofeedflowitemtargetvalidlistAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
