package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentflightsellticketingdetailAPIRequest 销售出票详情 API请求
// alitrip.agent.flight.sell.ticketing.detail
//
// 销售出票详情
type AlitripagentflightsellticketingdetailAPIRequest struct {
	model.Params
	// 飞猪订单号
	_orderId string
	// 国内国际标识
	_domesticIntl int64
}

// NewAlitripagentflightsellticketingdetailRequest 初始化AlitripagentflightsellticketingdetailAPIRequest对象
func NewAlitripagentflightsellticketingdetailRequest() *AlitripagentflightsellticketingdetailAPIRequest {
	return &AlitripagentflightsellticketingdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripagentflightsellticketingdetailAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.ticketing.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripagentflightsellticketingdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripagentflightsellticketingdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 飞猪订单号
func (r *AlitripagentflightsellticketingdetailAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitripagentflightsellticketingdetailAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetDomesticIntl is DomesticIntl Setter
// 国内国际标识
func (r *AlitripagentflightsellticketingdetailAPIRequest) SetDomesticIntl(_domesticIntl int64) error {
	r._domesticIntl = _domesticIntl
	r.Set("domestic_intl", _domesticIntl)
	return nil
}

// GetDomesticIntl DomesticIntl Getter
func (r AlitripagentflightsellticketingdetailAPIRequest) GetDomesticIntl() int64 {
	return r._domesticIntl
}
