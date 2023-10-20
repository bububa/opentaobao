package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensorcalendarAPIRequest 天猫互动游戏开放平台需要授权的传感器类接口(日历提醒) API请求
// alibaba.interact.sensor.calendar
//
// 天猫互动游戏开放平台需要授权的传感器类接口(日历提醒)
type AlibabainteractsensorcalendarAPIRequest struct {
	model.Params
}

// NewAlibabainteractsensorcalendarRequest 初始化AlibabainteractsensorcalendarAPIRequest对象
func NewAlibabainteractsensorcalendarRequest() *AlibabainteractsensorcalendarAPIRequest {
	return &AlibabainteractsensorcalendarAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensorcalendarAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.calendar"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensorcalendarAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensorcalendarAPIRequest) GetRawParams() model.Params {
	return r.Params
}
