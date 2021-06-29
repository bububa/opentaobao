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
    _deviceId   string
    // 来源
    _from   string
    // 订单信息
    _data   string
    // 支付方式
    _payType   string
    // 牌照方
    _license   string
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
func (r *TaobaoTvpayOrderPartnerpayRequest) SetDeviceId(_deviceId string) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r TaobaoTvpayOrderPartnerpayRequest) GetDeviceId() string {
    return r._deviceId
}
// From Setter
// 来源
func (r *TaobaoTvpayOrderPartnerpayRequest) SetFrom(_from string) error {
    r._from = _from
    r.Set("from", _from)
    return nil
}

// From Getter
func (r TaobaoTvpayOrderPartnerpayRequest) GetFrom() string {
    return r._from
}
// Data Setter
// 订单信息
func (r *TaobaoTvpayOrderPartnerpayRequest) SetData(_data string) error {
    r._data = _data
    r.Set("data", _data)
    return nil
}

// Data Getter
func (r TaobaoTvpayOrderPartnerpayRequest) GetData() string {
    return r._data
}
// PayType Setter
// 支付方式
func (r *TaobaoTvpayOrderPartnerpayRequest) SetPayType(_payType string) error {
    r._payType = _payType
    r.Set("pay_type", _payType)
    return nil
}

// PayType Getter
func (r TaobaoTvpayOrderPartnerpayRequest) GetPayType() string {
    return r._payType
}
// License Setter
// 牌照方
func (r *TaobaoTvpayOrderPartnerpayRequest) SetLicense(_license string) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r TaobaoTvpayOrderPartnerpayRequest) GetLicense() string {
    return r._license
}
