package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripAgentFlightSellModifyDetailAPIRequest
销售改签详情 API请求
alitrip.agent.flight.sell.modify.detail

销售改签详情 */
type AlitripAgentFlightSellModifyDetailAPIRequest struct {
	model.Params
	// 申请单号
	_applyId string
	// 国内国际标识
	_domesticIntl int64
}

// NewAlitripAgentFlightSellModifyDetailRequest 初始化AlitripAgentFlightSellModifyDetailAPIRequest对象
func NewAlitripAgentFlightSellModifyDetailRequest() *AlitripAgentFlightSellModifyDetailAPIRequest {
	return &AlitripAgentFlightSellModifyDetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellModifyDetailAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.modify.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellModifyDetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ApplyId Setter
// 申请单号
func (r *AlitripAgentFlightSellModifyDetailAPIRequest) SetApplyId(_applyId string) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// Get ApplyId Getter
func (r AlitripAgentFlightSellModifyDetailAPIRequest) GetApplyId() string {
	return r._applyId
}

// Set is DomesticIntl Setter
// 国内国际标识
func (r *AlitripAgentFlightSellModifyDetailAPIRequest) SetDomesticIntl(_domesticIntl int64) error {
	r._domesticIntl = _domesticIntl
	r.Set("domestic_intl", _domesticIntl)
	return nil
}

// Get DomesticIntl Getter
func (r AlitripAgentFlightSellModifyDetailAPIRequest) GetDomesticIntl() int64 {
	return r._domesticIntl
}
