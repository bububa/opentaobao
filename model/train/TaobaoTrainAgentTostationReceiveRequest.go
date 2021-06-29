package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下票送票至车站代理商确认用户已取票服务 API请求
taobao.train.agent.tostation.receive

送票至车站的订单，代理商确认用户已取票
*/
type TaobaoTrainAgentTostationReceiveRequest struct {
    model.Params
    // 淘宝的主订单号
    _mainOrderId   int64
    // 代理商id
    _agentId   int64
}

// 初始化TaobaoTrainAgentTostationReceiveRequest对象
func NewTaobaoTrainAgentTostationReceiveRequest() *TaobaoTrainAgentTostationReceiveRequest{
    return &TaobaoTrainAgentTostationReceiveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentTostationReceiveRequest) GetApiMethodName() string {
    return "taobao.train.agent.tostation.receive"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentTostationReceiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainOrderId Setter
// 淘宝的主订单号
func (r *TaobaoTrainAgentTostationReceiveRequest) SetMainOrderId(_mainOrderId int64) error {
    r._mainOrderId = _mainOrderId
    r.Set("main_order_id", _mainOrderId)
    return nil
}

// MainOrderId Getter
func (r TaobaoTrainAgentTostationReceiveRequest) GetMainOrderId() int64 {
    return r._mainOrderId
}
// AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentTostationReceiveRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r TaobaoTrainAgentTostationReceiveRequest) GetAgentId() int64 {
    return r._agentId
}
