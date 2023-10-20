package bus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusOrderSetAPIRequest 汽车票下单接口 API请求
// taobao.bus.order.set
//
// 提供给汽车票商家进行下单
type TaobaoBusOrderSetAPIRequest struct {
	model.Params
	// 下单参数
	_paramB2BCreateOrderRQ *B2BCreateOrderRq
}

// NewTaobaoBusOrderSetRequest 初始化TaobaoBusOrderSetAPIRequest对象
func NewTaobaoBusOrderSetRequest() *TaobaoBusOrderSetAPIRequest {
	return &TaobaoBusOrderSetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBusOrderSetAPIRequest) Reset() {
	r._paramB2BCreateOrderRQ = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusOrderSetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.order.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusOrderSetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusOrderSetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamB2BCreateOrderRQ is ParamB2BCreateOrderRQ Setter
// 下单参数
func (r *TaobaoBusOrderSetAPIRequest) SetParamB2BCreateOrderRQ(_paramB2BCreateOrderRQ *B2BCreateOrderRq) error {
	r._paramB2BCreateOrderRQ = _paramB2BCreateOrderRQ
	r.Set("param_b2_b_create_order_r_q", _paramB2BCreateOrderRQ)
	return nil
}

// GetParamB2BCreateOrderRQ ParamB2BCreateOrderRQ Getter
func (r TaobaoBusOrderSetAPIRequest) GetParamB2BCreateOrderRQ() *B2BCreateOrderRq {
	return r._paramB2BCreateOrderRQ
}

var poolTaobaoBusOrderSetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBusOrderSetRequest()
	},
}

// GetTaobaoBusOrderSetRequest 从 sync.Pool 获取 TaobaoBusOrderSetAPIRequest
func GetTaobaoBusOrderSetAPIRequest() *TaobaoBusOrderSetAPIRequest {
	return poolTaobaoBusOrderSetAPIRequest.Get().(*TaobaoBusOrderSetAPIRequest)
}

// ReleaseTaobaoBusOrderSetAPIRequest 将 TaobaoBusOrderSetAPIRequest 放入 sync.Pool
func ReleaseTaobaoBusOrderSetAPIRequest(v *TaobaoBusOrderSetAPIRequest) {
	v.Reset()
	poolTaobaoBusOrderSetAPIRequest.Put(v)
}
