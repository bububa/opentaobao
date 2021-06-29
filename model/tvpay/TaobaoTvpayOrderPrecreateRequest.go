package tvpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付预下单 API请求
taobao.tvpay.order.precreate

tv支付预下单
*/
type TaobaoTvpayOrderPrecreateRequest struct {
    model.Params
    // 设备id
    deviceId   string
    // 来源
    from   string
    // 订单详情
    data   string
    // 牌照方
    license   string
}

// 初始化TaobaoTvpayOrderPrecreateRequest对象
func NewTaobaoTvpayOrderPrecreateRequest() *TaobaoTvpayOrderPrecreateRequest{
    return &TaobaoTvpayOrderPrecreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTvpayOrderPrecreateRequest) GetApiMethodName() string {
    return "taobao.tvpay.order.precreate"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTvpayOrderPrecreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceId Setter
// 设备id
func (r *TaobaoTvpayOrderPrecreateRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

// DeviceId Getter
func (r TaobaoTvpayOrderPrecreateRequest) GetDeviceId() string {
    return r.deviceId
}
// From Setter
// 来源
func (r *TaobaoTvpayOrderPrecreateRequest) SetFrom(from string) error {
    r.from = from
    r.Set("from", from)
    return nil
}

// From Getter
func (r TaobaoTvpayOrderPrecreateRequest) GetFrom() string {
    return r.from
}
// Data Setter
// 订单详情
func (r *TaobaoTvpayOrderPrecreateRequest) SetData(data string) error {
    r.data = data
    r.Set("data", data)
    return nil
}

// Data Getter
func (r TaobaoTvpayOrderPrecreateRequest) GetData() string {
    return r.data
}
// License Setter
// 牌照方
func (r *TaobaoTvpayOrderPrecreateRequest) SetLicense(license string) error {
    r.license = license
    r.Set("license", license)
    return nil
}

// License Getter
func (r TaobaoTvpayOrderPrecreateRequest) GetLicense() string {
    return r.license
}
