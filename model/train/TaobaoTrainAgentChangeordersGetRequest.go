package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取待改签订单 APIRequest
taobao.train.agent.changeorders.get

代理商用来获取待改签的订单列表及数量，防止代理商掉单。
*/
type TaobaoTrainAgentChangeordersGetRequest struct {
    model.Params

    // 卖家id
    agentId   int64 

}

func NewTaobaoTrainAgentChangeordersGetRequest() *TaobaoTrainAgentChangeordersGetRequest{
    return &TaobaoTrainAgentChangeordersGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTrainAgentChangeordersGetRequest) GetApiMethodName() string {
    return "taobao.train.agent.changeorders.get"
}

func (r TaobaoTrainAgentChangeordersGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTrainAgentChangeordersGetRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r TaobaoTrainAgentChangeordersGetRequest) GetAgentId() int64 {
    return r.agentId
}

