package tvpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付第三方支付订单 API请求
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

// 初始化TaobaoTvpayOrderPartnerpayRequest对象
func NewTaobaoTvpayOrderPartnerpayRequest() *TaobaoTvpayOrderPartnerpayRequest{
    return &TaobaoTvpayOrderPartnerpayRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTvpayOrderPartnerpayRequest) GetApiMethodName() string {
    return "taobao.tvpay.order.partnerpay"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTvpayOrderPartnerpayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceId Setter
// 设备id
func (r *TaobaoTvpayOrderPartnerpayRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

// DeviceId Getter
func (r TaobaoTvpayOrderPartnerpayRequest) GetDeviceId() string {
    return r.deviceId
}
// From Setter
// 来源
func (r *TaobaoTvpayOrderPartnerpayRequest) SetFrom(from string) error {
    r.from = from
    r.Set("from", from)
    return nil
}

// From Getter
func (r TaobaoTvpayOrderPartnerpayRequest) GetFrom() string {
    return r.from
}
// Data Setter
// 订单信息
func (r *TaobaoTvpayOrderPartnerpayRequest) SetData(data string) error {
    r.data = data
    r.Set("data", data)
    return nil
}

// Data Getter
func (r TaobaoTvpayOrderPartnerpayRequest) GetData() string {
    return r.data
}
// PayType Setter
// 支付方式
func (r *TaobaoTvpayOrderPartnerpayRequest) SetPayType(payType string) error {
    r.payType = payType
    r.Set("pay_type", payType)
    return nil
}

// PayType Getter
func (r TaobaoTvpayOrderPartnerpayRequest) GetPayType() string {
    return r.payType
}
// License Setter
// 牌照方
func (r *TaobaoTvpayOrderPartnerpayRequest) SetLicense(license string) error {
    r.license = license
    r.Set("license", license)
    return nil
}

// License Getter
func (r TaobaoTvpayOrderPartnerpayRequest) GetLicense() string {
    return r.license
}
