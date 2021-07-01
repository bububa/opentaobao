package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商户门店创建接口 API请求
taobao.place.store.create

用于商家创建线下门店
*/
type TaobaoPlaceStoreCreateAPIRequest struct {
    model.Params
    // 门店创建入参
    _storeCreate   *StoreUpdateTopDTO
}

// 初始化TaobaoPlaceStoreCreateAPIRequest对象
func NewTaobaoPlaceStoreCreateRequest() *TaobaoPlaceStoreCreateAPIRequest{
    return &TaobaoPlaceStoreCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreCreateAPIRequest) GetApiMethodName() string {
    return "taobao.place.store.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreCreate Setter
// 门店创建入参
func (r *TaobaoPlaceStoreCreateAPIRequest) SetStoreCreate(_storeCreate *StoreUpdateTopDTO) error {
    r._storeCreate = _storeCreate
    r.Set("store_create", _storeCreate)
    return nil
}

// StoreCreate Getter
func (r TaobaoPlaceStoreCreateAPIRequest) GetStoreCreate() *StoreUpdateTopDTO {
    return r._storeCreate
}
