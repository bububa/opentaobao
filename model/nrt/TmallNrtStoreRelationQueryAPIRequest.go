package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtStoreRelationQueryAPIRequest 喵零门店关系查询 API请求
// tmall.nrt.store.relation.query
//
// 喵零门店关系查询
type TmallNrtStoreRelationQueryAPIRequest struct {
	model.Params
	// 查询对象
	_storeQuery *StoreQuery
}

// NewTmallNrtStoreRelationQueryRequest 初始化TmallNrtStoreRelationQueryAPIRequest对象
func NewTmallNrtStoreRelationQueryRequest() *TmallNrtStoreRelationQueryAPIRequest {
	return &TmallNrtStoreRelationQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtStoreRelationQueryAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.store.relation.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtStoreRelationQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStoreQuery is StoreQuery Setter
// 查询对象
func (r *TmallNrtStoreRelationQueryAPIRequest) SetStoreQuery(_storeQuery *StoreQuery) error {
	r._storeQuery = _storeQuery
	r.Set("store_query", _storeQuery)
	return nil
}

// GetStoreQuery StoreQuery Getter
func (r TmallNrtStoreRelationQueryAPIRequest) GetStoreQuery() *StoreQuery {
	return r._storeQuery
}
