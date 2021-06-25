package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
商家修改线下门店 
taobao.place.store.modify

用于商家修改线下门店信息
*/
func TaobaoPlaceStoreModify(clt *core.SDKClient, req *alsc.TaobaoPlaceStoreModifyRequest, session string) (*alsc.TaobaoPlaceStoreModifyResponse, error) {
    var resp alsc.TaobaoPlaceStoreModifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
