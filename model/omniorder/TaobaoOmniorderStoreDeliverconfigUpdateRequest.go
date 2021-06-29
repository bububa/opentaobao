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
    storeId   int64
    // 卖家发货配置
    storeDeliverConfig   *StoreDeliverConfig
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
func (r *TaobaoOmniorderStoreDeliverconfigUpdateRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r TaobaoOmniorderStoreDeliverconfigUpdateRequest) GetStoreId() int64 {
    return r.storeId
}
// StoreDeliverConfig Setter
// 卖家发货配置
func (r *TaobaoOmniorderStoreDeliverconfigUpdateRequest) SetStoreDeliverConfig(storeDeliverConfig *StoreDeliverConfig) error {
    r.storeDeliverConfig = storeDeliverConfig
    r.Set("store_deliver_config", storeDeliverConfig)
    return nil
}

// StoreDeliverConfig Getter
func (r TaobaoOmniorderStoreDeliverconfigUpdateRequest) GetStoreDeliverConfig() *StoreDeliverConfig {
    return r.storeDeliverConfig
}
