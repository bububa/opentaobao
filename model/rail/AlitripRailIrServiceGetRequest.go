package rail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国际火车票仓位坐席查询 APIRequest
alitrip.rail.ir.service.get

国际火车票标准仓位坐席查询
*/
type AlitripRailIrServiceGetRequest struct {
    model.Params

    // 6代表境外火车票
    bizType   int64 

    // 代理商id
    agentId   int64 

}

func NewAlitripRailIrServiceGetRequest() *AlitripRailIrServiceGetRequest{
    return &AlitripRailIrServiceGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripRailIrServiceGetRequest) GetApiMethodName() string {
    return "alitrip.rail.ir.service.get"
}

func (r AlitripRailIrServiceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripRailIrServiceGetRequest) SetBizType(bizType int64) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r AlitripRailIrServiceGetRequest) GetBizType() int64 {
    return r.bizType
}

func (r *AlitripRailIrServiceGetRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r AlitripRailIrServiceGetRequest) GetAgentId() int64 {
    return r.agentId
}

