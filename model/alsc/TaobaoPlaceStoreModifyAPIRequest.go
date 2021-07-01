package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家修改线下门店 API请求
taobao.place.store.modify

用于商家修改线下门店信息
*/
type TaobaoPlaceStoreModifyAPIRequest struct {
    model.Params
    // 门店创建入参
    _storeUpdate   *StoreUpdateTopDto
}

// 初始化TaobaoPlaceStoreModifyAPIRequest对象
func NewTaobaoPlaceStoreModifyRequest() *TaobaoPlaceStoreModifyAPIRequest{
    return &TaobaoPlaceStoreModifyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreModifyAPIRequest) GetApiMethodName() string {
    return "taobao.place.store.modify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreModifyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreUpdate Setter
// 门店创建入参
func (r *TaobaoPlaceStoreModifyAPIRequest) SetStoreUpdate(_storeUpdate *StoreUpdateTopDto) error {
    r._storeUpdate = _storeUpdate
    r.Set("store_update", _storeUpdate)
    return nil
}

// StoreUpdate Getter
func (r TaobaoPlaceStoreModifyAPIRequest) GetStoreUpdate() *StoreUpdateTopDto {
    return r._storeUpdate
}
