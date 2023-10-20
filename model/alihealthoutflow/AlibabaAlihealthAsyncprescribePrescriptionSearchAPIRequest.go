package alihealthoutflow

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
