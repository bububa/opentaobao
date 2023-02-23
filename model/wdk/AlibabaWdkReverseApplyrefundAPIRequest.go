package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkReverseApplyrefundAPIRequest 逆向申请接口 API请求
// alibaba.wdk.reverse.applyrefund
//
// 逆向渲染
type AlibabaWdkReverseApplyrefundAPIRequest struct {
	model.Params
	// 入参
	_paramApplyReverseReq *ApplyReverseReq
}

// NewAlibabaWdkReverseApplyrefundRequest 初始化AlibabaWdkReverseApplyrefundAPIRequest对象
func NewAlibabaWdkReverseApplyrefundRequest() *AlibabaWdkReverseApplyrefundAPIRequest {
	return &AlibabaWdkReverseApplyrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkReverseApplyrefundAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.reverse.applyrefund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkReverseApplyrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkReverseApplyrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamApplyReverseReq is ParamApplyReverseReq Setter
// 入参
func (r *AlibabaWdkReverseApplyrefundAPIRequest) SetParamApplyReverseReq(_paramApplyReverseReq *ApplyReverseReq) error {
	r._paramApplyReverseReq = _paramApplyReverseReq
	r.Set("param_apply_reverse_req", _paramApplyReverseReq)
	return nil
}

// GetParamApplyReverseReq ParamApplyReverseReq Getter
func (r AlibabaWdkReverseApplyrefundAPIRequest) GetParamApplyReverseReq() *ApplyReverseReq {
	return r._paramApplyReverseReq
}
