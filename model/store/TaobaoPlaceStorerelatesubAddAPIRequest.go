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
type TaobaoPlaceStorerelatesubAddAPIRequest struct {
    model.Params
    // 门店Id
    _storeId   int64
    // 子门店Id
    _subStoreIds   []int64
}

// 初始化TaobaoPlaceStorerelatesubAddAPIRequest对象
func NewTaobaoPlaceStorerelatesubAddRequest() *TaobaoPlaceStorerelatesubAddAPIRequest{
    return &TaobaoPlaceStorerelatesubAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStorerelatesubAddAPIRequest) GetApiMethodName() string {
    return "taobao.place.storerelatesub.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStorerelatesubAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店Id
func (r *TaobaoPlaceStorerelatesubAddAPIRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoPlaceStorerelatesubAddAPIRequest) GetStoreId() int64 {
    return r._storeId
}
// SubStoreIds Setter
// 子门店Id
func (r *TaobaoPlaceStorerelatesubAddAPIRequest) SetSubStoreIds(_subStoreIds []int64) error {
    r._subStoreIds = _subStoreIds
    r.Set("sub_store_ids", _subStoreIds)
    return nil
}

// SubStoreIds Getter
func (r TaobaoPlaceStorerelatesubAddAPIRequest) GetSubStoreIds() []int64 {
    return r._subStoreIds
}
