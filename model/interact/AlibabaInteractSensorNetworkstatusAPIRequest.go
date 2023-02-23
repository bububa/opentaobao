package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorNetworkstatusAPIRequest 网络状态 API请求
// alibaba.interact.sensor.networkstatus
//
// 客户端网络状态
type AlibabaInteractSensorNetworkstatusAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorNetworkstatusRequest 初始化AlibabaInteractSensorNetworkstatusAPIRequest对象
func NewAlibabaInteractSensorNetworkstatusRequest() *AlibabaInteractSensorNetworkstatusAPIRequest {
	return &AlibabaInteractSensorNetworkstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorNetworkstatusAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.networkstatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorNetworkstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorNetworkstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}
