package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobustvmcreateordersetAPIRequest 线下自助机创建订单 API请求
// taobao.bus.tvmcreateorder.set
//
// 提供给汽车票线下自助机的创建订单使用
type TaobaobustvmcreateordersetAPIRequest struct {
	model.Params
	// 创建订单对象
	_paramTVMCreateOrderRQ *TvmCreateOrderRq
}

// NewTaobaobustvmcreateordersetRequest 初始化TaobaobustvmcreateordersetAPIRequest对象
func NewTaobaobustvmcreateordersetRequest() *TaobaobustvmcreateordersetAPIRequest {
	return &TaobaobustvmcreateordersetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobustvmcreateordersetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.tvmcreateorder.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobustvmcreateordersetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobustvmcreateordersetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTVMCreateOrderRQ is ParamTVMCreateOrderRQ Setter
// 创建订单对象
func (r *TaobaobustvmcreateordersetAPIRequest) SetParamTVMCreateOrderRQ(_paramTVMCreateOrderRQ *TvmCreateOrderRq) error {
	r._paramTVMCreateOrderRQ = _paramTVMCreateOrderRQ
	r.Set("param_t_v_m_create_order_r_q", _paramTVMCreateOrderRQ)
	return nil
}

// GetParamTVMCreateOrderRQ ParamTVMCreateOrderRQ Getter
func (r TaobaobustvmcreateordersetAPIRequest) GetParamTVMCreateOrderRQ() *TvmCreateOrderRq {
	return r._paramTVMCreateOrderRQ
}
