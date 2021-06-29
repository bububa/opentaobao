package store

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店和子门店关系新增 APIRequest
taobao.place.storerelatesub.add

门店和子门店关系新增
*/
type TaobaoPlaceStorerelatesubAddRequest struct {
    model.Params

    // 门店Id
    storeId   int64 

    // 子门店Id
    subStoreIds   []int64 

}

func NewTaobaoPlaceStorerelatesubAddRequest() *TaobaoPlaceStorerelatesubAddRequest{
    return &TaobaoPlaceStorerelatesubAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPlaceStorerelatesubAddRequest) GetApiMethodName() string {
    return "taobao.place.storerelatesub.add"
}

func (r TaobaoPlaceStorerelatesubAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPlaceStorerelatesubAddRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoPlaceStorerelatesubAddRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoPlaceStorerelatesubAddRequest) SetSubStoreIds(subStoreIds []int64) error {
    r.subStoreIds = subStoreIds
    r.Set("sub_store_ids", subStoreIds)
    return nil
}

func (r TaobaoPlaceStorerelatesubAddRequest) GetSubStoreIds() []int64 {
    return r.subStoreIds
}

