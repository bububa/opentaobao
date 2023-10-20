package alihealthoutflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest 异步开方处方查询 API请求
// alibaba.alihealth.asyncprescribe.prescription.search
//
// 异步开方处方查询
type AlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest struct {
	model.Params
	// 查询入参
	_searchRequest *AsyncPrescribeSearchRequest
}

// NewAlibabaAlihealthAsyncprescribePrescriptionSearchRequest 初始化AlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest对象
func NewAlibabaAlihealthAsyncprescribePrescriptionSearchRequest() *AlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest {
	return &AlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest) Reset() {
	r._searchRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.asyncprescribe.prescription.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSearchRequest is SearchRequest Setter
// 查询入参
func (r *AlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest) SetSearchRequest(_searchRequest *AsyncPrescribeSearchRequest) error {
	r._searchRequest = _searchRequest
	r.Set("search_request", _searchRequest)
	return nil
}

// GetSearchRequest SearchRequest Getter
func (r AlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest) GetSearchRequest() *AsyncPrescribeSearchRequest {
	return r._searchRequest
}

var poolAlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthAsyncprescribePrescriptionSearchRequest()
	},
}

// GetAlibabaAlihealthAsyncprescribePrescriptionSearchRequest 从 sync.Pool 获取 AlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest
func GetAlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest() *AlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest {
	return poolAlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest.Get().(*AlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest)
}

// ReleaseAlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest 将 AlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest(v *AlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest.Put(v)
}
