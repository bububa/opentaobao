package store

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/store"
)

/* 
商户门店拓展信息更新接口 
taobao.place.store.extend.update

更新商户门店拓展信息（tags、attribute、bizAtrribute）更新接口
*/
func TaobaoPlaceStoreExtendUpdate(clt *core.SDKClient, req *store.TaobaoPlaceStoreExtendUpdateAPIRequest, session string) (*store.TaobaoPlaceStoreExtendUpdateAPIResponse, error) {
    var resp store.TaobaoPlaceStoreExtendUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
