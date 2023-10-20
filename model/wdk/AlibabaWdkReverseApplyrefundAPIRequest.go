package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkreverseapplyrefundAPIRequest 逆向申请接口 API请求
// alibaba.wdk.reverse.applyrefund
//
// 逆向渲染
type AlibabawdkreverseapplyrefundAPIRequest struct {
	model.Params
	// 入参
	_paramApplyReverseReq *ApplyReverseReq
}

// NewAlibabawdkreverseapplyrefundRequest 初始化AlibabawdkreverseapplyrefundAPIRequest对象
func NewAlibabawdkreverseapplyrefundRequest() *AlibabawdkreverseapplyrefundAPIRequest {
	return &AlibabawdkreverseapplyrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkreverseapplyrefundAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.reverse.applyrefund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkreverseapplyrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkreverseapplyrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamApplyReverseReq is ParamApplyReverseReq Setter
// 入参
func (r *AlibabawdkreverseapplyrefundAPIRequest) SetParamApplyReverseReq(_paramApplyReverseReq *ApplyReverseReq) error {
	r._paramApplyReverseReq = _paramApplyReverseReq
	r.Set("param_apply_reverse_req", _paramApplyReverseReq)
	return nil
}

// GetParamApplyReverseReq ParamApplyReverseReq Getter
func (r AlibabawdkreverseapplyrefundAPIRequest) GetParamApplyReverseReq() *ApplyReverseReq {
	return r._paramApplyReverseReq
}
