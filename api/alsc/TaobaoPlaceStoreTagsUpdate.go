package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
门店打标去标 
taobao.place.store.tags.update

门店打标去标
*/
func TaobaoPlaceStoreTagsUpdate(clt *core.SDKClient, req *alsc.TaobaoPlaceStoreTagsUpdateAPIRequest, session string) (*alsc.TaobaoPlaceStoreTagsUpdateAPIResponse, error) {
    var resp alsc.TaobaoPlaceStoreTagsUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
