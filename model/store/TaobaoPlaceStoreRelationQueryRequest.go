package store

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店关系查询 API请求
taobao.place.store.relation.query

查询门店关系
*/
type TaobaoPlaceStoreRelationQueryRequest struct {
    model.Params
    // 系统自动生成
    paramStoreRelationSimpleQuery   *StoreRelationSimpleQuery
}

// 初始化TaobaoPlaceStoreRelationQueryRequest对象
func NewTaobaoPlaceStoreRelationQueryRequest() *TaobaoPlaceStoreRelationQueryRequest{
    return &TaobaoPlaceStoreRelationQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreRelationQueryRequest) GetApiMethodName() string {
    return "taobao.place.store.relation.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreRelationQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamStoreRelationSimpleQuery Setter
// 系统自动生成
func (r *TaobaoPlaceStoreRelationQueryRequest) SetParamStoreRelationSimpleQuery(paramStoreRelationSimpleQuery *StoreRelationSimpleQuery) error {
    r.paramStoreRelationSimpleQuery = paramStoreRelationSimpleQuery
    r.Set("param_store_relation_simple_query", paramStoreRelationSimpleQuery)
    return nil
}

// ParamStoreRelationSimpleQuery Getter
func (r TaobaoPlaceStoreRelationQueryRequest) GetParamStoreRelationSimpleQuery() *StoreRelationSimpleQuery {
    return r.paramStoreRelationSimpleQuery
}
