package interact

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorCalendarAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorCalendarAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.calendar"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorCalendarAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorCalendarAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractSensorCalendarAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorCalendarRequest()
	},
}

// GetAlibabaInteractSensorCalendarRequest 从 sync.Pool 获取 AlibabaInteractSensorCalendarAPIRequest
func GetAlibabaInteractSensorCalendarAPIRequest() *AlibabaInteractSensorCalendarAPIRequest {
	return poolAlibabaInteractSensorCalendarAPIRequest.Get().(*AlibabaInteractSensorCalendarAPIRequest)
}

// ReleaseAlibabaInteractSensorCalendarAPIRequest 将 AlibabaInteractSensorCalendarAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorCalendarAPIRequest(v *AlibabaInteractSensorCalendarAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorCalendarAPIRequest.Put(v)
}
