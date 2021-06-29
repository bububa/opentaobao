package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票】查询订单详情 APIRequest
taobao.alitrip.ie.agent.order.get

根据订单ID查询订单详情
*/
type TaobaoAlitripIeAgentOrderGetRequest struct {
    model.Params

    // 代理商ID
    agentId   int64 

    // 交易订单ID
    tradeOrderId   int64 

}

func NewTaobaoAlitripIeAgentOrderGetRequest() *TaobaoAlitripIeAgentOrderGetRequest{
    return &TaobaoAlitripIeAgentOrderGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripIeAgentOrderGetRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.order.get"
}

func (r TaobaoAlitripIeAgentOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripIeAgentOrderGetRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r TaobaoAlitripIeAgentOrderGetRequest) GetAgentId() int64 {
    return r.agentId
}

func (r *TaobaoAlitripIeAgentOrderGetRequest) SetTradeOrderId(tradeOrderId int64) error {
    r.tradeOrderId = tradeOrderId
    r.Set("trade_order_id", tradeOrderId)
    return nil
}

func (r TaobaoAlitripIeAgentOrderGetRequest) GetTradeOrderId() int64 {
    return r.tradeOrderId
}

