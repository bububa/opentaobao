package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
新增门店扩展属性 
taobao.place.store.extend.add

新增授权用户的门店扩展属性
*/
func TaobaoPlaceStoreExtendAdd(clt *core.SDKClient, req *alsc.TaobaoPlaceStoreExtendAddRequest, session string) (*alsc.TaobaoPlaceStoreExtendAddAPIResponse, error) {
    var resp alsc.TaobaoPlaceStoreExtendAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
