package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询门店发货配置内容 APIRequest
taobao.omniorder.store.deliverconfig.get

查询门店发货配置内容
*/
type TaobaoOmniorderStoreDeliverconfigGetRequest struct {
    model.Params

    // 门店ID
    storeId   int64 

    // 是否是活动期
    activity   bool 

}

func NewTaobaoOmniorderStoreDeliverconfigGetRequest() *TaobaoOmniorderStoreDeliverconfigGetRequest{
    return &TaobaoOmniorderStoreDeliverconfigGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniorderStoreDeliverconfigGetRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.deliverconfig.get"
}

func (r TaobaoOmniorderStoreDeliverconfigGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniorderStoreDeliverconfigGetRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoOmniorderStoreDeliverconfigGetRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoOmniorderStoreDeliverconfigGetRequest) SetActivity(activity bool) error {
    r.activity = activity
    r.Set("activity", activity)
    return nil
}

func (r TaobaoOmniorderStoreDeliverconfigGetRequest) GetActivity() bool {
    return r.activity
}

