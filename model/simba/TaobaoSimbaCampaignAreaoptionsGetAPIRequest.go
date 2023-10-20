package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbacampaignareaoptionsgetAPIRequest 取得推广计划的可设置投放地域列表 API请求
// taobao.simba.campaign.areaoptions.get
//
// 取得推广计划的可设置投放地域列表
type TaobaosimbacampaignareaoptionsgetAPIRequest struct {
	model.Params
}

// NewTaobaosimbacampaignareaoptionsgetRequest 初始化TaobaosimbacampaignareaoptionsgetAPIRequest对象
func NewTaobaosimbacampaignareaoptionsgetRequest() *TaobaosimbacampaignareaoptionsgetAPIRequest {
	return &TaobaosimbacampaignareaoptionsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbacampaignareaoptionsgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.areaoptions.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbacampaignareaoptionsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbacampaignareaoptionsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
