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
    _bizCode   string
    // 业务外部id
    _outerId   string
    // 门店id
    _storeId   int64
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
func (r *TaobaoPlaceStoreQueryRequest) SetBizCode(_bizCode string) error {
    r._bizCode = _bizCode
    r.Set("biz_code", _bizCode)
    return nil
}

// BizCode Getter
func (r TaobaoPlaceStoreQueryRequest) GetBizCode() string {
    return r._bizCode
}
// OuterId Setter
// 业务外部id
func (r *TaobaoPlaceStoreQueryRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TaobaoPlaceStoreQueryRequest) GetOuterId() string {
    return r._outerId
}
// StoreId Setter
// 门店id
func (r *TaobaoPlaceStoreQueryRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoPlaceStoreQueryRequest) GetStoreId() int64 {
    return r._storeId
}
