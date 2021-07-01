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
type TaobaoTvpayAppinfoGetAPIRequest struct {
    model.Params
    // 设备id
    _deviceId   string
    // 来源
    _from   string
    // 客户端版本号
    _clientVersion   string
}

// 初始化TaobaoTvpayAppinfoGetAPIRequest对象
func NewTaobaoTvpayAppinfoGetRequest() *TaobaoTvpayAppinfoGetAPIRequest{
    return &TaobaoTvpayAppinfoGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTvpayAppinfoGetAPIRequest) GetApiMethodName() string {
    return "taobao.tvpay.appinfo.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTvpayAppinfoGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceId Setter
// 设备id
func (r *TaobaoTvpayAppinfoGetAPIRequest) SetDeviceId(_deviceId string) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r TaobaoTvpayAppinfoGetAPIRequest) GetDeviceId() string {
    return r._deviceId
}
// From Setter
// 来源
func (r *TaobaoTvpayAppinfoGetAPIRequest) SetFrom(_from string) error {
    r._from = _from
    r.Set("from", _from)
    return nil
}

// From Getter
func (r TaobaoTvpayAppinfoGetAPIRequest) GetFrom() string {
    return r._from
}
// ClientVersion Setter
// 客户端版本号
func (r *TaobaoTvpayAppinfoGetAPIRequest) SetClientVersion(_clientVersion string) error {
    r._clientVersion = _clientVersion
    r.Set("client_version", _clientVersion)
    return nil
}

// ClientVersion Getter
func (r TaobaoTvpayAppinfoGetAPIRequest) GetClientVersion() string {
    return r._clientVersion
}
