package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkReverseApplyrefundAPIRequest) Reset() {
	r._paramApplyReverseReq = nil
	r.Params.ToZero()
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

var poolAlibabaWdkReverseApplyrefundAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkReverseApplyrefundRequest()
	},
}

// GetAlibabaWdkReverseApplyrefundRequest 从 sync.Pool 获取 AlibabaWdkReverseApplyrefundAPIRequest
func GetAlibabaWdkReverseApplyrefundAPIRequest() *AlibabaWdkReverseApplyrefundAPIRequest {
	return poolAlibabaWdkReverseApplyrefundAPIRequest.Get().(*AlibabaWdkReverseApplyrefundAPIRequest)
}

// ReleaseAlibabaWdkReverseApplyrefundAPIRequest 将 AlibabaWdkReverseApplyrefundAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkReverseApplyrefundAPIRequest(v *AlibabaWdkReverseApplyrefundAPIRequest) {
	v.Reset()
	poolAlibabaWdkReverseApplyrefundAPIRequest.Put(v)
}
