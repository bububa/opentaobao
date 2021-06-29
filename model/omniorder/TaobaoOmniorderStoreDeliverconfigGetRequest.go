package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询门店发货配置内容 API请求
taobao.omniorder.store.deliverconfig.get

查询门店发货配置内容
*/
type TaobaoOmniorderStoreDeliverconfigGetRequest struct {
    model.Params
    // 门店ID
    _storeId   int64
    // 是否是活动期
    _activity   bool
}

// 初始化TaobaoOmniorderStoreDeliverconfigGetRequest对象
func NewTaobaoOmniorderStoreDeliverconfigGetRequest() *TaobaoOmniorderStoreDeliverconfigGetRequest{
    return &TaobaoOmniorderStoreDeliverconfigGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreDeliverconfigGetRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.deliverconfig.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreDeliverconfigGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店ID
func (r *TaobaoOmniorderStoreDeliverconfigGetRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoOmniorderStoreDeliverconfigGetRequest) GetStoreId() int64 {
    return r._storeId
}
// Activity Setter
// 是否是活动期
func (r *TaobaoOmniorderStoreDeliverconfigGetRequest) SetActivity(_activity bool) error {
    r._activity = _activity
    r.Set("activity", _activity)
    return nil
}

// Activity Getter
func (r TaobaoOmniorderStoreDeliverconfigGetRequest) GetActivity() bool {
    return r._activity
}
