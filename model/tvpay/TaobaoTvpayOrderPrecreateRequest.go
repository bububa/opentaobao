package tvpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付预下单 APIRequest
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

func NewTaobaoTvpayOrderPrecreateRequest() *TaobaoTvpayOrderPrecreateRequest{
    return &TaobaoTvpayOrderPrecreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTvpayOrderPrecreateRequest) GetApiMethodName() string {
    return "taobao.tvpay.order.precreate"
}

func (r TaobaoTvpayOrderPrecreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTvpayOrderPrecreateRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

func (r TaobaoTvpayOrderPrecreateRequest) GetDeviceId() string {
    return r.deviceId
}

func (r *TaobaoTvpayOrderPrecreateRequest) SetFrom(from string) error {
    r.from = from
    r.Set("from", from)
    return nil
}

func (r TaobaoTvpayOrderPrecreateRequest) GetFrom() string {
    return r.from
}

func (r *TaobaoTvpayOrderPrecreateRequest) SetData(data string) error {
    r.data = data
    r.Set("data", data)
    return nil
}

func (r TaobaoTvpayOrderPrecreateRequest) GetData() string {
    return r.data
}

func (r *TaobaoTvpayOrderPrecreateRequest) SetLicense(license string) error {
    r.license = license
    r.Set("license", license)
    return nil
}

func (r TaobaoTvpayOrderPrecreateRequest) GetLicense() string {
    return r.license
}

