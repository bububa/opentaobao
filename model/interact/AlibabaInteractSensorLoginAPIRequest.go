package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractSensorLoginAPIRequest
获取登陆页面 API请求
alibaba.interact.sensor.login

获取登陆页面 */
type AlibabaInteractSensorLoginAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorLoginRequest 初始化AlibabaInteractSensorLoginAPIRequest对象
func NewAlibabaInteractSensorLoginRequest() *AlibabaInteractSensorLoginAPIRequest {
	return &AlibabaInteractSensorLoginAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorLoginAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.login"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorLoginAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
