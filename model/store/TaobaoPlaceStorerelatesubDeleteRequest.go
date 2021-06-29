package store

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店和子门店关系删除 API请求
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

// 初始化TaobaoPlaceStorerelatesubDeleteRequest对象
func NewTaobaoPlaceStorerelatesubDeleteRequest() *TaobaoPlaceStorerelatesubDeleteRequest{
    return &TaobaoPlaceStorerelatesubDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStorerelatesubDeleteRequest) GetApiMethodName() string {
    return "taobao.place.storerelatesub.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStorerelatesubDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店Id
func (r *TaobaoPlaceStorerelatesubDeleteRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r TaobaoPlaceStorerelatesubDeleteRequest) GetStoreId() int64 {
    return r.storeId
}
// SubStoreIds Setter
// 子门店id
func (r *TaobaoPlaceStorerelatesubDeleteRequest) SetSubStoreIds(subStoreIds []int64) error {
    r.subStoreIds = subStoreIds
    r.Set("sub_store_ids", subStoreIds)
    return nil
}

// SubStoreIds Getter
func (r TaobaoPlaceStorerelatesubDeleteRequest) GetSubStoreIds() []int64 {
    return r.subStoreIds
}
