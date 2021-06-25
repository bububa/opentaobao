package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下自助机未付款取消订单 APIRequest
taobao.bus.tvmcancelorder.set

自助机汽车票未付款取消订单
*/
type TaobaoBusTvmcancelorderSetRequest struct {
    model.Params

    // 飞猪订单号
    alitripOrderId   string 

}

func NewTaobaoBusTvmcancelorderSetRequest() *TaobaoBusTvmcancelorderSetRequest{
    return &TaobaoBusTvmcancelorderSetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusTvmcancelorderSetRequest) GetApiMethodName() string {
    return "taobao.bus.tvmcancelorder.set"
}

func (r TaobaoBusTvmcancelorderSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBusTvmcancelorderSetRequest) SetAlitripOrderId(alitripOrderId string) error {
    r.alitripOrderId = alitripOrderId
    r.Set("alitrip_order_id", alitripOrderId)
    return nil
}

func (r TaobaoBusTvmcancelorderSetRequest) GetAlitripOrderId() string {
    return r.alitripOrderId
}

