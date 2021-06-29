package rail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国际火车票铁路承运公司查询 API请求
alitrip.rail.ir.carrier.get

国际火车票提供给代理商用于查询标准铁路承运公司carrier信息，用于代理商自己的carrier与飞猪平台的carrier做映射
*/
type AlitripRailIrCarrierGetRequest struct {
    model.Params
    // 商家id
    agentId   int64
}

// 初始化AlitripRailIrCarrierGetRequest对象
func NewAlitripRailIrCarrierGetRequest() *AlitripRailIrCarrierGetRequest{
    return &AlitripRailIrCarrierGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripRailIrCarrierGetRequest) GetApiMethodName() string {
    return "alitrip.rail.ir.carrier.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripRailIrCarrierGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AgentId Setter
// 商家id
func (r *AlitripRailIrCarrierGetRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

// AgentId Getter
func (r AlitripRailIrCarrierGetRequest) GetAgentId() int64 {
    return r.agentId
}
