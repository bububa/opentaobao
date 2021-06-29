package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商获取待出票订单列表 API请求
taobao.train.agent.bookorders.get

代理商获取待出票订单列表，只返回订单号
*/
type TaobaoTrainAgentBookordersGetRequest struct {
    model.Params
    // 代理商id
    agentId   int64
}

// 初始化TaobaoTrainAgentBookordersGetRequest对象
func NewTaobaoTrainAgentBookordersGetRequest() *TaobaoTrainAgentBookordersGetRequest{
    return &TaobaoTrainAgentBookordersGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentBookordersGetRequest) GetApiMethodName() string {
    return "taobao.train.agent.bookorders.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentBookordersGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentBookordersGetRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

// AgentId Getter
func (r TaobaoTrainAgentBookordersGetRequest) GetAgentId() int64 {
    return r.agentId
}
