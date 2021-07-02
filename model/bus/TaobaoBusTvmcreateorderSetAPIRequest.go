package bus

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusTvmcreateorderSetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.tvmcreateorder.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusTvmcreateorderSetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamTVMCreateOrderRQ Setter
// 创建订单对象
func (r *TaobaoBusTvmcreateorderSetAPIRequest) SetParamTVMCreateOrderRQ(_paramTVMCreateOrderRQ *TvmCreateOrderRq) error {
	r._paramTVMCreateOrderRQ = _paramTVMCreateOrderRQ
	r.Set("param_t_v_m_create_order_r_q", _paramTVMCreateOrderRQ)
	return nil
}

// Get ParamTVMCreateOrderRQ Getter
func (r TaobaoBusTvmcreateorderSetAPIRequest) GetParamTVMCreateOrderRQ() *TvmCreateOrderRq {
	return r._paramTVMCreateOrderRQ
}
