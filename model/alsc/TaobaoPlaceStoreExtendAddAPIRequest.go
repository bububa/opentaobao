package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增门店扩展属性 API请求
taobao.place.store.extend.add

新增授权用户的门店扩展属性
*/
type TaobaoPlaceStoreExtendAddAPIRequest struct {
    model.Params
    // 门店ID
    _storeId   int64
    // 扩展信息
    _etv   []ExtendTypeValueTopDto
}

// 初始化TaobaoPlaceStoreExtendAddAPIRequest对象
func NewTaobaoPlaceStoreExtendAddRequest() *TaobaoPlaceStoreExtendAddAPIRequest{
    return &TaobaoPlaceStoreExtendAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreExtendAddAPIRequest) GetApiMethodName() string {
    return "taobao.place.store.extend.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreExtendAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店ID
func (r *TaobaoPlaceStoreExtendAddAPIRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoPlaceStoreExtendAddAPIRequest) GetStoreId() int64 {
    return r._storeId
}
// Etv Setter
// 扩展信息
func (r *TaobaoPlaceStoreExtendAddAPIRequest) SetEtv(_etv []ExtendTypeValueTopDto) error {
    r._etv = _etv
    r.Set("etv", _etv)
    return nil
}

// Etv Getter
func (r TaobaoPlaceStoreExtendAddAPIRequest) GetEtv() []ExtendTypeValueTopDto {
    return r._etv
}
