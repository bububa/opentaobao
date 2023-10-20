package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoplacestorerelationqueryAPIRequest 门店关系查询 API请求
// taobao.place.store.relation.query
//
// 查询门店关系
type TaobaoplacestorerelationqueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramStoreRelationSimpleQuery *StoreRelationSimpleQuery
}

// NewTaobaoplacestorerelationqueryRequest 初始化TaobaoplacestorerelationqueryAPIRequest对象
func NewTaobaoplacestorerelationqueryRequest() *TaobaoplacestorerelationqueryAPIRequest {
	return &TaobaoplacestorerelationqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoplacestorerelationqueryAPIRequest) GetApiMethodName() string {
	return "taobao.place.store.relation.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoplacestorerelationqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoplacestorerelationqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamStoreRelationSimpleQuery is ParamStoreRelationSimpleQuery Setter
// 系统自动生成
func (r *TaobaoplacestorerelationqueryAPIRequest) SetParamStoreRelationSimpleQuery(_paramStoreRelationSimpleQuery *StoreRelationSimpleQuery) error {
	r._paramStoreRelationSimpleQuery = _paramStoreRelationSimpleQuery
	r.Set("param_store_relation_simple_query", _paramStoreRelationSimpleQuery)
	return nil
}

// GetParamStoreRelationSimpleQuery ParamStoreRelationSimpleQuery Getter
func (r TaobaoplacestorerelationqueryAPIRequest) GetParamStoreRelationSimpleQuery() *StoreRelationSimpleQuery {
	return r._paramStoreRelationSimpleQuery
}
