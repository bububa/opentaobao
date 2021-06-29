package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取改签单详情v2--增加鉴权校验 API请求
taobao.train.agent.change.get.vtwo

卖家获取待处理的改签单详情
*/
type TaobaoTrainAgentChangeGetVtwoRequest struct {
    model.Params
    // 代理商id
    agentId   int64
    // 申请单id
    applyId   int64
}

// 初始化TaobaoTrainAgentChangeGetVtwoRequest对象
func NewTaobaoTrainAgentChangeGetVtwoRequest() *TaobaoTrainAgentChangeGetVtwoRequest{
    return &TaobaoTrainAgentChangeGetVtwoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentChangeGetVtwoRequest) GetApiMethodName() string {
    return "taobao.train.agent.change.get.vtwo"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentChangeGetVtwoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentChangeGetVtwoRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

// AgentId Getter
func (r TaobaoTrainAgentChangeGetVtwoRequest) GetAgentId() int64 {
    return r.agentId
}
// ApplyId Setter
// 申请单id
func (r *TaobaoTrainAgentChangeGetVtwoRequest) SetApplyId(applyId int64) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

// ApplyId Getter
func (r TaobaoTrainAgentChangeGetVtwoRequest) GetApplyId() int64 {
    return r.applyId
}
