package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
廉航辅营正向订单查询详情接口 APIRequest
alitrip.tripvp.agent.order.get

【国际机票】查询辅营订单详情
*/
type AlitripTripvpAgentOrderGetRequest struct {
    model.Params

    // 代理商ID
    agentId   int64 

    // 辅营的订单号
    tradeOrderId   int64 

}

func NewAlitripTripvpAgentOrderGetRequest() *AlitripTripvpAgentOrderGetRequest{
    return &AlitripTripvpAgentOrderGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTripvpAgentOrderGetRequest) GetApiMethodName() string {
    return "alitrip.tripvp.agent.order.get"
}

func (r AlitripTripvpAgentOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTripvpAgentOrderGetRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r AlitripTripvpAgentOrderGetRequest) GetAgentId() int64 {
    return r.agentId
}

func (r *AlitripTripvpAgentOrderGetRequest) SetTradeOrderId(tradeOrderId int64) error {
    r.tradeOrderId = tradeOrderId
    r.Set("trade_order_id", tradeOrderId)
    return nil
}

func (r AlitripTripvpAgentOrderGetRequest) GetTradeOrderId() int64 {
    return r.tradeOrderId
}

