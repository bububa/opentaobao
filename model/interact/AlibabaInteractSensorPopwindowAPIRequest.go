package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractSensorPopwindowAPIRequest
popwindow API请求
alibaba.interact.sensor.popwindow

popwindow */
type AlibabaInteractSensorPopwindowAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorPopwindowRequest 初始化AlibabaInteractSensorPopwindowAPIRequest对象
func NewAlibabaInteractSensorPopwindowRequest() *AlibabaInteractSensorPopwindowAPIRequest {
	return &AlibabaInteractSensorPopwindowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorPopwindowAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.popwindow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorPopwindowAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
