package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售退票单列表 API请求
alitrip.agent.flight.sell.refund.list

销售退票单列表
*/
type AlitripAgentFlightSellRefundListRequest struct {
    model.Params
    // 请求对象
    param   *RefundListRequestDto
}

// 初始化AlitripAgentFlightSellRefundListRequest对象
func NewAlitripAgentFlightSellRefundListRequest() *AlitripAgentFlightSellRefundListRequest{
    return &AlitripAgentFlightSellRefundListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellRefundListRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.refund.list"
}

// IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellRefundListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 请求对象
func (r *AlitripAgentFlightSellRefundListRequest) SetParam(param *RefundListRequestDto) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlitripAgentFlightSellRefundListRequest) GetParam() *RefundListRequestDto {
    return r.param
}
