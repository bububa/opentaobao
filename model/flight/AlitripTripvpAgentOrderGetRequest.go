package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
廉航辅营正向订单查询详情接口 API请求
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

// 初始化AlitripTripvpAgentOrderGetRequest对象
func NewAlitripTripvpAgentOrderGetRequest() *AlitripTripvpAgentOrderGetRequest{
    return &AlitripTripvpAgentOrderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTripvpAgentOrderGetRequest) GetApiMethodName() string {
    return "alitrip.tripvp.agent.order.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTripvpAgentOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AgentId Setter
// 代理商ID
func (r *AlitripTripvpAgentOrderGetRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

// AgentId Getter
func (r AlitripTripvpAgentOrderGetRequest) GetAgentId() int64 {
    return r.agentId
}
// TradeOrderId Setter
// 辅营的订单号
func (r *AlitripTripvpAgentOrderGetRequest) SetTradeOrderId(tradeOrderId int64) error {
    r.tradeOrderId = tradeOrderId
    r.Set("trade_order_id", tradeOrderId)
    return nil
}

// TradeOrderId Getter
func (r AlitripTripvpAgentOrderGetRequest) GetTradeOrderId() int64 {
    return r.tradeOrderId
}
