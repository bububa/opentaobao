package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店删除接口 API请求
taobao.qimen.store.delete

商家在ERP等系统中调用该接口，删除线下门店
*/
type TaobaoQimenStoreDeleteRequest struct {
    model.Params
    // 要删除的门店id
    storeId   int64
}

// 初始化TaobaoQimenStoreDeleteRequest对象
func NewTaobaoQimenStoreDeleteRequest() *TaobaoQimenStoreDeleteRequest{
    return &TaobaoQimenStoreDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenStoreDeleteRequest) GetApiMethodName() string {
    return "taobao.qimen.store.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenStoreDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 要删除的门店id
func (r *TaobaoQimenStoreDeleteRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r TaobaoQimenStoreDeleteRequest) GetStoreId() int64 {
    return r.storeId
}
