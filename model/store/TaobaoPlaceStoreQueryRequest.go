package store

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店信息查询接口 API请求
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

// 初始化TaobaoPlaceStoreQueryRequest对象
func NewTaobaoPlaceStoreQueryRequest() *TaobaoPlaceStoreQueryRequest{
    return &TaobaoPlaceStoreQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreQueryRequest) GetApiMethodName() string {
    return "taobao.place.store.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizCode Setter
// 业务code，用于区分业务
func (r *TaobaoPlaceStoreQueryRequest) SetBizCode(bizCode string) error {
    r.bizCode = bizCode
    r.Set("biz_code", bizCode)
    return nil
}

// BizCode Getter
func (r TaobaoPlaceStoreQueryRequest) GetBizCode() string {
    return r.bizCode
}
// OuterId Setter
// 业务外部id
func (r *TaobaoPlaceStoreQueryRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoPlaceStoreQueryRequest) GetOuterId() string {
    return r.outerId
}
// StoreId Setter
// 门店id
func (r *TaobaoPlaceStoreQueryRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r TaobaoPlaceStoreQueryRequest) GetStoreId() int64 {
    return r.storeId
}
