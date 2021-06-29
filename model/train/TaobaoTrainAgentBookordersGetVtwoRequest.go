package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商获取待出票订单列表v2--增加鉴权校验 API请求
taobao.train.agent.bookorders.get.vtwo

代理商获取待出票订单列表，只返回订单号
*/
type TaobaoTrainAgentBookordersGetVtwoRequest struct {
    model.Params
    // 代理商id
    agentId   int64
}

// 初始化TaobaoTrainAgentBookordersGetVtwoRequest对象
func NewTaobaoTrainAgentBookordersGetVtwoRequest() *TaobaoTrainAgentBookordersGetVtwoRequest{
    return &TaobaoTrainAgentBookordersGetVtwoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentBookordersGetVtwoRequest) GetApiMethodName() string {
    return "taobao.train.agent.bookorders.get.vtwo"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentBookordersGetVtwoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentBookordersGetVtwoRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

// AgentId Getter
func (r TaobaoTrainAgentBookordersGetVtwoRequest) GetAgentId() int64 {
    return r.agentId
}
