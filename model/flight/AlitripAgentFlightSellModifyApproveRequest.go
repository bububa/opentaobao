package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售改签确认 APIRequest
alitrip.agent.flight.sell.modify.approve

销售改签确认
*/
type AlitripAgentFlightSellModifyApproveRequest struct {
    model.Params

    // 入参对象
    param   *ModifyApproveRequestDto 

}

func NewAlitripAgentFlightSellModifyApproveRequest() *AlitripAgentFlightSellModifyApproveRequest{
    return &AlitripAgentFlightSellModifyApproveRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripAgentFlightSellModifyApproveRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.modify.approve"
}

func (r AlitripAgentFlightSellModifyApproveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripAgentFlightSellModifyApproveRequest) SetParam(param *ModifyApproveRequestDto) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlitripAgentFlightSellModifyApproveRequest) GetParam() *ModifyApproveRequestDto {
    return r.param
}

