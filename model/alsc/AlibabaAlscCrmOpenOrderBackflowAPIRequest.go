package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmopenorderbackflowAPIRequest 订单回流接口 API请求
// alibaba.alsc.crm.open.order.backflow
//
// 回流isv订单接口
type AlibabaalsccrmopenorderbackflowAPIRequest struct {
	model.Params
	// 入参
	_paramOrderBackflowOpenReq *OrderBackflowOpenReq
}

// NewAlibabaalsccrmopenorderbackflowRequest 初始化AlibabaalsccrmopenorderbackflowAPIRequest对象
func NewAlibabaalsccrmopenorderbackflowRequest() *AlibabaalsccrmopenorderbackflowAPIRequest {
	return &AlibabaalsccrmopenorderbackflowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmopenorderbackflowAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.open.order.backflow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmopenorderbackflowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmopenorderbackflowAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOrderBackflowOpenReq is ParamOrderBackflowOpenReq Setter
// 入参
func (r *AlibabaalsccrmopenorderbackflowAPIRequest) SetParamOrderBackflowOpenReq(_paramOrderBackflowOpenReq *OrderBackflowOpenReq) error {
	r._paramOrderBackflowOpenReq = _paramOrderBackflowOpenReq
	r.Set("param_order_backflow_open_req", _paramOrderBackflowOpenReq)
	return nil
}

// GetParamOrderBackflowOpenReq ParamOrderBackflowOpenReq Getter
func (r AlibabaalsccrmopenorderbackflowAPIRequest) GetParamOrderBackflowOpenReq() *OrderBackflowOpenReq {
	return r._paramOrderBackflowOpenReq
}
