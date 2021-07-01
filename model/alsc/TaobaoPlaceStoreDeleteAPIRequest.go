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
type TaobaoPlaceStoreDeleteAPIRequest struct {
    model.Params
    // 门店id
    _storeId   int64
}

// 初始化TaobaoPlaceStoreDeleteAPIRequest对象
func NewTaobaoPlaceStoreDeleteRequest() *TaobaoPlaceStoreDeleteAPIRequest{
    return &TaobaoPlaceStoreDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreDeleteAPIRequest) GetApiMethodName() string {
    return "taobao.place.store.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店id
func (r *TaobaoPlaceStoreDeleteAPIRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoPlaceStoreDeleteAPIRequest) GetStoreId() int64 {
    return r._storeId
}
