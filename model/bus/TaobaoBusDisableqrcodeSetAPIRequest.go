package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
自助机失效二维码 API请求
taobao.bus.disableqrcode.set

使创建的二维码失效
*/
type TaobaoBusDisableqrcodeSetAPIRequest struct {
    model.Params
    // 飞猪订单号
    _alitripOrderId   string
}

// 初始化TaobaoBusDisableqrcodeSetAPIRequest对象
func NewTaobaoBusDisableqrcodeSetRequest() *TaobaoBusDisableqrcodeSetAPIRequest{
    return &TaobaoBusDisableqrcodeSetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusDisableqrcodeSetAPIRequest) GetApiMethodName() string {
    return "taobao.bus.disableqrcode.set"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusDisableqrcodeSetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlitripOrderId Setter
// 飞猪订单号
func (r *TaobaoBusDisableqrcodeSetAPIRequest) SetAlitripOrderId(_alitripOrderId string) error {
    r._alitripOrderId = _alitripOrderId
    r.Set("alitrip_order_id", _alitripOrderId)
    return nil
}

// AlitripOrderId Getter
func (r TaobaoBusDisableqrcodeSetAPIRequest) GetAlitripOrderId() string {
    return r._alitripOrderId
}
