package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下门店删除 APIRequest
taobao.place.store.delete

用于商家删除线下门店
*/
type TaobaoPlaceStoreDeleteRequest struct {
    model.Params

    // 门店id
    storeId   int64 

}

func NewTaobaoPlaceStoreDeleteRequest() *TaobaoPlaceStoreDeleteRequest{
    return &TaobaoPlaceStoreDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPlaceStoreDeleteRequest) GetApiMethodName() string {
    return "taobao.place.store.delete"
}

func (r TaobaoPlaceStoreDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPlaceStoreDeleteRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoPlaceStoreDeleteRequest) GetStoreId() int64 {
    return r.storeId
}

