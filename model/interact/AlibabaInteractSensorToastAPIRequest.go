package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorToastAPIRequest toast API请求
// alibaba.interact.sensor.toast
//
// toast提示
type AlibabaInteractSensorToastAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorToastRequest 初始化AlibabaInteractSensorToastAPIRequest对象
func NewAlibabaInteractSensorToastRequest() *AlibabaInteractSensorToastAPIRequest {
	return &AlibabaInteractSensorToastAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorToastAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.toast"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorToastAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorToastAPIRequest) GetRawParams() model.Params {
	return r.Params
}
