package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorCalendarAPIRequest 天猫互动游戏开放平台需要授权的传感器类接口(日历提醒) API请求
// alibaba.interact.sensor.calendar
//
// 天猫互动游戏开放平台需要授权的传感器类接口(日历提醒)
type AlibabaInteractSensorCalendarAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorCalendarRequest 初始化AlibabaInteractSensorCalendarAPIRequest对象
func NewAlibabaInteractSensorCalendarRequest() *AlibabaInteractSensorCalendarAPIRequest {
	return &AlibabaInteractSensorCalendarAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorCalendarAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.calendar"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorCalendarAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
