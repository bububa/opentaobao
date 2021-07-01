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
type AlitripAgentFlightSellRefundListAPIRequest struct {
    model.Params
    // 请求对象
    _param   *RefundListRequestDto
}

// 初始化AlitripAgentFlightSellRefundListAPIRequest对象
func NewAlitripAgentFlightSellRefundListRequest() *AlitripAgentFlightSellRefundListAPIRequest{
    return &AlitripAgentFlightSellRefundListAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellRefundListAPIRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.refund.list"
}

// IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellRefundListAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 请求对象
func (r *AlitripAgentFlightSellRefundListAPIRequest) SetParam(_param *RefundListRequestDto) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlitripAgentFlightSellRefundListAPIRequest) GetParam() *RefundListRequestDto {
    return r._param
}
