package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售出票列表 API请求
alitrip.agent.flight.sell.ticketing.list

销售出票列表
*/
type AlitripAgentFlightSellTicketingListRequest struct {
    model.Params
    // 入参对象
    _param   *TicketingListRequestDto
}

// 初始化AlitripAgentFlightSellTicketingListRequest对象
func NewAlitripAgentFlightSellTicketingListRequest() *AlitripAgentFlightSellTicketingListRequest{
    return &AlitripAgentFlightSellTicketingListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellTicketingListRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.ticketing.list"
}

// IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellTicketingListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参对象
func (r *AlitripAgentFlightSellTicketingListRequest) SetParam(_param *TicketingListRequestDto) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlitripAgentFlightSellTicketingListRequest) GetParam() *TicketingListRequestDto {
    return r._param
}
