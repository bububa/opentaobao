package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下门店删除 API请求
taobao.place.store.delete

用于商家删除线下门店
*/
type TaobaoPlaceStoreDeleteRequest struct {
    model.Params
    // 门店id
    _storeId   int64
}

// 初始化TaobaoPlaceStoreDeleteRequest对象
func NewTaobaoPlaceStoreDeleteRequest() *TaobaoPlaceStoreDeleteRequest{
    return &TaobaoPlaceStoreDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreDeleteRequest) GetApiMethodName() string {
    return "taobao.place.store.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店id
func (r *TaobaoPlaceStoreDeleteRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoPlaceStoreDeleteRequest) GetStoreId() int64 {
    return r._storeId
}
