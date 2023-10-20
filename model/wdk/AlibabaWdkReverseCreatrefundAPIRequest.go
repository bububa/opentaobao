package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkReverseCreatrefundAPIRequest 逆向提交 API请求
// alibaba.wdk.reverse.creatrefund
//
// 逆向申请提交
type AlibabaWdkReverseCreatrefundAPIRequest struct {
	model.Params
	// CreateReverseReq
	_paramCreateReverseReq *CreateReverseReq
}

// NewAlibabaWdkReverseCreatrefundRequest 初始化AlibabaWdkReverseCreatrefundAPIRequest对象
func NewAlibabaWdkReverseCreatrefundRequest() *AlibabaWdkReverseCreatrefundAPIRequest {
	return &AlibabaWdkReverseCreatrefundAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkReverseCreatrefundAPIRequest) Reset() {
	r._paramCreateReverseReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkReverseCreatrefundAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.reverse.creatrefund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkReverseCreatrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkReverseCreatrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCreateReverseReq is ParamCreateReverseReq Setter
// CreateReverseReq
func (r *AlibabaWdkReverseCreatrefundAPIRequest) SetParamCreateReverseReq(_paramCreateReverseReq *CreateReverseReq) error {
	r._paramCreateReverseReq = _paramCreateReverseReq
	r.Set("param_create_reverse_req", _paramCreateReverseReq)
	return nil
}

// GetParamCreateReverseReq ParamCreateReverseReq Getter
func (r AlibabaWdkReverseCreatrefundAPIRequest) GetParamCreateReverseReq() *CreateReverseReq {
	return r._paramCreateReverseReq
}

var poolAlibabaWdkReverseCreatrefundAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkReverseCreatrefundRequest()
	},
}

// GetAlibabaWdkReverseCreatrefundRequest 从 sync.Pool 获取 AlibabaWdkReverseCreatrefundAPIRequest
func GetAlibabaWdkReverseCreatrefundAPIRequest() *AlibabaWdkReverseCreatrefundAPIRequest {
	return poolAlibabaWdkReverseCreatrefundAPIRequest.Get().(*AlibabaWdkReverseCreatrefundAPIRequest)
}

// ReleaseAlibabaWdkReverseCreatrefundAPIRequest 将 AlibabaWdkReverseCreatrefundAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkReverseCreatrefundAPIRequest(v *AlibabaWdkReverseCreatrefundAPIRequest) {
	v.Reset()
	poolAlibabaWdkReverseCreatrefundAPIRequest.Put(v)
}
