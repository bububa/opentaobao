package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家汽车票车次更新通知接口 APIRequest
taobao.bus.busnumber.set

商家汽车票车次更新后，调用该接口通知平台。
*/
type TaobaoBusBusnumberSetRequest struct {
    model.Params

    // 车次更新通知参数
    pushParam   *TopBusNumerPushRq 

}

func NewTaobaoBusBusnumberSetRequest() *TaobaoBusBusnumberSetRequest{
    return &TaobaoBusBusnumberSetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusBusnumberSetRequest) GetApiMethodName() string {
    return "taobao.bus.busnumber.set"
}

func (r TaobaoBusBusnumberSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBusBusnumberSetRequest) SetPushParam(pushParam *TopBusNumerPushRq) error {
    r.pushParam = pushParam
    r.Set("push_param", pushParam)
    return nil
}

func (r TaobaoBusBusnumberSetRequest) GetPushParam() *TopBusNumerPushRq {
    return r.pushParam
}

