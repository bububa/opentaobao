package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripCarOrderAcceptAPIRequest
确认接单 API请求
taobao.alitrip.car.order.accept

用来接收服务商确认接单信息 */
type TaobaoAlitripCarOrderAcceptAPIRequest struct {
	model.Params
	// 确认订单请求
	_paramOrderAccept *OrderAccept
}

// NewTaobaoAlitripCarOrderAcceptRequest 初始化TaobaoAlitripCarOrderAcceptAPIRequest对象
func NewTaobaoAlitripCarOrderAcceptRequest() *TaobaoAlitripCarOrderAcceptAPIRequest {
	return &TaobaoAlitripCarOrderAcceptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripCarOrderAcceptAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.car.order.accept"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripCarOrderAcceptAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamOrderAccept Setter
// 确认订单请求
func (r *TaobaoAlitripCarOrderAcceptAPIRequest) SetParamOrderAccept(_paramOrderAccept *OrderAccept) error {
	r._paramOrderAccept = _paramOrderAccept
	r.Set("param_order_accept", _paramOrderAccept)
	return nil
}

// Get ParamOrderAccept Getter
func (r TaobaoAlitripCarOrderAcceptAPIRequest) GetParamOrderAccept() *OrderAccept {
	return r._paramOrderAccept
}
