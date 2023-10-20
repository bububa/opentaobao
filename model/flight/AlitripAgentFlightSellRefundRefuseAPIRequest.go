package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentflightsellrefundrefuseAPIRequest 销售退票拒绝 API请求
// alitrip.agent.flight.sell.refund.refuse
//
// 销售退票拒绝
type AlitripagentflightsellrefundrefuseAPIRequest struct {
	model.Params
	// 申请单号
	_applyId string
	// 拒绝原因
	_refuseReason string
	// 国内国际标识:1:国内,2:国际
	_domesticIntl int64
}

// NewAlitripagentflightsellrefundrefuseRequest 初始化AlitripagentflightsellrefundrefuseAPIRequest对象
func NewAlitripagentflightsellrefundrefuseRequest() *AlitripagentflightsellrefundrefuseAPIRequest {
	return &AlitripagentflightsellrefundrefuseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripagentflightsellrefundrefuseAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.refund.refuse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripagentflightsellrefundrefuseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripagentflightsellrefundrefuseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyId is ApplyId Setter
// 申请单号
func (r *AlitripagentflightsellrefundrefuseAPIRequest) SetApplyId(_applyId string) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r AlitripagentflightsellrefundrefuseAPIRequest) GetApplyId() string {
	return r._applyId
}

// SetRefuseReason is RefuseReason Setter
// 拒绝原因
func (r *AlitripagentflightsellrefundrefuseAPIRequest) SetRefuseReason(_refuseReason string) error {
	r._refuseReason = _refuseReason
	r.Set("refuse_reason", _refuseReason)
	return nil
}

// GetRefuseReason RefuseReason Getter
func (r AlitripagentflightsellrefundrefuseAPIRequest) GetRefuseReason() string {
	return r._refuseReason
}

// SetDomesticIntl is DomesticIntl Setter
// 国内国际标识:1:国内,2:国际
func (r *AlitripagentflightsellrefundrefuseAPIRequest) SetDomesticIntl(_domesticIntl int64) error {
	r._domesticIntl = _domesticIntl
	r.Set("domestic_intl", _domesticIntl)
	return nil
}

// GetDomesticIntl DomesticIntl Getter
func (r AlitripagentflightsellrefundrefuseAPIRequest) GetDomesticIntl() int64 {
	return r._domesticIntl
}
