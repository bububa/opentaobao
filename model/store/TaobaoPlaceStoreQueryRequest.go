package store

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店信息查询接口 APIRequest
taobao.place.store.query

根据用户授权信息，获取用户的门店公开信息
*/
type TaobaoPlaceStoreQueryRequest struct {
    model.Params

    // 业务code，用于区分业务
    bizCode   string 

    // 业务外部id
    outerId   string 

    // 门店id
    storeId   int64 

}

func NewTaobaoPlaceStoreQueryRequest() *TaobaoPlaceStoreQueryRequest{
    return &TaobaoPlaceStoreQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPlaceStoreQueryRequest) GetApiMethodName() string {
    return "taobao.place.store.query"
}

func (r TaobaoPlaceStoreQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPlaceStoreQueryRequest) SetBizCode(bizCode string) error {
    r.bizCode = bizCode
    r.Set("biz_code", bizCode)
    return nil
}

func (r TaobaoPlaceStoreQueryRequest) GetBizCode() string {
    return r.bizCode
}

func (r *TaobaoPlaceStoreQueryRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoPlaceStoreQueryRequest) GetOuterId() string {
    return r.outerId
}

func (r *TaobaoPlaceStoreQueryRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoPlaceStoreQueryRequest) GetStoreId() int64 {
    return r.storeId
}

