package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店信息查询接口 API请求
taobao.qimen.store.query

商家在ERP等系统中调用该接口，查询门店相关信息
*/
type TaobaoQimenStoreQueryRequest struct {
    model.Params
    // 已分配的线上门店ID
    storeId   int64
}

// 初始化TaobaoQimenStoreQueryRequest对象
func NewTaobaoQimenStoreQueryRequest() *TaobaoQimenStoreQueryRequest{
    return &TaobaoQimenStoreQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenStoreQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.store.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenStoreQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 已分配的线上门店ID
func (r *TaobaoQimenStoreQueryRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r TaobaoQimenStoreQueryRequest) GetStoreId() int64 {
    return r.storeId
}
