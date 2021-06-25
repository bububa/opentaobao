package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店信息查询接口 APIRequest
taobao.qimen.store.query

商家在ERP等系统中调用该接口，查询门店相关信息
*/
type TaobaoQimenStoreQueryRequest struct {
    model.Params

    // 已分配的线上门店ID
    storeId   int64 

}

func NewTaobaoQimenStoreQueryRequest() *TaobaoQimenStoreQueryRequest{
    return &TaobaoQimenStoreQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenStoreQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.store.query"
}

func (r TaobaoQimenStoreQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenStoreQueryRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoQimenStoreQueryRequest) GetStoreId() int64 {
    return r.storeId
}

