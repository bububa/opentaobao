package store

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店关系查询 APIRequest
taobao.place.store.relation.query

查询门店关系
*/
type TaobaoPlaceStoreRelationQueryRequest struct {
    model.Params

    // 系统自动生成
    paramStoreRelationSimpleQuery   *StoreRelationSimpleQuery 

}

func NewTaobaoPlaceStoreRelationQueryRequest() *TaobaoPlaceStoreRelationQueryRequest{
    return &TaobaoPlaceStoreRelationQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPlaceStoreRelationQueryRequest) GetApiMethodName() string {
    return "taobao.place.store.relation.query"
}

func (r TaobaoPlaceStoreRelationQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPlaceStoreRelationQueryRequest) SetParamStoreRelationSimpleQuery(paramStoreRelationSimpleQuery *StoreRelationSimpleQuery) error {
    r.paramStoreRelationSimpleQuery = paramStoreRelationSimpleQuery
    r.Set("param_store_relation_simple_query", paramStoreRelationSimpleQuery)
    return nil
}

func (r TaobaoPlaceStoreRelationQueryRequest) GetParamStoreRelationSimpleQuery() *StoreRelationSimpleQuery {
    return r.paramStoreRelationSimpleQuery
}

