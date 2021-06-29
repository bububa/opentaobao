package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
自助机条形码被动支付 API请求
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

// 初始化TaobaoBusTvmpayorderSetRequest对象
func NewTaobaoBusTvmpayorderSetRequest() *TaobaoBusTvmpayorderSetRequest{
    return &TaobaoBusTvmpayorderSetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusTvmpayorderSetRequest) GetApiMethodName() string {
    return "taobao.bus.tvmpayorder.set"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusTvmpayorderSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlipayAuthCode Setter
// 条形码认证码
func (r *TaobaoBusTvmpayorderSetRequest) SetAlipayAuthCode(alipayAuthCode string) error {
    r.alipayAuthCode = alipayAuthCode
    r.Set("alipay_auth_code", alipayAuthCode)
    return nil
}

// AlipayAuthCode Getter
func (r TaobaoBusTvmpayorderSetRequest) GetAlipayAuthCode() string {
    return r.alipayAuthCode
}
// AlitripOrderId Setter
// 飞猪订单号
func (r *TaobaoBusTvmpayorderSetRequest) SetAlitripOrderId(alitripOrderId string) error {
    r.alitripOrderId = alitripOrderId
    r.Set("alitrip_order_id", alitripOrderId)
    return nil
}

// AlitripOrderId Getter
func (r TaobaoBusTvmpayorderSetRequest) GetAlitripOrderId() string {
    return r.alitripOrderId
}
