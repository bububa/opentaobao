package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorGcanvasAPIRequest gcanvas API请求
// alibaba.interact.sensor.gcanvas
//
// gcanvas 功能
type AlibabaInteractSensorGcanvasAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorGcanvasRequest 初始化AlibabaInteractSensorGcanvasAPIRequest对象
func NewAlibabaInteractSensorGcanvasRequest() *AlibabaInteractSensorGcanvasAPIRequest {
	return &AlibabaInteractSensorGcanvasAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorGcanvasAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.gcanvas"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorGcanvasAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorGcanvasAPIRequest) GetRawParams() model.Params {
	return r.Params
}
