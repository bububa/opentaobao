package store

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/store"
)

/* 
门店商品关联绑定接口 
taobao.place.store.itemstore.band

商品和多个门店关系绑定接口
*/
func TaobaoPlaceStoreItemstoreBand(clt *core.SDKClient, req *store.TaobaoPlaceStoreItemstoreBandAPIRequest, session string) (*store.TaobaoPlaceStoreItemstoreBandAPIResponse, error) {
    var resp store.TaobaoPlaceStoreItemstoreBandAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
