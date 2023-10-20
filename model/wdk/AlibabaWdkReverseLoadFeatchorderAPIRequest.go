package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkReverseLoadFeatchorderAPIRequest 取货详情 API请求
// alibaba.wdk.reverse.load.featchorder
//
// 取货详情
type AlibabaWdkReverseLoadFeatchorderAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramLoadFetchReq *LoadFetchReq
}

// NewAlibabaWdkReverseLoadFeatchorderRequest 初始化AlibabaWdkReverseLoadFeatchorderAPIRequest对象
func NewAlibabaWdkReverseLoadFeatchorderRequest() *AlibabaWdkReverseLoadFeatchorderAPIRequest {
	return &AlibabaWdkReverseLoadFeatchorderAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkReverseLoadFeatchorderAPIRequest) Reset() {
	r._paramLoadFetchReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkReverseLoadFeatchorderAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.reverse.load.featchorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkReverseLoadFeatchorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkReverseLoadFeatchorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamLoadFetchReq is ParamLoadFetchReq Setter
// 系统自动生成
func (r *AlibabaWdkReverseLoadFeatchorderAPIRequest) SetParamLoadFetchReq(_paramLoadFetchReq *LoadFetchReq) error {
	r._paramLoadFetchReq = _paramLoadFetchReq
	r.Set("param_load_fetch_req", _paramLoadFetchReq)
	return nil
}

// GetParamLoadFetchReq ParamLoadFetchReq Getter
func (r AlibabaWdkReverseLoadFeatchorderAPIRequest) GetParamLoadFetchReq() *LoadFetchReq {
	return r._paramLoadFetchReq
}

var poolAlibabaWdkReverseLoadFeatchorderAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkReverseLoadFeatchorderRequest()
	},
}

// GetAlibabaWdkReverseLoadFeatchorderRequest 从 sync.Pool 获取 AlibabaWdkReverseLoadFeatchorderAPIRequest
func GetAlibabaWdkReverseLoadFeatchorderAPIRequest() *AlibabaWdkReverseLoadFeatchorderAPIRequest {
	return poolAlibabaWdkReverseLoadFeatchorderAPIRequest.Get().(*AlibabaWdkReverseLoadFeatchorderAPIRequest)
}

// ReleaseAlibabaWdkReverseLoadFeatchorderAPIRequest 将 AlibabaWdkReverseLoadFeatchorderAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkReverseLoadFeatchorderAPIRequest(v *AlibabaWdkReverseLoadFeatchorderAPIRequest) {
	v.Reset()
	poolAlibabaWdkReverseLoadFeatchorderAPIRequest.Put(v)
}
