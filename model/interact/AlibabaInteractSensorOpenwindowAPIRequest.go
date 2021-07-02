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
func (r AlibabaInteractSensorOpenwindowAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
