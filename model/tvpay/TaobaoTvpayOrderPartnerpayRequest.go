package tvpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付第三方支付订单 APIRequest
taobao.tvpay.order.partnerpay

tv支付第三方发起并支付订单（使用设备授权）
*/
type TaobaoTvpayOrderPartnerpayRequest struct {
    model.Params

    // 设备id
    deviceId   string 

    // 来源
    from   string 

    // 订单信息
    data   string 

    // 支付方式
    payType   string 

    // 牌照方
    license   string 

}

func NewTaobaoTvpayOrderPartnerpayRequest() *TaobaoTvpayOrderPartnerpayRequest{
    return &TaobaoTvpayOrderPartnerpayRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTvpayOrderPartnerpayRequest) GetApiMethodName() string {
    return "taobao.tvpay.order.partnerpay"
}

func (r TaobaoTvpayOrderPartnerpayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTvpayOrderPartnerpayRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

func (r TaobaoTvpayOrderPartnerpayRequest) GetDeviceId() string {
    return r.deviceId
}

func (r *TaobaoTvpayOrderPartnerpayRequest) SetFrom(from string) error {
    r.from = from
    r.Set("from", from)
    return nil
}

func (r TaobaoTvpayOrderPartnerpayRequest) GetFrom() string {
    return r.from
}

func (r *TaobaoTvpayOrderPartnerpayRequest) SetData(data string) error {
    r.data = data
    r.Set("data", data)
    return nil
}

func (r TaobaoTvpayOrderPartnerpayRequest) GetData() string {
    return r.data
}

func (r *TaobaoTvpayOrderPartnerpayRequest) SetPayType(payType string) error {
    r.payType = payType
    r.Set("pay_type", payType)
    return nil
}

func (r TaobaoTvpayOrderPartnerpayRequest) GetPayType() string {
    return r.payType
}

func (r *TaobaoTvpayOrderPartnerpayRequest) SetLicense(license string) error {
    r.license = license
    r.Set("license", license)
    return nil
}

func (r TaobaoTvpayOrderPartnerpayRequest) GetLicense() string {
    return r.license
}

