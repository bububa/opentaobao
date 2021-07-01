package store

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/store"
)

/* 
门店和子门店关系查找 
taobao.place.storerelatesub.get

门店和子门店关系查找
*/
func TaobaoPlaceStorerelatesubGet(clt *core.SDKClient, req *store.TaobaoPlaceStorerelatesubGetAPIRequest, session string) (*store.TaobaoPlaceStorerelatesubGetAPIResponse, error) {
    var resp store.TaobaoPlaceStorerelatesubGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
