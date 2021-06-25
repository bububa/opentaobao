package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售退票确认 APIRequest
alitrip.agent.flight.sell.refund.approve

销售退票确认
*/
type AlitripAgentFlightSellRefundApproveRequest struct {
    model.Params

    // 入参
    param   *RefundApproveRequestDto 

}

func NewAlitripAgentFlightSellRefundApproveRequest() *AlitripAgentFlightSellRefundApproveRequest{
    return &AlitripAgentFlightSellRefundApproveRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripAgentFlightSellRefundApproveRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.refund.approve"
}

func (r AlitripAgentFlightSellRefundApproveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripAgentFlightSellRefundApproveRequest) SetParam(param *RefundApproveRequestDto) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlitripAgentFlightSellRefundApproveRequest) GetParam() *RefundApproveRequestDto {
    return r.param
}

