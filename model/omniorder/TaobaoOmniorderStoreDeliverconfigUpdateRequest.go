package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改门店发货配置内容 APIRequest
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

func NewTaobaoOmniorderStoreDeliverconfigUpdateRequest() *TaobaoOmniorderStoreDeliverconfigUpdateRequest{
    return &TaobaoOmniorderStoreDeliverconfigUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniorderStoreDeliverconfigUpdateRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.deliverconfig.update"
}

func (r TaobaoOmniorderStoreDeliverconfigUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniorderStoreDeliverconfigUpdateRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoOmniorderStoreDeliverconfigUpdateRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoOmniorderStoreDeliverconfigUpdateRequest) SetStoreDeliverConfig(storeDeliverConfig *StoreDeliverConfig) error {
    r.storeDeliverConfig = storeDeliverConfig
    r.Set("store_deliver_config", storeDeliverConfig)
    return nil
}

func (r TaobaoOmniorderStoreDeliverconfigUpdateRequest) GetStoreDeliverConfig() *StoreDeliverConfig {
    return r.storeDeliverConfig
}

