package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
自助机失效二维码 APIRequest
taobao.bus.disableqrcode.set

使创建的二维码失效
*/
type TaobaoBusDisableqrcodeSetRequest struct {
    model.Params

    // 飞猪订单号
    alitripOrderId   string 

}

func NewTaobaoBusDisableqrcodeSetRequest() *TaobaoBusDisableqrcodeSetRequest{
    return &TaobaoBusDisableqrcodeSetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusDisableqrcodeSetRequest) GetApiMethodName() string {
    return "taobao.bus.disableqrcode.set"
}

func (r TaobaoBusDisableqrcodeSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBusDisableqrcodeSetRequest) SetAlitripOrderId(alitripOrderId string) error {
    r.alitripOrderId = alitripOrderId
    r.Set("alitrip_order_id", alitripOrderId)
    return nil
}

func (r TaobaoBusDisableqrcodeSetRequest) GetAlitripOrderId() string {
    return r.alitripOrderId
}

