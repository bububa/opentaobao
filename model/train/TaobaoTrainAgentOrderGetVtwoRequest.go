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
type TaobaoTrainAgentOrderGetVtwoRequest struct {
    model.Params
    // 淘宝的主订单号
    _mainOrderId   int64
    // 代理商id
    _agentId   int64
}

// 初始化TaobaoTrainAgentOrderGetVtwoRequest对象
func NewTaobaoTrainAgentOrderGetVtwoRequest() *TaobaoTrainAgentOrderGetVtwoRequest{
    return &TaobaoTrainAgentOrderGetVtwoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentOrderGetVtwoRequest) GetApiMethodName() string {
    return "taobao.train.agent.order.get.vtwo"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentOrderGetVtwoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainOrderId Setter
// 淘宝的主订单号
func (r *TaobaoTrainAgentOrderGetVtwoRequest) SetMainOrderId(_mainOrderId int64) error {
    r._mainOrderId = _mainOrderId
    r.Set("main_order_id", _mainOrderId)
    return nil
}

// MainOrderId Getter
func (r TaobaoTrainAgentOrderGetVtwoRequest) GetMainOrderId() int64 {
    return r._mainOrderId
}
// AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentOrderGetVtwoRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r TaobaoTrainAgentOrderGetVtwoRequest) GetAgentId() int64 {
    return r._agentId
}
