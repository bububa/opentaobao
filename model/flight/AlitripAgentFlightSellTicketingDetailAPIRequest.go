package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightSellTicketingDetailAPIRequest 销售出票详情 API请求
// alitrip.agent.flight.sell.ticketing.detail
//
// 销售出票详情
type AlitripAgentFlightSellTicketingDetailAPIRequest struct {
	model.Params
	// 国内国际标识
	_domesticIntl int64
	// 飞猪订单号
	_orderId string
}

// NewAlitripAgentFlightSellTicketingDetailRequest 初始化AlitripAgentFlightSellTicketingDetailAPIRequest对象
func NewAlitripAgentFlightSellTicketingDetailRequest() *AlitripAgentFlightSellTicketingDetailAPIRequest {
	return &AlitripAgentFlightSellTicketingDetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellTicketingDetailAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.ticketing.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellTicketingDetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DomesticIntl Setter
// 国内国际标识
func (r *AlitripAgentFlightSellTicketingDetailAPIRequest) SetDomesticIntl(_domesticIntl int64) error {
	r._domesticIntl = _domesticIntl
	r.Set("domestic_intl", _domesticIntl)
	return nil
}

// Get DomesticIntl Getter
func (r AlitripAgentFlightSellTicketingDetailAPIRequest) GetDomesticIntl() int64 {
	return r._domesticIntl
}

// Set is OrderId Setter
// 飞猪订单号
func (r *AlitripAgentFlightSellTicketingDetailAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r AlitripAgentFlightSellTicketingDetailAPIRequest) GetOrderId() string {
	return r._orderId
}
