package jipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商订单】订单回填票号/成功订单 APIRequest
taobao.jipiao.agent.order.ticket

淘宝机票代理商回填票号/成功订单
*/
type TaobaoJipiaoAgentOrderTicketRequest struct {
    model.Params

    // 淘宝系统订单id
    orderId   int64 

    // 成功订单参数：列表元素结构为——<br/>1.航班号（注：是订单里的航班号，非实际承运航班号）;<br/>2.旧乘机人姓名;<br/>3.新乘机人姓名;<br/>4.票号 （乘机人、航段对应的票号）<br/>说明：<br/>1.往返订单，2段航班对应1个票号的，需要2条success_info记录，分别对应去程、回程；<br/>2.有时用户输入的乘机人姓名输入错误或者有生僻字，代理商必须输入新的名字以保证验真通过；即旧乘机人姓名和新乘机人姓名不需要变化时可以相同
    successInfo   []string 

}

func NewTaobaoJipiaoAgentOrderTicketRequest() *TaobaoJipiaoAgentOrderTicketRequest{
    return &TaobaoJipiaoAgentOrderTicketRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJipiaoAgentOrderTicketRequest) GetApiMethodName() string {
    return "taobao.jipiao.agent.order.ticket"
}

func (r TaobaoJipiaoAgentOrderTicketRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJipiaoAgentOrderTicketRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TaobaoJipiaoAgentOrderTicketRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *TaobaoJipiaoAgentOrderTicketRequest) SetSuccessInfo(successInfo []string) error {
    r.successInfo = successInfo
    r.Set("success_info", successInfo)
    return nil
}

func (r TaobaoJipiaoAgentOrderTicketRequest) GetSuccessInfo() []string {
    return r.successInfo
}

