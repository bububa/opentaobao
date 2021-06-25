package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
自助机条形码被动支付 APIRequest
taobao.bus.tvmpayorder.set

汽车票线下自助机条形码支付
*/
type TaobaoBusTvmpayorderSetRequest struct {
    model.Params

    // 条形码认证码
    alipayAuthCode   string 

    // 飞猪订单号
    alitripOrderId   string 

}

func NewTaobaoBusTvmpayorderSetRequest() *TaobaoBusTvmpayorderSetRequest{
    return &TaobaoBusTvmpayorderSetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusTvmpayorderSetRequest) GetApiMethodName() string {
    return "taobao.bus.tvmpayorder.set"
}

func (r TaobaoBusTvmpayorderSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBusTvmpayorderSetRequest) SetAlipayAuthCode(alipayAuthCode string) error {
    r.alipayAuthCode = alipayAuthCode
    r.Set("alipay_auth_code", alipayAuthCode)
    return nil
}

func (r TaobaoBusTvmpayorderSetRequest) GetAlipayAuthCode() string {
    return r.alipayAuthCode
}

func (r *TaobaoBusTvmpayorderSetRequest) SetAlitripOrderId(alitripOrderId string) error {
    r.alitripOrderId = alitripOrderId
    r.Set("alitrip_order_id", alitripOrderId)
    return nil
}

func (r TaobaoBusTvmpayorderSetRequest) GetAlitripOrderId() string {
    return r.alitripOrderId
}

