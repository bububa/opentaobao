package store

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/store"
)

/* 
门店和子门店关系新增 
taobao.place.storerelatesub.add

门店和子门店关系新增
*/
func TaobaoPlaceStorerelatesubAdd(clt *core.SDKClient, req *store.TaobaoPlaceStorerelatesubAddRequest, session string) (*store.TaobaoPlaceStorerelatesubAddAPIResponse, error) {
    var resp store.TaobaoPlaceStorerelatesubAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
