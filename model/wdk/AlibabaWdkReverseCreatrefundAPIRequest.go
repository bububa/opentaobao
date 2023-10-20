package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkreversecreatrefundAPIRequest 逆向提交 API请求
// alibaba.wdk.reverse.creatrefund
//
// 逆向申请提交
type AlibabawdkreversecreatrefundAPIRequest struct {
	model.Params
	// CreateReverseReq
	_paramCreateReverseReq *CreateReverseReq
}

// NewAlibabawdkreversecreatrefundRequest 初始化AlibabawdkreversecreatrefundAPIRequest对象
func NewAlibabawdkreversecreatrefundRequest() *AlibabawdkreversecreatrefundAPIRequest {
	return &AlibabawdkreversecreatrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkreversecreatrefundAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.reverse.creatrefund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkreversecreatrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkreversecreatrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCreateReverseReq is ParamCreateReverseReq Setter
// CreateReverseReq
func (r *AlibabawdkreversecreatrefundAPIRequest) SetParamCreateReverseReq(_paramCreateReverseReq *CreateReverseReq) error {
	r._paramCreateReverseReq = _paramCreateReverseReq
	r.Set("param_create_reverse_req", _paramCreateReverseReq)
	return nil
}

// GetParamCreateReverseReq ParamCreateReverseReq Getter
func (r AlibabawdkreversecreatrefundAPIRequest) GetParamCreateReverseReq() *CreateReverseReq {
	return r._paramCreateReverseReq
}
