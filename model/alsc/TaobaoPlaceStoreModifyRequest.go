package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家修改线下门店 APIRequest
taobao.place.store.modify

用于商家修改线下门店信息
*/
type TaobaoPlaceStoreModifyRequest struct {
    model.Params

    // 门店创建入参
    storeUpdate   *StoreUpdateTopDto 

}

func NewTaobaoPlaceStoreModifyRequest() *TaobaoPlaceStoreModifyRequest{
    return &TaobaoPlaceStoreModifyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPlaceStoreModifyRequest) GetApiMethodName() string {
    return "taobao.place.store.modify"
}

func (r TaobaoPlaceStoreModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPlaceStoreModifyRequest) SetStoreUpdate(storeUpdate *StoreUpdateTopDto) error {
    r.storeUpdate = storeUpdate
    r.Set("store_update", storeUpdate)
    return nil
}

func (r TaobaoPlaceStoreModifyRequest) GetStoreUpdate() *StoreUpdateTopDto {
    return r.storeUpdate
}

