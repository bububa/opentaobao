package jipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商订单】订单详情 API请求
taobao.jipiao.agent.order.detail

根据淘宝系统订单号获取订单详情信息
*/
type TaobaoJipiaoAgentOrderDetailRequest struct {
    model.Params
    // 淘宝订单id列表，当前支持列表长度为1，即当前只支持单个订单详情查询
    orderIds   []int64
}

// 初始化TaobaoJipiaoAgentOrderDetailRequest对象
func NewTaobaoJipiaoAgentOrderDetailRequest() *TaobaoJipiaoAgentOrderDetailRequest{
    return &TaobaoJipiaoAgentOrderDetailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJipiaoAgentOrderDetailRequest) GetApiMethodName() string {
    return "taobao.jipiao.agent.order.detail"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJipiaoAgentOrderDetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderIds Setter
// 淘宝订单id列表，当前支持列表长度为1，即当前只支持单个订单详情查询
func (r *TaobaoJipiaoAgentOrderDetailRequest) SetOrderIds(orderIds []int64) error {
    r.orderIds = orderIds
    r.Set("order_ids", orderIds)
    return nil
}

// OrderIds Getter
func (r TaobaoJipiaoAgentOrderDetailRequest) GetOrderIds() []int64 {
    return r.orderIds
}
