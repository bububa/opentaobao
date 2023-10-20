package ottpay

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttPayOrderQuerycporderAPIRequest 查询支付订单对应cp订单号 API请求
// youku.ott.pay.order.querycporder
//
// 根据支付订单查询对应cp订单号
type YoukuOttPayOrderQuerycporderAPIRequest struct {
	model.Params
	// 支付对应订单
	_gatewayOrder string
}

// NewYoukuOttPayOrderQuerycporderRequest 初始化YoukuOttPayOrderQuerycporderAPIRequest对象
func NewYoukuOttPayOrderQuerycporderRequest() *YoukuOttPayOrderQuerycporderAPIRequest {
	return &YoukuOttPayOrderQuerycporderAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YoukuOttPayOrderQuerycporderAPIRequest) Reset() {
	r._gatewayOrder = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuOttPayOrderQuerycporderAPIRequest) GetApiMethodName() string {
	return "youku.ott.pay.order.querycporder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuOttPayOrderQuerycporderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuOttPayOrderQuerycporderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGatewayOrder is GatewayOrder Setter
// 支付对应订单
func (r *YoukuOttPayOrderQuerycporderAPIRequest) SetGatewayOrder(_gatewayOrder string) error {
	r._gatewayOrder = _gatewayOrder
	r.Set("gateway_order", _gatewayOrder)
	return nil
}

// GetGatewayOrder GatewayOrder Getter
func (r YoukuOttPayOrderQuerycporderAPIRequest) GetGatewayOrder() string {
	return r._gatewayOrder
}

var poolYoukuOttPayOrderQuerycporderAPIRequest = sync.Pool{
	New: func() any {
		return NewYoukuOttPayOrderQuerycporderRequest()
	},
}

// GetYoukuOttPayOrderQuerycporderRequest 从 sync.Pool 获取 YoukuOttPayOrderQuerycporderAPIRequest
func GetYoukuOttPayOrderQuerycporderAPIRequest() *YoukuOttPayOrderQuerycporderAPIRequest {
	return poolYoukuOttPayOrderQuerycporderAPIRequest.Get().(*YoukuOttPayOrderQuerycporderAPIRequest)
}

// ReleaseYoukuOttPayOrderQuerycporderAPIRequest 将 YoukuOttPayOrderQuerycporderAPIRequest 放入 sync.Pool
func ReleaseYoukuOttPayOrderQuerycporderAPIRequest(v *YoukuOttPayOrderQuerycporderAPIRequest) {
	v.Reset()
	poolYoukuOttPayOrderQuerycporderAPIRequest.Put(v)
}
