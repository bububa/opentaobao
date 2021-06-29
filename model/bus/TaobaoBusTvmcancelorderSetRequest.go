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
    alitripOrderId   string
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
func (r *TaobaoBusTvmcancelorderSetRequest) SetAlitripOrderId(alitripOrderId string) error {
    r.alitripOrderId = alitripOrderId
    r.Set("alitrip_order_id", alitripOrderId)
    return nil
}

// AlitripOrderId Getter
func (r TaobaoBusTvmcancelorderSetRequest) GetAlitripOrderId() string {
    return r.alitripOrderId
}
