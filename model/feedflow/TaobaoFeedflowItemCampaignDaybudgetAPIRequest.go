package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemcampaigndaybudgetAPIRequest 获取当日投放日预算总额 API请求
// taobao.feedflow.item.campaign.daybudget
//
// 获取当日投放日预算总额
type TaobaofeedflowitemcampaigndaybudgetAPIRequest struct {
	model.Params
}

// NewTaobaofeedflowitemcampaigndaybudgetRequest 初始化TaobaofeedflowitemcampaigndaybudgetAPIRequest对象
func NewTaobaofeedflowitemcampaigndaybudgetRequest() *TaobaofeedflowitemcampaigndaybudgetAPIRequest {
	return &TaobaofeedflowitemcampaigndaybudgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemcampaigndaybudgetAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.campaign.daybudget"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemcampaigndaybudgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemcampaigndaybudgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
