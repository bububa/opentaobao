package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOcOrderCreateAPIRequest
创建OC订单 API请求
taobao.oc.order.create

创建OC订单接口 */
type TaobaoOcOrderCreateAPIRequest struct {
	model.Params
	// OC订单
	_paramOCOrder *OcOrder
}

// NewTaobaoOcOrderCreateRequest 初始化TaobaoOcOrderCreateAPIRequest对象
func NewTaobaoOcOrderCreateRequest() *TaobaoOcOrderCreateAPIRequest {
	return &TaobaoOcOrderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOcOrderCreateAPIRequest) GetApiMethodName() string {
	return "taobao.oc.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOcOrderCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamOCOrder Setter
// OC订单
func (r *TaobaoOcOrderCreateAPIRequest) SetParamOCOrder(_paramOCOrder *OcOrder) error {
	r._paramOCOrder = _paramOCOrder
	r.Set("param_o_c_order", _paramOCOrder)
	return nil
}

// Get ParamOCOrder Getter
func (r TaobaoOcOrderCreateAPIRequest) GetParamOCOrder() *OcOrder {
	return r._paramOCOrder
}
