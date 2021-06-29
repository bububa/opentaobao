package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店自提配置修改 API请求
taobao.omniorder.store.collectconfig.update

修改门店自提配置内容
*/
type TaobaoOmniorderStoreCollectconfigUpdateRequest struct {
    model.Params
    // 门店自提配置
    _storeCollectConfig   *StoreCollectConfig
    // 门店ID
    _storeId   int64
}

// 初始化TaobaoOmniorderStoreCollectconfigUpdateRequest对象
func NewTaobaoOmniorderStoreCollectconfigUpdateRequest() *TaobaoOmniorderStoreCollectconfigUpdateRequest{
    return &TaobaoOmniorderStoreCollectconfigUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreCollectconfigUpdateRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.collectconfig.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreCollectconfigUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreCollectConfig Setter
// 门店自提配置
func (r *TaobaoOmniorderStoreCollectconfigUpdateRequest) SetStoreCollectConfig(_storeCollectConfig *StoreCollectConfig) error {
    r._storeCollectConfig = _storeCollectConfig
    r.Set("store_collect_config", _storeCollectConfig)
    return nil
}

// StoreCollectConfig Getter
func (r TaobaoOmniorderStoreCollectconfigUpdateRequest) GetStoreCollectConfig() *StoreCollectConfig {
    return r._storeCollectConfig
}
// StoreId Setter
// 门店ID
func (r *TaobaoOmniorderStoreCollectconfigUpdateRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoOmniorderStoreCollectconfigUpdateRequest) GetStoreId() int64 {
    return r._storeId
}
