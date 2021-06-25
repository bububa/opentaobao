package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商获取待出票订单列表 APIRequest
taobao.train.agent.bookorders.get

代理商获取待出票订单列表，只返回订单号
*/
type TaobaoTrainAgentBookordersGetRequest struct {
    model.Params

    // 代理商id
    agentId   int64 

}

func NewTaobaoTrainAgentBookordersGetRequest() *TaobaoTrainAgentBookordersGetRequest{
    return &TaobaoTrainAgentBookordersGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTrainAgentBookordersGetRequest) GetApiMethodName() string {
    return "taobao.train.agent.bookorders.get"
}

func (r TaobaoTrainAgentBookordersGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTrainAgentBookordersGetRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r TaobaoTrainAgentBookordersGetRequest) GetAgentId() int64 {
    return r.agentId
}

