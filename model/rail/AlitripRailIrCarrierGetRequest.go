package rail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国际火车票铁路承运公司查询 APIRequest
alitrip.rail.ir.carrier.get

国际火车票提供给代理商用于查询标准铁路承运公司carrier信息，用于代理商自己的carrier与飞猪平台的carrier做映射
*/
type AlitripRailIrCarrierGetRequest struct {
    model.Params

    // 商家id
    agentId   int64 

}

func NewAlitripRailIrCarrierGetRequest() *AlitripRailIrCarrierGetRequest{
    return &AlitripRailIrCarrierGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripRailIrCarrierGetRequest) GetApiMethodName() string {
    return "alitrip.rail.ir.carrier.get"
}

func (r AlitripRailIrCarrierGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripRailIrCarrierGetRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r AlitripRailIrCarrierGetRequest) GetAgentId() int64 {
    return r.agentId
}

