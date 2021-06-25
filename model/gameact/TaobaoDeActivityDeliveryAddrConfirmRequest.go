package gameact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户收件地址确认 APIRequest
taobao.de.activity.delivery.addr.confirm

用户收件地址确认
*/
type TaobaoDeActivityDeliveryAddrConfirmRequest struct {
    model.Params

    // 加密流水号
    serialNumber   string 

    // 地址Sign
    addressSign   string 

}

func NewTaobaoDeActivityDeliveryAddrConfirmRequest() *TaobaoDeActivityDeliveryAddrConfirmRequest{
    return &TaobaoDeActivityDeliveryAddrConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoDeActivityDeliveryAddrConfirmRequest) GetApiMethodName() string {
    return "taobao.de.activity.delivery.addr.confirm"
}

func (r TaobaoDeActivityDeliveryAddrConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoDeActivityDeliveryAddrConfirmRequest) SetSerialNumber(serialNumber string) error {
    r.serialNumber = serialNumber
    r.Set("serial_number", serialNumber)
    return nil
}

func (r TaobaoDeActivityDeliveryAddrConfirmRequest) GetSerialNumber() string {
    return r.serialNumber
}

func (r *TaobaoDeActivityDeliveryAddrConfirmRequest) SetAddressSign(addressSign string) error {
    r.addressSign = addressSign
    r.Set("address_sign", addressSign)
    return nil
}

func (r TaobaoDeActivityDeliveryAddrConfirmRequest) GetAddressSign() string {
    return r.addressSign
}

