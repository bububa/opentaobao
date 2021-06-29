package store

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/store"
)

/* 
门店和子门店关系删除 
taobao.place.storerelatesub.delete

门店和子门店关系删除
*/
func TaobaoPlaceStorerelatesubDelete(clt *core.SDKClient, req *store.TaobaoPlaceStorerelatesubDeleteRequest, session string) (*store.TaobaoPlaceStorerelatesubDeleteAPIResponse, error) {
    var resp store.TaobaoPlaceStorerelatesubDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
