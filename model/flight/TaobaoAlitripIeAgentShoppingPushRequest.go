package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国际机票大卖家Shopping推送 APIRequest
taobao.alitrip.ie.agent.shopping.push

用于国际机票大卖家主动推送Shopping结果更新缓存报价。
*/
type TaobaoAlitripIeAgentShoppingPushRequest struct {
    model.Params

    // 政策推送结构体
    param0   *ShoppingPushRq 

}

func NewTaobaoAlitripIeAgentShoppingPushRequest() *TaobaoAlitripIeAgentShoppingPushRequest{
    return &TaobaoAlitripIeAgentShoppingPushRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripIeAgentShoppingPushRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.shopping.push"
}

func (r TaobaoAlitripIeAgentShoppingPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripIeAgentShoppingPushRequest) SetParam0(param0 *ShoppingPushRq) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoAlitripIeAgentShoppingPushRequest) GetParam0() *ShoppingPushRq {
    return r.param0
}

