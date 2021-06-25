package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取待改签订单v2--增加鉴权校验 APIRequest
taobao.train.agent.changeorders.get.vtwo

代理商用来获取待改签的订单列表及数量，防止代理商掉单。
*/
type TaobaoTrainAgentChangeordersGetVtwoRequest struct {
    model.Params

    // 卖家id
    agentId   int64 

}

func NewTaobaoTrainAgentChangeordersGetVtwoRequest() *TaobaoTrainAgentChangeordersGetVtwoRequest{
    return &TaobaoTrainAgentChangeordersGetVtwoRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTrainAgentChangeordersGetVtwoRequest) GetApiMethodName() string {
    return "taobao.train.agent.changeorders.get.vtwo"
}

func (r TaobaoTrainAgentChangeordersGetVtwoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTrainAgentChangeordersGetVtwoRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r TaobaoTrainAgentChangeordersGetVtwoRequest) GetAgentId() int64 {
    return r.agentId
}

