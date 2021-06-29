package tvpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付获取应用信息 API请求
taobao.tvpay.appinfo.get

tv支付获取应用信息
*/
type TaobaoTvpayAppinfoGetRequest struct {
    model.Params
    // 设备id
    deviceId   string
    // 来源
    from   string
    // 客户端版本号
    clientVersion   string
}

// 初始化TaobaoTvpayAppinfoGetRequest对象
func NewTaobaoTvpayAppinfoGetRequest() *TaobaoTvpayAppinfoGetRequest{
    return &TaobaoTvpayAppinfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTvpayAppinfoGetRequest) GetApiMethodName() string {
    return "taobao.tvpay.appinfo.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTvpayAppinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceId Setter
// 设备id
func (r *TaobaoTvpayAppinfoGetRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

// DeviceId Getter
func (r TaobaoTvpayAppinfoGetRequest) GetDeviceId() string {
    return r.deviceId
}
// From Setter
// 来源
func (r *TaobaoTvpayAppinfoGetRequest) SetFrom(from string) error {
    r.from = from
    r.Set("from", from)
    return nil
}

// From Getter
func (r TaobaoTvpayAppinfoGetRequest) GetFrom() string {
    return r.from
}
// ClientVersion Setter
// 客户端版本号
func (r *TaobaoTvpayAppinfoGetRequest) SetClientVersion(clientVersion string) error {
    r.clientVersion = clientVersion
    r.Set("client_version", clientVersion)
    return nil
}

// ClientVersion Getter
func (r TaobaoTvpayAppinfoGetRequest) GetClientVersion() string {
    return r.clientVersion
}
