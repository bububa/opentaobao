package store

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店和子门店关系删除 APIRequest
taobao.place.storerelatesub.delete

门店和子门店关系删除
*/
type TaobaoPlaceStorerelatesubDeleteRequest struct {
    model.Params

    // 门店Id
    storeId   int64 

    // 子门店id
    subStoreIds   []int64 

}

func NewTaobaoPlaceStorerelatesubDeleteRequest() *TaobaoPlaceStorerelatesubDeleteRequest{
    return &TaobaoPlaceStorerelatesubDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPlaceStorerelatesubDeleteRequest) GetApiMethodName() string {
    return "taobao.place.storerelatesub.delete"
}

func (r TaobaoPlaceStorerelatesubDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPlaceStorerelatesubDeleteRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoPlaceStorerelatesubDeleteRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoPlaceStorerelatesubDeleteRequest) SetSubStoreIds(subStoreIds []int64) error {
    r.subStoreIds = subStoreIds
    r.Set("sub_store_ids", subStoreIds)
    return nil
}

func (r TaobaoPlaceStorerelatesubDeleteRequest) GetSubStoreIds() []int64 {
    return r.subStoreIds
}

