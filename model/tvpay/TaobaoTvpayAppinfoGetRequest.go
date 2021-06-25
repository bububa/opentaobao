package tvpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付获取应用信息 APIRequest
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

func NewTaobaoTvpayAppinfoGetRequest() *TaobaoTvpayAppinfoGetRequest{
    return &TaobaoTvpayAppinfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTvpayAppinfoGetRequest) GetApiMethodName() string {
    return "taobao.tvpay.appinfo.get"
}

func (r TaobaoTvpayAppinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTvpayAppinfoGetRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

func (r TaobaoTvpayAppinfoGetRequest) GetDeviceId() string {
    return r.deviceId
}

func (r *TaobaoTvpayAppinfoGetRequest) SetFrom(from string) error {
    r.from = from
    r.Set("from", from)
    return nil
}

func (r TaobaoTvpayAppinfoGetRequest) GetFrom() string {
    return r.from
}

func (r *TaobaoTvpayAppinfoGetRequest) SetClientVersion(clientVersion string) error {
    r.clientVersion = clientVersion
    r.Set("client_version", clientVersion)
    return nil
}

func (r TaobaoTvpayAppinfoGetRequest) GetClientVersion() string {
    return r.clientVersion
}

