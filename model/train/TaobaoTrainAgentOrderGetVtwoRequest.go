package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商获取订单信息回调APIv2--增加鉴权校验 API请求
taobao.train.agent.order.get.vtwo

代理商获取订单信息回调API
*/
type TaobaoTrainAgentOrderGetVtwoAPIRequest struct {
    model.Params
    // 淘宝的主订单号
    _mainOrderId   int64
    // 代理商id
    _agentId   int64
}

// 初始化TaobaoTrainAgentOrderGetVtwoAPIRequest对象
func NewTaobaoTrainAgentOrderGetVtwoRequest() *TaobaoTrainAgentOrderGetVtwoAPIRequest{
    return &TaobaoTrainAgentOrderGetVtwoAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentOrderGetVtwoAPIRequest) GetApiMethodName() string {
    return "taobao.train.agent.order.get.vtwo"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentOrderGetVtwoAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainOrderId Setter
// 淘宝的主订单号
func (r *TaobaoTrainAgentOrderGetVtwoAPIRequest) SetMainOrderId(_mainOrderId int64) error {
    r._mainOrderId = _mainOrderId
    r.Set("main_order_id", _mainOrderId)
    return nil
}

// MainOrderId Getter
func (r TaobaoTrainAgentOrderGetVtwoAPIRequest) GetMainOrderId() int64 {
    return r._mainOrderId
}
// AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentOrderGetVtwoAPIRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r TaobaoTrainAgentOrderGetVtwoAPIRequest) GetAgentId() int64 {
    return r._agentId
}
