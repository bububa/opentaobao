package tvpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付 APIRequest
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

func NewTaobaoTvpayAccessDataGetRequest() *TaobaoTvpayAccessDataGetRequest{
    return &TaobaoTvpayAccessDataGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTvpayAccessDataGetRequest) GetApiMethodName() string {
    return "taobao.tvpay.access.data.get"
}

func (r TaobaoTvpayAccessDataGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTvpayAccessDataGetRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

func (r TaobaoTvpayAccessDataGetRequest) GetDeviceId() string {
    return r.deviceId
}

func (r *TaobaoTvpayAccessDataGetRequest) SetFrom(from string) error {
    r.from = from
    r.Set("from", from)
    return nil
}

func (r TaobaoTvpayAccessDataGetRequest) GetFrom() string {
    return r.from
}

func (r *TaobaoTvpayAccessDataGetRequest) SetOutOrderNo(outOrderNo string) error {
    r.outOrderNo = outOrderNo
    r.Set("out_order_no", outOrderNo)
    return nil
}

func (r TaobaoTvpayAccessDataGetRequest) GetOutOrderNo() string {
    return r.outOrderNo
}

func (r *TaobaoTvpayAccessDataGetRequest) SetAccountClientVersion(accountClientVersion string) error {
    r.accountClientVersion = accountClientVersion
    r.Set("account_client_version", accountClientVersion)
    return nil
}

func (r TaobaoTvpayAccessDataGetRequest) GetAccountClientVersion() string {
    return r.accountClientVersion
}

