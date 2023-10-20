package flight

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightSellTicketingDetailAPIRequest 销售出票详情 API请求
// alitrip.agent.flight.sell.ticketing.detail
//
// 销售出票详情
type AlitripAgentFlightSellTicketingDetailAPIRequest struct {
	model.Params
	// 飞猪订单号
	_orderId string
	// 国内国际标识
	_domesticIntl int64
}

// NewAlitripAgentFlightSellTicketingDetailRequest 初始化AlitripAgentFlightSellTicketingDetailAPIRequest对象
func NewAlitripAgentFlightSellTicketingDetailRequest() *AlitripAgentFlightSellTicketingDetailAPIRequest {
	return &AlitripAgentFlightSellTicketingDetailAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripAgentFlightSellTicketingDetailAPIRequest) Reset() {
	r._orderId = ""
	r._domesticIntl = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellTicketingDetailAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.ticketing.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellTicketingDetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripAgentFlightSellTicketingDetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 飞猪订单号
func (r *AlitripAgentFlightSellTicketingDetailAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitripAgentFlightSellTicketingDetailAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetDomesticIntl is DomesticIntl Setter
// 国内国际标识
func (r *AlitripAgentFlightSellTicketingDetailAPIRequest) SetDomesticIntl(_domesticIntl int64) error {
	r._domesticIntl = _domesticIntl
	r.Set("domestic_intl", _domesticIntl)
	return nil
}

// GetDomesticIntl DomesticIntl Getter
func (r AlitripAgentFlightSellTicketingDetailAPIRequest) GetDomesticIntl() int64 {
	return r._domesticIntl
}

var poolAlitripAgentFlightSellTicketingDetailAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripAgentFlightSellTicketingDetailRequest()
	},
}

// GetAlitripAgentFlightSellTicketingDetailRequest 从 sync.Pool 获取 AlitripAgentFlightSellTicketingDetailAPIRequest
func GetAlitripAgentFlightSellTicketingDetailAPIRequest() *AlitripAgentFlightSellTicketingDetailAPIRequest {
	return poolAlitripAgentFlightSellTicketingDetailAPIRequest.Get().(*AlitripAgentFlightSellTicketingDetailAPIRequest)
}

// ReleaseAlitripAgentFlightSellTicketingDetailAPIRequest 将 AlitripAgentFlightSellTicketingDetailAPIRequest 放入 sync.Pool
func ReleaseAlitripAgentFlightSellTicketingDetailAPIRequest(v *AlitripAgentFlightSellTicketingDetailAPIRequest) {
	v.Reset()
	poolAlitripAgentFlightSellTicketingDetailAPIRequest.Put(v)
}
