package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtstorerelationqueryAPIRequest 喵零门店关系查询 API请求
// tmall.nrt.store.relation.query
//
// 喵零门店关系查询
type TmallnrtstorerelationqueryAPIRequest struct {
	model.Params
	// 查询对象
	_storeQuery *StoreQuery
}

// NewTmallnrtstorerelationqueryRequest 初始化TmallnrtstorerelationqueryAPIRequest对象
func NewTmallnrtstorerelationqueryRequest() *TmallnrtstorerelationqueryAPIRequest {
	return &TmallnrtstorerelationqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrtstorerelationqueryAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.store.relation.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrtstorerelationqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrtstorerelationqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreQuery is StoreQuery Setter
// 查询对象
func (r *TmallnrtstorerelationqueryAPIRequest) SetStoreQuery(_storeQuery *StoreQuery) error {
	r._storeQuery = _storeQuery
	r.Set("store_query", _storeQuery)
	return nil
}

// GetStoreQuery StoreQuery Getter
func (r TmallnrtstorerelationqueryAPIRequest) GetStoreQuery() *StoreQuery {
	return r._storeQuery
}
