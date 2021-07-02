package store

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreRelationQueryAPIRequest) GetApiMethodName() string {
	return "taobao.place.store.relation.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreRelationQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamStoreRelationSimpleQuery Setter
// 系统自动生成
func (r *TaobaoPlaceStoreRelationQueryAPIRequest) SetParamStoreRelationSimpleQuery(_paramStoreRelationSimpleQuery *StoreRelationSimpleQuery) error {
	r._paramStoreRelationSimpleQuery = _paramStoreRelationSimpleQuery
	r.Set("param_store_relation_simple_query", _paramStoreRelationSimpleQuery)
	return nil
}

// Get ParamStoreRelationSimpleQuery Getter
func (r TaobaoPlaceStoreRelationQueryAPIRequest) GetParamStoreRelationSimpleQuery() *StoreRelationSimpleQuery {
	return r._paramStoreRelationSimpleQuery
}
