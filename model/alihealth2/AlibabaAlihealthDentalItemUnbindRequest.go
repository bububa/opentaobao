package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV解绑商品 API请求
alibaba.alihealth.dental.item.unbind

ISV解绑商品
*/
type AlibabaAlihealthDentalItemUnbindRequest struct {
    model.Params
    // ISV门店ID
    _spStoreId   string
    // ISV商品ID
    _spItemId   string
}

// 初始化AlibabaAlihealthDentalItemUnbindRequest对象
func NewAlibabaAlihealthDentalItemUnbindRequest() *AlibabaAlihealthDentalItemUnbindRequest{
    return &AlibabaAlihealthDentalItemUnbindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDentalItemUnbindRequest) GetApiMethodName() string {
    return "alibaba.alihealth.dental.item.unbind"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDentalItemUnbindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SpStoreId Setter
// ISV门店ID
func (r *AlibabaAlihealthDentalItemUnbindRequest) SetSpStoreId(_spStoreId string) error {
    r._spStoreId = _spStoreId
    r.Set("sp_store_id", _spStoreId)
    return nil
}

// SpStoreId Getter
func (r AlibabaAlihealthDentalItemUnbindRequest) GetSpStoreId() string {
    return r._spStoreId
}
// SpItemId Setter
// ISV商品ID
func (r *AlibabaAlihealthDentalItemUnbindRequest) SetSpItemId(_spItemId string) error {
    r._spItemId = _spItemId
    r.Set("sp_item_id", _spItemId)
    return nil
}

// SpItemId Getter
func (r AlibabaAlihealthDentalItemUnbindRequest) GetSpItemId() string {
    return r._spItemId
}
