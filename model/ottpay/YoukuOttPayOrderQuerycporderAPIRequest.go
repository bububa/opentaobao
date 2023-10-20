package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukuottpayorderquerycporderAPIRequest 查询支付订单对应cp订单号 API请求
// youku.ott.pay.order.querycporder
//
// 根据支付订单查询对应cp订单号
type YoukuottpayorderquerycporderAPIRequest struct {
	model.Params
	// 支付对应订单
	_gatewayOrder string
}

// NewYoukuottpayorderquerycporderRequest 初始化YoukuottpayorderquerycporderAPIRequest对象
func NewYoukuottpayorderquerycporderRequest() *YoukuottpayorderquerycporderAPIRequest {
	return &YoukuottpayorderquerycporderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuottpayorderquerycporderAPIRequest) GetApiMethodName() string {
	return "youku.ott.pay.order.querycporder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuottpayorderquerycporderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuottpayorderquerycporderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGatewayOrder is GatewayOrder Setter
// 支付对应订单
func (r *YoukuottpayorderquerycporderAPIRequest) SetGatewayOrder(_gatewayOrder string) error {
	r._gatewayOrder = _gatewayOrder
	r.Set("gateway_order", _gatewayOrder)
	return nil
}

// GetGatewayOrder GatewayOrder Getter
func (r YoukuottpayorderquerycporderAPIRequest) GetGatewayOrder() string {
	return r._gatewayOrder
}
