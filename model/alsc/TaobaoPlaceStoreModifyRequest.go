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
type TaobaoPlaceStoreModifyRequest struct {
    model.Params
    // 门店创建入参
    _storeUpdate   *StoreUpdateTopDTO
}

// 初始化TaobaoPlaceStoreModifyRequest对象
func NewTaobaoPlaceStoreModifyRequest() *TaobaoPlaceStoreModifyRequest{
    return &TaobaoPlaceStoreModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreModifyRequest) GetApiMethodName() string {
    return "taobao.place.store.modify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreUpdate Setter
// 门店创建入参
func (r *TaobaoPlaceStoreModifyRequest) SetStoreUpdate(_storeUpdate *StoreUpdateTopDTO) error {
    r._storeUpdate = _storeUpdate
    r.Set("store_update", _storeUpdate)
    return nil
}

// StoreUpdate Getter
func (r TaobaoPlaceStoreModifyRequest) GetStoreUpdate() *StoreUpdateTopDTO {
    return r._storeUpdate
}
