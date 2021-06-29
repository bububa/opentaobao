package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取待改签订单 API请求
taobao.train.agent.changeorders.get

代理商用来获取待改签的订单列表及数量，防止代理商掉单。
*/
type TaobaoTrainAgentChangeordersGetRequest struct {
    model.Params
    // 卖家id
    _agentId   int64
}

// 初始化TaobaoTrainAgentChangeordersGetRequest对象
func NewTaobaoTrainAgentChangeordersGetRequest() *TaobaoTrainAgentChangeordersGetRequest{
    return &TaobaoTrainAgentChangeordersGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentChangeordersGetRequest) GetApiMethodName() string {
    return "taobao.train.agent.changeorders.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentChangeordersGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AgentId Setter
// 卖家id
func (r *TaobaoTrainAgentChangeordersGetRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r TaobaoTrainAgentChangeordersGetRequest) GetAgentId() int64 {
    return r._agentId
}
