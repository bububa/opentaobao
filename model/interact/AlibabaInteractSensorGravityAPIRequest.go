package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorGravityAPIRequest 重力感应 API请求
// alibaba.interact.sensor.gravity
//
// native获取重力感应
type AlibabaInteractSensorGravityAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorGravityRequest 初始化AlibabaInteractSensorGravityAPIRequest对象
func NewAlibabaInteractSensorGravityRequest() *AlibabaInteractSensorGravityAPIRequest {
	return &AlibabaInteractSensorGravityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorGravityAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.gravity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorGravityAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
