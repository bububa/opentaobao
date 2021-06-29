package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店自提配置修改 APIRequest
taobao.omniorder.store.collectconfig.update

修改门店自提配置内容
*/
type TaobaoOmniorderStoreCollectconfigUpdateRequest struct {
    model.Params

    // 门店自提配置
    storeCollectConfig   *StoreCollectConfig 

    // 门店ID
    storeId   int64 

}

func NewTaobaoOmniorderStoreCollectconfigUpdateRequest() *TaobaoOmniorderStoreCollectconfigUpdateRequest{
    return &TaobaoOmniorderStoreCollectconfigUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniorderStoreCollectconfigUpdateRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.collectconfig.update"
}

func (r TaobaoOmniorderStoreCollectconfigUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniorderStoreCollectconfigUpdateRequest) SetStoreCollectConfig(storeCollectConfig *StoreCollectConfig) error {
    r.storeCollectConfig = storeCollectConfig
    r.Set("store_collect_config", storeCollectConfig)
    return nil
}

func (r TaobaoOmniorderStoreCollectconfigUpdateRequest) GetStoreCollectConfig() *StoreCollectConfig {
    return r.storeCollectConfig
}

func (r *TaobaoOmniorderStoreCollectconfigUpdateRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoOmniorderStoreCollectconfigUpdateRequest) GetStoreId() int64 {
    return r.storeId
}

