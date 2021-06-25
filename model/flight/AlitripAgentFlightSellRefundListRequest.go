package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售退票单列表 APIRequest
alitrip.agent.flight.sell.refund.list

销售退票单列表
*/
type AlitripAgentFlightSellRefundListRequest struct {
    model.Params

    // 请求对象
    param   *RefundListRequestDto 

}

func NewAlitripAgentFlightSellRefundListRequest() *AlitripAgentFlightSellRefundListRequest{
    return &AlitripAgentFlightSellRefundListRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripAgentFlightSellRefundListRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.refund.list"
}

func (r AlitripAgentFlightSellRefundListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripAgentFlightSellRefundListRequest) SetParam(param *RefundListRequestDto) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlitripAgentFlightSellRefundListRequest) GetParam() *RefundListRequestDto {
    return r.param
}

