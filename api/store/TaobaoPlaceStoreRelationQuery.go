package store

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/store"
)

/* 
门店关系查询 
taobao.place.store.relation.query

查询门店关系
*/
func TaobaoPlaceStoreRelationQuery(clt *core.SDKClient, req *store.TaobaoPlaceStoreRelationQueryRequest, session string) (*store.TaobaoPlaceStoreRelationQueryAPIResponse, error) {
    var resp store.TaobaoPlaceStoreRelationQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
