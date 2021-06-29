package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下票确认送票至车站服务 API请求
taobao.train.agent.tostation.confirm

送票至车站的订单，代理商确认配送到站
*/
type TaobaoTrainAgentTostationConfirmRequest struct {
    model.Params
    // 淘宝的主订单号
    mainOrderId   int64
    // 代理商id
    agentId   int64
}

// 初始化TaobaoTrainAgentTostationConfirmRequest对象
func NewTaobaoTrainAgentTostationConfirmRequest() *TaobaoTrainAgentTostationConfirmRequest{
    return &TaobaoTrainAgentTostationConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentTostationConfirmRequest) GetApiMethodName() string {
    return "taobao.train.agent.tostation.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentTostationConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainOrderId Setter
// 淘宝的主订单号
func (r *TaobaoTrainAgentTostationConfirmRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

// MainOrderId Getter
func (r TaobaoTrainAgentTostationConfirmRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}
// AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentTostationConfirmRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

// AgentId Getter
func (r TaobaoTrainAgentTostationConfirmRequest) GetAgentId() int64 {
    return r.agentId
}
