package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店删除接口 APIRequest
taobao.qimen.store.delete

商家在ERP等系统中调用该接口，删除线下门店
*/
type TaobaoQimenStoreDeleteRequest struct {
    model.Params

    // 要删除的门店id
    storeId   int64 

}

func NewTaobaoQimenStoreDeleteRequest() *TaobaoQimenStoreDeleteRequest{
    return &TaobaoQimenStoreDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenStoreDeleteRequest) GetApiMethodName() string {
    return "taobao.qimen.store.delete"
}

func (r TaobaoQimenStoreDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenStoreDeleteRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoQimenStoreDeleteRequest) GetStoreId() int64 {
    return r.storeId
}

