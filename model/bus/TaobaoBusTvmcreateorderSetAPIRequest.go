package bus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusTvmcreateorderSetAPIRequest 线下自助机创建订单 API请求
// taobao.bus.tvmcreateorder.set
//
// 提供给汽车票线下自助机的创建订单使用
type TaobaoBusTvmcreateorderSetAPIRequest struct {
	model.Params
	// 创建订单对象
	_paramTVMCreateOrderRQ *TvmCreateOrderRq
}

// NewTaobaoBusTvmcreateorderSetRequest 初始化TaobaoBusTvmcreateorderSetAPIRequest对象
func NewTaobaoBusTvmcreateorderSetRequest() *TaobaoBusTvmcreateorderSetAPIRequest {
	return &TaobaoBusTvmcreateorderSetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBusTvmcreateorderSetAPIRequest) Reset() {
	r._paramTVMCreateOrderRQ = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusTvmcreateorderSetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.tvmcreateorder.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusTvmcreateorderSetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusTvmcreateorderSetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTVMCreateOrderRQ is ParamTVMCreateOrderRQ Setter
// 创建订单对象
func (r *TaobaoBusTvmcreateorderSetAPIRequest) SetParamTVMCreateOrderRQ(_paramTVMCreateOrderRQ *TvmCreateOrderRq) error {
	r._paramTVMCreateOrderRQ = _paramTVMCreateOrderRQ
	r.Set("param_t_v_m_create_order_r_q", _paramTVMCreateOrderRQ)
	return nil
}

// GetParamTVMCreateOrderRQ ParamTVMCreateOrderRQ Getter
func (r TaobaoBusTvmcreateorderSetAPIRequest) GetParamTVMCreateOrderRQ() *TvmCreateOrderRq {
	return r._paramTVMCreateOrderRQ
}

var poolTaobaoBusTvmcreateorderSetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBusTvmcreateorderSetRequest()
	},
}

// GetTaobaoBusTvmcreateorderSetRequest 从 sync.Pool 获取 TaobaoBusTvmcreateorderSetAPIRequest
func GetTaobaoBusTvmcreateorderSetAPIRequest() *TaobaoBusTvmcreateorderSetAPIRequest {
	return poolTaobaoBusTvmcreateorderSetAPIRequest.Get().(*TaobaoBusTvmcreateorderSetAPIRequest)
}

// ReleaseTaobaoBusTvmcreateorderSetAPIRequest 将 TaobaoBusTvmcreateorderSetAPIRequest 放入 sync.Pool
func ReleaseTaobaoBusTvmcreateorderSetAPIRequest(v *TaobaoBusTvmcreateorderSetAPIRequest) {
	v.Reset()
	poolTaobaoBusTvmcreateorderSetAPIRequest.Put(v)
}
