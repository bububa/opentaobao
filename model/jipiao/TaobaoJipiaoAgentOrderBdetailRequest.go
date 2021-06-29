package jipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商订单】采购订单详情 API请求
taobao.jipiao.agent.order.bdetail

根据淘宝系统订单号获取订单详情信息
*/
type TaobaoJipiaoAgentOrderBdetailRequest struct {
    model.Params
    // 订单号
    orderId   int64
}

// 初始化TaobaoJipiaoAgentOrderBdetailRequest对象
func NewTaobaoJipiaoAgentOrderBdetailRequest() *TaobaoJipiaoAgentOrderBdetailRequest{
    return &TaobaoJipiaoAgentOrderBdetailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJipiaoAgentOrderBdetailRequest) GetApiMethodName() string {
    return "taobao.jipiao.agent.order.bdetail"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJipiaoAgentOrderBdetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单号
func (r *TaobaoJipiaoAgentOrderBdetailRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TaobaoJipiaoAgentOrderBdetailRequest) GetOrderId() int64 {
    return r.orderId
}
