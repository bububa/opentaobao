package store

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/store"
)

/* 
商户门店库创建接口 
taobao.place.storegroup.create

用于商家创建线下门店库
*/
func TaobaoPlaceStoregroupCreate(clt *core.SDKClient, req *store.TaobaoPlaceStoregroupCreateRequest, session string) (*store.TaobaoPlaceStoregroupCreateAPIResponse, error) {
    var resp store.TaobaoPlaceStoregroupCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
