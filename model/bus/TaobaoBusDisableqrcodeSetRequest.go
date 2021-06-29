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
type TaobaoBusDisableqrcodeSetRequest struct {
    model.Params
    // 飞猪订单号
    _alitripOrderId   string
}

// 初始化TaobaoBusDisableqrcodeSetRequest对象
func NewTaobaoBusDisableqrcodeSetRequest() *TaobaoBusDisableqrcodeSetRequest{
    return &TaobaoBusDisableqrcodeSetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusDisableqrcodeSetRequest) GetApiMethodName() string {
    return "taobao.bus.disableqrcode.set"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusDisableqrcodeSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlitripOrderId Setter
// 飞猪订单号
func (r *TaobaoBusDisableqrcodeSetRequest) SetAlitripOrderId(_alitripOrderId string) error {
    r._alitripOrderId = _alitripOrderId
    r.Set("alitrip_order_id", _alitripOrderId)
    return nil
}

// AlitripOrderId Getter
func (r TaobaoBusDisableqrcodeSetRequest) GetAlitripOrderId() string {
    return r._alitripOrderId
}
