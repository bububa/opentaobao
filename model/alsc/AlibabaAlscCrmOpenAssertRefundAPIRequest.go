package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmopenassertrefundAPIRequest 资产核销回退接口 API请求
// alibaba.alsc.crm.open.assert.refund
//
// 回退已经核销的储值积分券资产
type AlibabaalsccrmopenassertrefundAPIRequest struct {
	model.Params
	// 入参
	_paramPropertyRefundOpenReq *PropertyRefundOpenReq
}

// NewAlibabaalsccrmopenassertrefundRequest 初始化AlibabaalsccrmopenassertrefundAPIRequest对象
func NewAlibabaalsccrmopenassertrefundRequest() *AlibabaalsccrmopenassertrefundAPIRequest {
	return &AlibabaalsccrmopenassertrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmopenassertrefundAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.open.assert.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmopenassertrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmopenassertrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPropertyRefundOpenReq is ParamPropertyRefundOpenReq Setter
// 入参
func (r *AlibabaalsccrmopenassertrefundAPIRequest) SetParamPropertyRefundOpenReq(_paramPropertyRefundOpenReq *PropertyRefundOpenReq) error {
	r._paramPropertyRefundOpenReq = _paramPropertyRefundOpenReq
	r.Set("param_property_refund_open_req", _paramPropertyRefundOpenReq)
	return nil
}

// GetParamPropertyRefundOpenReq ParamPropertyRefundOpenReq Getter
func (r AlibabaalsccrmopenassertrefundAPIRequest) GetParamPropertyRefundOpenReq() *PropertyRefundOpenReq {
	return r._paramPropertyRefundOpenReq
}
