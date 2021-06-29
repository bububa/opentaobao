package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国际机票大卖家Shopping推送 API请求
taobao.alitrip.ie.agent.shopping.push

用于国际机票大卖家主动推送Shopping结果更新缓存报价。
*/
type TaobaoAlitripIeAgentShoppingPushRequest struct {
    model.Params
    // 政策推送结构体
    _param0   *ShoppingPushRq
}

// 初始化TaobaoAlitripIeAgentShoppingPushRequest对象
func NewTaobaoAlitripIeAgentShoppingPushRequest() *TaobaoAlitripIeAgentShoppingPushRequest{
    return &TaobaoAlitripIeAgentShoppingPushRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentShoppingPushRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.shopping.push"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentShoppingPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 政策推送结构体
func (r *TaobaoAlitripIeAgentShoppingPushRequest) SetParam0(_param0 *ShoppingPushRq) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TaobaoAlitripIeAgentShoppingPushRequest) GetParam0() *ShoppingPushRq {
    return r._param0
}
