package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractSensorMaAPIRequest
码相关API API请求
alibaba.interact.sensor.ma

码相关API */
type AlibabaInteractSensorMaAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorMaRequest 初始化AlibabaInteractSensorMaAPIRequest对象
func NewAlibabaInteractSensorMaRequest() *AlibabaInteractSensorMaAPIRequest {
	return &AlibabaInteractSensorMaAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorMaAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.ma"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorMaAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
