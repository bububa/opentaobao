package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商获取订单退票信息 API请求
taobao.train.agent.get.refund

代理商获取订单信息回调API
*/
type TaobaoTrainAgentGetRefundAPIRequest struct {
    model.Params
    // 淘宝的主订单号
    _mainOrderId   int64
    // 代理商id
    _agentId   int64
}

// 初始化TaobaoTrainAgentGetRefundAPIRequest对象
func NewTaobaoTrainAgentGetRefundRequest() *TaobaoTrainAgentGetRefundAPIRequest{
    return &TaobaoTrainAgentGetRefundAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentGetRefundAPIRequest) GetApiMethodName() string {
    return "taobao.train.agent.get.refund"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentGetRefundAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainOrderId Setter
// 淘宝的主订单号
func (r *TaobaoTrainAgentGetRefundAPIRequest) SetMainOrderId(_mainOrderId int64) error {
    r._mainOrderId = _mainOrderId
    r.Set("main_order_id", _mainOrderId)
    return nil
}

// MainOrderId Getter
func (r TaobaoTrainAgentGetRefundAPIRequest) GetMainOrderId() int64 {
    return r._mainOrderId
}
// AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentGetRefundAPIRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r TaobaoTrainAgentGetRefundAPIRequest) GetAgentId() int64 {
    return r._agentId
}
