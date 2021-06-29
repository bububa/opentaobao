package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售退票确认 API请求
alitrip.agent.flight.sell.refund.approve

销售退票确认
*/
type AlitripAgentFlightSellRefundApproveRequest struct {
    model.Params
    // 入参
    _param   *RefundApproveRequestDto
}

// 初始化AlitripAgentFlightSellRefundApproveRequest对象
func NewAlitripAgentFlightSellRefundApproveRequest() *AlitripAgentFlightSellRefundApproveRequest{
    return &AlitripAgentFlightSellRefundApproveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellRefundApproveRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.refund.approve"
}

// IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellRefundApproveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *AlitripAgentFlightSellRefundApproveRequest) SetParam(_param *RefundApproveRequestDto) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlitripAgentFlightSellRefundApproveRequest) GetParam() *RefundApproveRequestDto {
    return r._param
}
