package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票】查询订单详情 API请求
taobao.alitrip.ie.agent.order.get

根据订单ID查询订单详情
*/
type TaobaoAlitripIeAgentOrderGetRequest struct {
    model.Params
    // 代理商ID
    _agentId   int64
    // 交易订单ID
    _tradeOrderId   int64
}

// 初始化TaobaoAlitripIeAgentOrderGetRequest对象
func NewTaobaoAlitripIeAgentOrderGetRequest() *TaobaoAlitripIeAgentOrderGetRequest{
    return &TaobaoAlitripIeAgentOrderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentOrderGetRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.order.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AgentId Setter
// 代理商ID
func (r *TaobaoAlitripIeAgentOrderGetRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r TaobaoAlitripIeAgentOrderGetRequest) GetAgentId() int64 {
    return r._agentId
}
// TradeOrderId Setter
// 交易订单ID
func (r *TaobaoAlitripIeAgentOrderGetRequest) SetTradeOrderId(_tradeOrderId int64) error {
    r._tradeOrderId = _tradeOrderId
    r.Set("trade_order_id", _tradeOrderId)
    return nil
}

// TradeOrderId Getter
func (r TaobaoAlitripIeAgentOrderGetRequest) GetTradeOrderId() int64 {
    return r._tradeOrderId
}
