package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询门店自提配置内容 API请求
taobao.omniorder.store.collectconfig.get

查询门店自提配置内容
*/
type TaobaoOmniorderStoreCollectconfigGetRequest struct {
    model.Params
    // 门店ID
    _storeId   int64
    // 是否是活动期
    _activity   bool
}

// 初始化TaobaoOmniorderStoreCollectconfigGetRequest对象
func NewTaobaoOmniorderStoreCollectconfigGetRequest() *TaobaoOmniorderStoreCollectconfigGetRequest{
    return &TaobaoOmniorderStoreCollectconfigGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreCollectconfigGetRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.collectconfig.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreCollectconfigGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店ID
func (r *TaobaoOmniorderStoreCollectconfigGetRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoOmniorderStoreCollectconfigGetRequest) GetStoreId() int64 {
    return r._storeId
}
// Activity Setter
// 是否是活动期
func (r *TaobaoOmniorderStoreCollectconfigGetRequest) SetActivity(_activity bool) error {
    r._activity = _activity
    r.Set("activity", _activity)
    return nil
}

// Activity Getter
func (r TaobaoOmniorderStoreCollectconfigGetRequest) GetActivity() bool {
    return r._activity
}
