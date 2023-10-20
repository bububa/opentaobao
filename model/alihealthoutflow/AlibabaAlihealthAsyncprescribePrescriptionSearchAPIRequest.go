package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthasyncprescribeprescriptionsearchAPIRequest 异步开方处方查询 API请求
// alibaba.alihealth.asyncprescribe.prescription.search
//
// 异步开方处方查询
type AlibabaalihealthasyncprescribeprescriptionsearchAPIRequest struct {
	model.Params
	// 查询入参
	_searchRequest *AsyncPrescribeSearchRequest
}

// NewAlibabaalihealthasyncprescribeprescriptionsearchRequest 初始化AlibabaalihealthasyncprescribeprescriptionsearchAPIRequest对象
func NewAlibabaalihealthasyncprescribeprescriptionsearchRequest() *AlibabaalihealthasyncprescribeprescriptionsearchAPIRequest {
	return &AlibabaalihealthasyncprescribeprescriptionsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthasyncprescribeprescriptionsearchAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.asyncprescribe.prescription.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthasyncprescribeprescriptionsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthasyncprescribeprescriptionsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSearchRequest is SearchRequest Setter
// 查询入参
func (r *AlibabaalihealthasyncprescribeprescriptionsearchAPIRequest) SetSearchRequest(_searchRequest *AsyncPrescribeSearchRequest) error {
	r._searchRequest = _searchRequest
	r.Set("search_request", _searchRequest)
	return nil
}

// GetSearchRequest SearchRequest Getter
func (r AlibabaalihealthasyncprescribeprescriptionsearchAPIRequest) GetSearchRequest() *AsyncPrescribeSearchRequest {
	return r._searchRequest
}
