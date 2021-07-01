package store

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/store"
)

/* 
门店信息查询接口 
taobao.place.store.query

根据用户授权信息，获取用户的门店公开信息
*/
func TaobaoPlaceStoreQuery(clt *core.SDKClient, req *store.TaobaoPlaceStoreQueryAPIRequest, session string) (*store.TaobaoPlaceStoreQueryAPIResponse, error) {
    var resp store.TaobaoPlaceStoreQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
