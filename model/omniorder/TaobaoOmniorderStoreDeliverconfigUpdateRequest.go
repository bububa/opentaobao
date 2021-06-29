package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改门店发货配置内容 API请求
taobao.omniorder.store.deliverconfig.update

修改门店发货配置内容
*/
type TaobaoOmniorderStoreDeliverconfigUpdateRequest struct {
    model.Params
    // 门店ID
    _storeId   int64
    // 卖家发货配置
    _storeDeliverConfig   *StoreDeliverConfig
}

// 初始化TaobaoOmniorderStoreDeliverconfigUpdateRequest对象
func NewTaobaoOmniorderStoreDeliverconfigUpdateRequest() *TaobaoOmniorderStoreDeliverconfigUpdateRequest{
    return &TaobaoOmniorderStoreDeliverconfigUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreDeliverconfigUpdateRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.deliverconfig.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreDeliverconfigUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店ID
func (r *TaobaoOmniorderStoreDeliverconfigUpdateRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoOmniorderStoreDeliverconfigUpdateRequest) GetStoreId() int64 {
    return r._storeId
}
// StoreDeliverConfig Setter
// 卖家发货配置
func (r *TaobaoOmniorderStoreDeliverconfigUpdateRequest) SetStoreDeliverConfig(_storeDeliverConfig *StoreDeliverConfig) error {
    r._storeDeliverConfig = _storeDeliverConfig
    r.Set("store_deliver_config", _storeDeliverConfig)
    return nil
}

// StoreDeliverConfig Getter
func (r TaobaoOmniorderStoreDeliverconfigUpdateRequest) GetStoreDeliverConfig() *StoreDeliverConfig {
    return r._storeDeliverConfig
}
