package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
解绑商品信息 API请求
alibaba.alihealth.reserve.dental.unbinditem

绑定门店信息，商品信息
*/
type AlibabaAlihealthReserveDentalUnbinditemRequest struct {
    model.Params
    // 服务商门店id
    _spStoreId   string
    // 服务商商品id
    _spItemId   string
}

// 初始化AlibabaAlihealthReserveDentalUnbinditemRequest对象
func NewAlibabaAlihealthReserveDentalUnbinditemRequest() *AlibabaAlihealthReserveDentalUnbinditemRequest{
    return &AlibabaAlihealthReserveDentalUnbinditemRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthReserveDentalUnbinditemRequest) GetApiMethodName() string {
    return "alibaba.alihealth.reserve.dental.unbinditem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthReserveDentalUnbinditemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SpStoreId Setter
// 服务商门店id
func (r *AlibabaAlihealthReserveDentalUnbinditemRequest) SetSpStoreId(_spStoreId string) error {
    r._spStoreId = _spStoreId
    r.Set("sp_store_id", _spStoreId)
    return nil
}

// SpStoreId Getter
func (r AlibabaAlihealthReserveDentalUnbinditemRequest) GetSpStoreId() string {
    return r._spStoreId
}
// SpItemId Setter
// 服务商商品id
func (r *AlibabaAlihealthReserveDentalUnbinditemRequest) SetSpItemId(_spItemId string) error {
    r._spItemId = _spItemId
    r.Set("sp_item_id", _spItemId)
    return nil
}

// SpItemId Getter
func (r AlibabaAlihealthReserveDentalUnbinditemRequest) GetSpItemId() string {
    return r._spItemId
}
