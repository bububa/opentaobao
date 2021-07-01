package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取待改签订单v2--增加鉴权校验 API请求
taobao.train.agent.changeorders.get.vtwo

代理商用来获取待改签的订单列表及数量，防止代理商掉单。
*/
type TaobaoTrainAgentChangeordersGetVtwoAPIRequest struct {
    model.Params
    // 卖家id
    _agentId   int64
}

// 初始化TaobaoTrainAgentChangeordersGetVtwoAPIRequest对象
func NewTaobaoTrainAgentChangeordersGetVtwoRequest() *TaobaoTrainAgentChangeordersGetVtwoAPIRequest{
    return &TaobaoTrainAgentChangeordersGetVtwoAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentChangeordersGetVtwoAPIRequest) GetApiMethodName() string {
    return "taobao.train.agent.changeorders.get.vtwo"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentChangeordersGetVtwoAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AgentId Setter
// 卖家id
func (r *TaobaoTrainAgentChangeordersGetVtwoAPIRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r TaobaoTrainAgentChangeordersGetVtwoAPIRequest) GetAgentId() int64 {
    return r._agentId
}
