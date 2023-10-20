package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentflightsellmodifyrefuseAPIRequest 销售改签拒绝 API请求
// alitrip.agent.flight.sell.modify.refuse
//
// 销售改签拒绝
type AlitripagentflightsellmodifyrefuseAPIRequest struct {
	model.Params
	// 申请单号
	_applyId string
	// 拒绝原因
	_refuseReason string
	// 国际国内标识:1:国内,2:国际
	_domesticIntl int64
}

// NewAlitripagentflightsellmodifyrefuseRequest 初始化AlitripagentflightsellmodifyrefuseAPIRequest对象
func NewAlitripagentflightsellmodifyrefuseRequest() *AlitripagentflightsellmodifyrefuseAPIRequest {
	return &AlitripagentflightsellmodifyrefuseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripagentflightsellmodifyrefuseAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.modify.refuse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripagentflightsellmodifyrefuseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripagentflightsellmodifyrefuseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyId is ApplyId Setter
// 申请单号
func (r *AlitripagentflightsellmodifyrefuseAPIRequest) SetApplyId(_applyId string) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r AlitripagentflightsellmodifyrefuseAPIRequest) GetApplyId() string {
	return r._applyId
}

// SetRefuseReason is RefuseReason Setter
// 拒绝原因
func (r *AlitripagentflightsellmodifyrefuseAPIRequest) SetRefuseReason(_refuseReason string) error {
	r._refuseReason = _refuseReason
	r.Set("refuse_reason", _refuseReason)
	return nil
}

// GetRefuseReason RefuseReason Getter
func (r AlitripagentflightsellmodifyrefuseAPIRequest) GetRefuseReason() string {
	return r._refuseReason
}

// SetDomesticIntl is DomesticIntl Setter
// 国际国内标识:1:国内,2:国际
func (r *AlitripagentflightsellmodifyrefuseAPIRequest) SetDomesticIntl(_domesticIntl int64) error {
	r._domesticIntl = _domesticIntl
	r.Set("domestic_intl", _domesticIntl)
	return nil
}

// GetDomesticIntl DomesticIntl Getter
func (r AlitripagentflightsellmodifyrefuseAPIRequest) GetDomesticIntl() int64 {
	return r._domesticIntl
}
