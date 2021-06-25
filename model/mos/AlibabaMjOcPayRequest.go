package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
POS收银成功后订单同步 APIRequest
alibaba.mj.oc.pay

此API用于在银泰商场中，消费者在收银台收银/退款时， POS系统在收银或退款成功后，调用此接口进行订单同步
*/
type AlibabaMjOcPayRequest struct {
    model.Params

    // 订单数据
    posOrder   *PosOrderDto 

}

func NewAlibabaMjOcPayRequest() *AlibabaMjOcPayRequest{
    return &AlibabaMjOcPayRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMjOcPayRequest) GetApiMethodName() string {
    return "alibaba.mj.oc.pay"
}

func (r AlibabaMjOcPayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMjOcPayRequest) SetPosOrder(posOrder *PosOrderDto) error {
    r.posOrder = posOrder
    r.Set("pos_order", posOrder)
    return nil
}

func (r AlibabaMjOcPayRequest) GetPosOrder() *PosOrderDto {
    return r.posOrder
}

