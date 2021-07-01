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
type TaobaoBusTvmpayorderSetAPIRequest struct {
    model.Params
    // 条形码认证码
    _alipayAuthCode   string
    // 飞猪订单号
    _alitripOrderId   string
}

// 初始化TaobaoBusTvmpayorderSetAPIRequest对象
func NewTaobaoBusTvmpayorderSetRequest() *TaobaoBusTvmpayorderSetAPIRequest{
    return &TaobaoBusTvmpayorderSetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusTvmpayorderSetAPIRequest) GetApiMethodName() string {
    return "taobao.bus.tvmpayorder.set"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusTvmpayorderSetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlipayAuthCode Setter
// 条形码认证码
func (r *TaobaoBusTvmpayorderSetAPIRequest) SetAlipayAuthCode(_alipayAuthCode string) error {
    r._alipayAuthCode = _alipayAuthCode
    r.Set("alipay_auth_code", _alipayAuthCode)
    return nil
}

// AlipayAuthCode Getter
func (r TaobaoBusTvmpayorderSetAPIRequest) GetAlipayAuthCode() string {
    return r._alipayAuthCode
}
// AlitripOrderId Setter
// 飞猪订单号
func (r *TaobaoBusTvmpayorderSetAPIRequest) SetAlitripOrderId(_alitripOrderId string) error {
    r._alitripOrderId = _alitripOrderId
    r.Set("alitrip_order_id", _alitripOrderId)
    return nil
}

// AlitripOrderId Getter
func (r TaobaoBusTvmpayorderSetAPIRequest) GetAlitripOrderId() string {
    return r._alitripOrderId
}
