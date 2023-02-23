package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkAxStoreQueryAPIRequest 翱象经营店查询接口 API请求
// alibaba.wdk.ax.store.query
//
// 翱象经营店查询接口
type AlibabaWdkAxStoreQueryAPIRequest struct {
	model.Params
	// 查询入参
	_queryRequest *AxStoreQueryRequest
}

// NewAlibabaWdkAxStoreQueryRequest 初始化AlibabaWdkAxStoreQueryAPIRequest对象
func NewAlibabaWdkAxStoreQueryRequest() *AlibabaWdkAxStoreQueryAPIRequest {
	return &AlibabaWdkAxStoreQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkAxStoreQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ax.store.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkAxStoreQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkAxStoreQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryRequest is QueryRequest Setter
// 查询入参
func (r *AlibabaWdkAxStoreQueryAPIRequest) SetQueryRequest(_queryRequest *AxStoreQueryRequest) error {
	r._queryRequest = _queryRequest
	r.Set("query_request", _queryRequest)
	return nil
}

// GetQueryRequest QueryRequest Getter
func (r AlibabaWdkAxStoreQueryAPIRequest) GetQueryRequest() *AxStoreQueryRequest {
	return r._queryRequest
}
