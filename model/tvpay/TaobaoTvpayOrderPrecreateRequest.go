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
type TaobaoTvpayOrderPrecreateAPIRequest struct {
    model.Params
    // 设备id
    _deviceId   string
    // 来源
    _from   string
    // 订单详情
    _data   string
    // 牌照方
    _license   string
}

// 初始化TaobaoTvpayOrderPrecreateAPIRequest对象
func NewTaobaoTvpayOrderPrecreateRequest() *TaobaoTvpayOrderPrecreateAPIRequest{
    return &TaobaoTvpayOrderPrecreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTvpayOrderPrecreateAPIRequest) GetApiMethodName() string {
    return "taobao.tvpay.order.precreate"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTvpayOrderPrecreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceId Setter
// 设备id
func (r *TaobaoTvpayOrderPrecreateAPIRequest) SetDeviceId(_deviceId string) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r TaobaoTvpayOrderPrecreateAPIRequest) GetDeviceId() string {
    return r._deviceId
}
// From Setter
// 来源
func (r *TaobaoTvpayOrderPrecreateAPIRequest) SetFrom(_from string) error {
    r._from = _from
    r.Set("from", _from)
    return nil
}

// From Getter
func (r TaobaoTvpayOrderPrecreateAPIRequest) GetFrom() string {
    return r._from
}
// Data Setter
// 订单详情
func (r *TaobaoTvpayOrderPrecreateAPIRequest) SetData(_data string) error {
    r._data = _data
    r.Set("data", _data)
    return nil
}

// Data Getter
func (r TaobaoTvpayOrderPrecreateAPIRequest) GetData() string {
    return r._data
}
// License Setter
// 牌照方
func (r *TaobaoTvpayOrderPrecreateAPIRequest) SetLicense(_license string) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r TaobaoTvpayOrderPrecreateAPIRequest) GetLicense() string {
    return r._license
}
