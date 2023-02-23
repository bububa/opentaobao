package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorOpenwindowAPIRequest 客户端打开新页面 API请求
// alibaba.interact.sensor.openwindow
//
// 客户端打开新页面
type AlibabaInteractSensorOpenwindowAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorOpenwindowRequest 初始化AlibabaInteractSensorOpenwindowAPIRequest对象
func NewAlibabaInteractSensorOpenwindowRequest() *AlibabaInteractSensorOpenwindowAPIRequest {
	return &AlibabaInteractSensorOpenwindowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorOpenwindowAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.openwindow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorOpenwindowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorOpenwindowAPIRequest) GetRawParams() model.Params {
	return r.Params
}
