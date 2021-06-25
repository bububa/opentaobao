package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商户门店创建接口 APIRequest
taobao.place.store.create

用于商家创建线下门店
*/
type TaobaoPlaceStoreCreateRequest struct {
    model.Params

    // 门店创建入参
    storeCreate   *StoreUpdateTopDto 

}

func NewTaobaoPlaceStoreCreateRequest() *TaobaoPlaceStoreCreateRequest{
    return &TaobaoPlaceStoreCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPlaceStoreCreateRequest) GetApiMethodName() string {
    return "taobao.place.store.create"
}

func (r TaobaoPlaceStoreCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPlaceStoreCreateRequest) SetStoreCreate(storeCreate *StoreUpdateTopDto) error {
    r.storeCreate = storeCreate
    r.Set("store_create", storeCreate)
    return nil
}

func (r TaobaoPlaceStoreCreateRequest) GetStoreCreate() *StoreUpdateTopDto {
    return r.storeCreate
}

