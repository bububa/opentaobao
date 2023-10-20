package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorVibrateAPIRequest 客户端震动 API请求
// alibaba.interact.sensor.vibrate
//
// 客户端震动
type AlibabaInteractSensorVibrateAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorVibrateRequest 初始化AlibabaInteractSensorVibrateAPIRequest对象
func NewAlibabaInteractSensorVibrateRequest() *AlibabaInteractSensorVibrateAPIRequest {
	return &AlibabaInteractSensorVibrateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorVibrateAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.vibrate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorVibrateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorVibrateAPIRequest) GetRawParams() model.Params {
	return r.Params
}
