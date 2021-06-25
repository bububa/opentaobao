package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
B2B退票申请接口 APIRequest
taobao.bus.refund.set

B2B业务支持退票
*/
type TaobaoBusRefundSetRequest struct {
    model.Params

    // 入参
    param0   *B2BRefundOrderRq 

}

func NewTaobaoBusRefundSetRequest() *TaobaoBusRefundSetRequest{
    return &TaobaoBusRefundSetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusRefundSetRequest) GetApiMethodName() string {
    return "taobao.bus.refund.set"
}

func (r TaobaoBusRefundSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBusRefundSetRequest) SetParam0(param0 *B2BRefundOrderRq) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoBusRefundSetRequest) GetParam0() *B2BRefundOrderRq {
    return r.param0
}

