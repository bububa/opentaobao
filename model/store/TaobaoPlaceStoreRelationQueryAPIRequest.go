package store

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoreRelationQueryAPIRequest 门店关系查询 API请求
// taobao.place.store.relation.query
//
// 查询门店关系
type TaobaoPlaceStoreRelationQueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramStoreRelationSimpleQuery *StoreRelationSimpleQuery
}

// NewTaobaoPlaceStoreRelationQueryRequest 初始化TaobaoPlaceStoreRelationQueryAPIRequest对象
func NewTaobaoPlaceStoreRelationQueryRequest() *TaobaoPlaceStoreRelationQueryAPIRequest {
	return &TaobaoPlaceStoreRelationQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPlaceStoreRelationQueryAPIRequest) Reset() {
	r._paramStoreRelationSimpleQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreRelationQueryAPIRequest) GetApiMethodName() string {
	return "taobao.place.store.relation.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreRelationQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPlaceStoreRelationQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamStoreRelationSimpleQuery is ParamStoreRelationSimpleQuery Setter
// 系统自动生成
func (r *TaobaoPlaceStoreRelationQueryAPIRequest) SetParamStoreRelationSimpleQuery(_paramStoreRelationSimpleQuery *StoreRelationSimpleQuery) error {
	r._paramStoreRelationSimpleQuery = _paramStoreRelationSimpleQuery
	r.Set("param_store_relation_simple_query", _paramStoreRelationSimpleQuery)
	return nil
}

// GetParamStoreRelationSimpleQuery ParamStoreRelationSimpleQuery Getter
func (r TaobaoPlaceStoreRelationQueryAPIRequest) GetParamStoreRelationSimpleQuery() *StoreRelationSimpleQuery {
	return r._paramStoreRelationSimpleQuery
}

var poolTaobaoPlaceStoreRelationQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPlaceStoreRelationQueryRequest()
	},
}

// GetTaobaoPlaceStoreRelationQueryRequest 从 sync.Pool 获取 TaobaoPlaceStoreRelationQueryAPIRequest
func GetTaobaoPlaceStoreRelationQueryAPIRequest() *TaobaoPlaceStoreRelationQueryAPIRequest {
	return poolTaobaoPlaceStoreRelationQueryAPIRequest.Get().(*TaobaoPlaceStoreRelationQueryAPIRequest)
}

// ReleaseTaobaoPlaceStoreRelationQueryAPIRequest 将 TaobaoPlaceStoreRelationQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoPlaceStoreRelationQueryAPIRequest(v *TaobaoPlaceStoreRelationQueryAPIRequest) {
	v.Reset()
	poolTaobaoPlaceStoreRelationQueryAPIRequest.Put(v)
}
