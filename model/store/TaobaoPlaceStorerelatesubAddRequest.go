package store

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店和子门店关系新增 API请求
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

// 初始化TaobaoPlaceStorerelatesubAddRequest对象
func NewTaobaoPlaceStorerelatesubAddRequest() *TaobaoPlaceStorerelatesubAddRequest{
    return &TaobaoPlaceStorerelatesubAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStorerelatesubAddRequest) GetApiMethodName() string {
    return "taobao.place.storerelatesub.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStorerelatesubAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店Id
func (r *TaobaoPlaceStorerelatesubAddRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r TaobaoPlaceStorerelatesubAddRequest) GetStoreId() int64 {
    return r.storeId
}
// SubStoreIds Setter
// 子门店Id
func (r *TaobaoPlaceStorerelatesubAddRequest) SetSubStoreIds(subStoreIds []int64) error {
    r.subStoreIds = subStoreIds
    r.Set("sub_store_ids", subStoreIds)
    return nil
}

// SubStoreIds Getter
func (r TaobaoPlaceStorerelatesubAddRequest) GetSubStoreIds() []int64 {
    return r.subStoreIds
}
