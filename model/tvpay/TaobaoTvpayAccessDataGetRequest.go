package tvpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付 API请求
taobao.tvpay.access.data.get

在匿名用户支付后尝试为其登陆绑定的淘宝账号
*/
type TaobaoTvpayAccessDataGetRequest struct {
    model.Params
    // 设备id
    deviceId   string
    // 来源
    from   string
    // 订单id
    outOrderNo   string
    // 账号客户端版本
    accountClientVersion   string
}

// 初始化TaobaoTvpayAccessDataGetRequest对象
func NewTaobaoTvpayAccessDataGetRequest() *TaobaoTvpayAccessDataGetRequest{
    return &TaobaoTvpayAccessDataGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTvpayAccessDataGetRequest) GetApiMethodName() string {
    return "taobao.tvpay.access.data.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTvpayAccessDataGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceId Setter
// 设备id
func (r *TaobaoTvpayAccessDataGetRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

// DeviceId Getter
func (r TaobaoTvpayAccessDataGetRequest) GetDeviceId() string {
    return r.deviceId
}
// From Setter
// 来源
func (r *TaobaoTvpayAccessDataGetRequest) SetFrom(from string) error {
    r.from = from
    r.Set("from", from)
    return nil
}

// From Getter
func (r TaobaoTvpayAccessDataGetRequest) GetFrom() string {
    return r.from
}
// OutOrderNo Setter
// 订单id
func (r *TaobaoTvpayAccessDataGetRequest) SetOutOrderNo(outOrderNo string) error {
    r.outOrderNo = outOrderNo
    r.Set("out_order_no", outOrderNo)
    return nil
}

// OutOrderNo Getter
func (r TaobaoTvpayAccessDataGetRequest) GetOutOrderNo() string {
    return r.outOrderNo
}
// AccountClientVersion Setter
// 账号客户端版本
func (r *TaobaoTvpayAccessDataGetRequest) SetAccountClientVersion(accountClientVersion string) error {
    r.accountClientVersion = accountClientVersion
    r.Set("account_client_version", accountClientVersion)
    return nil
}

// AccountClientVersion Getter
func (r TaobaoTvpayAccessDataGetRequest) GetAccountClientVersion() string {
    return r.accountClientVersion
}
