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
    storeId   int64
    // 是否是活动期
    activity   bool
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
func (r *TaobaoOmniorderStoreCollectconfigGetRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r TaobaoOmniorderStoreCollectconfigGetRequest) GetStoreId() int64 {
    return r.storeId
}
// Activity Setter
// 是否是活动期
func (r *TaobaoOmniorderStoreCollectconfigGetRequest) SetActivity(activity bool) error {
    r.activity = activity
    r.Set("activity", activity)
    return nil
}

// Activity Getter
func (r TaobaoOmniorderStoreCollectconfigGetRequest) GetActivity() bool {
    return r.activity
}
