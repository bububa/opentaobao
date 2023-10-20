package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemcampaigngetAPIRequest 通过计划id查询计划 API请求
// taobao.feedflow.item.campaign.get
//
// 通过计划id查询计划
type TaobaofeedflowitemcampaigngetAPIRequest struct {
	model.Params
	// 计划id
	_campaginId int64
}

// NewTaobaofeedflowitemcampaigngetRequest 初始化TaobaofeedflowitemcampaigngetAPIRequest对象
func NewTaobaofeedflowitemcampaigngetRequest() *TaobaofeedflowitemcampaigngetAPIRequest {
	return &TaobaofeedflowitemcampaigngetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemcampaigngetAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.campaign.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemcampaigngetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemcampaigngetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCampaginId is CampaginId Setter
// 计划id
func (r *TaobaofeedflowitemcampaigngetAPIRequest) SetCampaginId(_campaginId int64) error {
	r._campaginId = _campaginId
	r.Set("campagin_id", _campaginId)
	return nil
}

// GetCampaginId CampaginId Getter
func (r TaobaofeedflowitemcampaigngetAPIRequest) GetCampaginId() int64 {
	return r._campaginId
}
