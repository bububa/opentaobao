package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensornetworkstatusAPIRequest 网络状态 API请求
// alibaba.interact.sensor.networkstatus
//
// 客户端网络状态
type AlibabainteractsensornetworkstatusAPIRequest struct {
	model.Params
}

// NewAlibabainteractsensornetworkstatusRequest 初始化AlibabainteractsensornetworkstatusAPIRequest对象
func NewAlibabainteractsensornetworkstatusRequest() *AlibabainteractsensornetworkstatusAPIRequest {
	return &AlibabainteractsensornetworkstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensornetworkstatusAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.networkstatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensornetworkstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensornetworkstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}
