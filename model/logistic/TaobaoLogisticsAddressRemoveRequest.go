package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除卖家地址库 API请求
taobao.logistics.address.remove

用此接口删除卖家地址库
*/
type TaobaoLogisticsAddressRemoveRequest struct {
    model.Params
    // 地址库ID
    _contactId   int64
}

// 初始化TaobaoLogisticsAddressRemoveRequest对象
func NewTaobaoLogisticsAddressRemoveRequest() *TaobaoLogisticsAddressRemoveRequest{
    return &TaobaoLogisticsAddressRemoveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsAddressRemoveRequest) GetApiMethodName() string {
    return "taobao.logistics.address.remove"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsAddressRemoveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ContactId Setter
// 地址库ID
func (r *TaobaoLogisticsAddressRemoveRequest) SetContactId(_contactId int64) error {
    r._contactId = _contactId
    r.Set("contact_id", _contactId)
    return nil
}

// ContactId Getter
func (r TaobaoLogisticsAddressRemoveRequest) GetContactId() int64 {
    return r._contactId
}
