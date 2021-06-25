package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商获取待出票订单列表v2--增加鉴权校验 APIRequest
taobao.train.agent.bookorders.get.vtwo

代理商获取待出票订单列表，只返回订单号
*/
type TaobaoTrainAgentBookordersGetVtwoRequest struct {
    model.Params

    // 代理商id
    agentId   int64 

}

func NewTaobaoTrainAgentBookordersGetVtwoRequest() *TaobaoTrainAgentBookordersGetVtwoRequest{
    return &TaobaoTrainAgentBookordersGetVtwoRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTrainAgentBookordersGetVtwoRequest) GetApiMethodName() string {
    return "taobao.train.agent.bookorders.get.vtwo"
}

func (r TaobaoTrainAgentBookordersGetVtwoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTrainAgentBookordersGetVtwoRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r TaobaoTrainAgentBookordersGetVtwoRequest) GetAgentId() int64 {
    return r.agentId
}

