package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下自助机未付款取消订单 API请求
taobao.bus.tvmcancelorder.set

自助机汽车票未付款取消订单
*/
type TaobaoBusTvmcancelorderSetRequest struct {
    model.Params
    // 飞猪订单号
    _alitripOrderId   string
}

// 初始化TaobaoBusTvmcancelorderSetRequest对象
func NewTaobaoBusTvmcancelorderSetRequest() *TaobaoBusTvmcancelorderSetRequest{
    return &TaobaoBusTvmcancelorderSetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusTvmcancelorderSetRequest) GetApiMethodName() string {
    return "taobao.bus.tvmcancelorder.set"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusTvmcancelorderSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlitripOrderId Setter
// 飞猪订单号
func (r *TaobaoBusTvmcancelorderSetRequest) SetAlitripOrderId(_alitripOrderId string) error {
    r._alitripOrderId = _alitripOrderId
    r.Set("alitrip_order_id", _alitripOrderId)
    return nil
}

// AlitripOrderId Getter
func (r TaobaoBusTvmcancelorderSetRequest) GetAlitripOrderId() string {
    return r._alitripOrderId
}
