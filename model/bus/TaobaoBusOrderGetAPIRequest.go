package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusOrderGetAPIRequest 汽车票订单查询 API请求
// taobao.bus.order.get
//
// 商家汽车票订单查询
type TaobaoBusOrderGetAPIRequest struct {
	model.Params
	// 订单查询对象
	_paramB2BOrderQueryRQ *B2BOrderQueryRq
}

// NewTaobaoBusOrderGetRequest 初始化TaobaoBusOrderGetAPIRequest对象
func NewTaobaoBusOrderGetRequest() *TaobaoBusOrderGetAPIRequest {
	return &TaobaoBusOrderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusOrderGetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamB2BOrderQueryRQ is ParamB2BOrderQueryRQ Setter
// 订单查询对象
func (r *TaobaoBusOrderGetAPIRequest) SetParamB2BOrderQueryRQ(_paramB2BOrderQueryRQ *B2BOrderQueryRq) error {
	r._paramB2BOrderQueryRQ = _paramB2BOrderQueryRQ
	r.Set("param_b2_b_order_query_r_q", _paramB2BOrderQueryRQ)
	return nil
}

// GetParamB2BOrderQueryRQ ParamB2BOrderQueryRQ Getter
func (r TaobaoBusOrderGetAPIRequest) GetParamB2BOrderQueryRQ() *B2BOrderQueryRq {
	return r._paramB2BOrderQueryRQ
}
