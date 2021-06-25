package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售改签回填 APIRequest
alitrip.agent.flight.sell.modify.backfill

销售改签回填
*/
type AlitripAgentFlightSellModifyBackfillRequest struct {
    model.Params

    // 入参
    param   *ModifyBackFillRequestDto 

}

func NewAlitripAgentFlightSellModifyBackfillRequest() *AlitripAgentFlightSellModifyBackfillRequest{
    return &AlitripAgentFlightSellModifyBackfillRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripAgentFlightSellModifyBackfillRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.modify.backfill"
}

func (r AlitripAgentFlightSellModifyBackfillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripAgentFlightSellModifyBackfillRequest) SetParam(param *ModifyBackFillRequestDto) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlitripAgentFlightSellModifyBackfillRequest) GetParam() *ModifyBackFillRequestDto {
    return r.param
}

