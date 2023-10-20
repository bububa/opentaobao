package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkReverseCreatefeatchAPIRequest 逆向取货 API请求
// alibaba.wdk.reverse.createfeatch
//
// 发起逆向取货
type AlibabaWdkReverseCreatefeatchAPIRequest struct {
	model.Params
	// 入参
	_paramCreateFetchReq *CreateFetchReq
}

// NewAlibabaWdkReverseCreatefeatchRequest 初始化AlibabaWdkReverseCreatefeatchAPIRequest对象
func NewAlibabaWdkReverseCreatefeatchRequest() *AlibabaWdkReverseCreatefeatchAPIRequest {
	return &AlibabaWdkReverseCreatefeatchAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkReverseCreatefeatchAPIRequest) Reset() {
	r._paramCreateFetchReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkReverseCreatefeatchAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.reverse.createfeatch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkReverseCreatefeatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkReverseCreatefeatchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCreateFetchReq is ParamCreateFetchReq Setter
// 入参
func (r *AlibabaWdkReverseCreatefeatchAPIRequest) SetParamCreateFetchReq(_paramCreateFetchReq *CreateFetchReq) error {
	r._paramCreateFetchReq = _paramCreateFetchReq
	r.Set("param_create_fetch_req", _paramCreateFetchReq)
	return nil
}

// GetParamCreateFetchReq ParamCreateFetchReq Getter
func (r AlibabaWdkReverseCreatefeatchAPIRequest) GetParamCreateFetchReq() *CreateFetchReq {
	return r._paramCreateFetchReq
}

var poolAlibabaWdkReverseCreatefeatchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkReverseCreatefeatchRequest()
	},
}

// GetAlibabaWdkReverseCreatefeatchRequest 从 sync.Pool 获取 AlibabaWdkReverseCreatefeatchAPIRequest
func GetAlibabaWdkReverseCreatefeatchAPIRequest() *AlibabaWdkReverseCreatefeatchAPIRequest {
	return poolAlibabaWdkReverseCreatefeatchAPIRequest.Get().(*AlibabaWdkReverseCreatefeatchAPIRequest)
}

// ReleaseAlibabaWdkReverseCreatefeatchAPIRequest 将 AlibabaWdkReverseCreatefeatchAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkReverseCreatefeatchAPIRequest(v *AlibabaWdkReverseCreatefeatchAPIRequest) {
	v.Reset()
	poolAlibabaWdkReverseCreatefeatchAPIRequest.Put(v)
}
