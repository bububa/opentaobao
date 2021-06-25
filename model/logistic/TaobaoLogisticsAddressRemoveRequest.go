package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除卖家地址库 APIRequest
taobao.logistics.address.remove

用此接口删除卖家地址库
*/
type TaobaoLogisticsAddressRemoveRequest struct {
    model.Params

    // 地址库ID
    contactId   int64 

}

func NewTaobaoLogisticsAddressRemoveRequest() *TaobaoLogisticsAddressRemoveRequest{
    return &TaobaoLogisticsAddressRemoveRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoLogisticsAddressRemoveRequest) GetApiMethodName() string {
    return "taobao.logistics.address.remove"
}

func (r TaobaoLogisticsAddressRemoveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoLogisticsAddressRemoveRequest) SetContactId(contactId int64) error {
    r.contactId = contactId
    r.Set("contact_id", contactId)
    return nil
}

func (r TaobaoLogisticsAddressRemoveRequest) GetContactId() int64 {
    return r.contactId
}

