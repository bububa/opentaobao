package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取待退票的订单v2--增加鉴权校验 API请求
taobao.train.agent.returnorders.get.vtwo

代理商用来获取待退票的订单列表及数量，防止代理商掉单。
*/
type TaobaoTrainAgentReturnordersGetVtwoRequest struct {
    model.Params
    // 卖家ID
    agentId   int64
    // 0 线上退票 1线下退票
    offline   int64
}

// 初始化TaobaoTrainAgentReturnordersGetVtwoRequest对象
func NewTaobaoTrainAgentReturnordersGetVtwoRequest() *TaobaoTrainAgentReturnordersGetVtwoRequest{
    return &TaobaoTrainAgentReturnordersGetVtwoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentReturnordersGetVtwoRequest) GetApiMethodName() string {
    return "taobao.train.agent.returnorders.get.vtwo"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentReturnordersGetVtwoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AgentId Setter
// 卖家ID
func (r *TaobaoTrainAgentReturnordersGetVtwoRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

// AgentId Getter
func (r TaobaoTrainAgentReturnordersGetVtwoRequest) GetAgentId() int64 {
    return r.agentId
}
// Offline Setter
// 0 线上退票 1线下退票
func (r *TaobaoTrainAgentReturnordersGetVtwoRequest) SetOffline(offline int64) error {
    r.offline = offline
    r.Set("offline", offline)
    return nil
}

// Offline Getter
func (r TaobaoTrainAgentReturnordersGetVtwoRequest) GetOffline() int64 {
    return r.offline
}
