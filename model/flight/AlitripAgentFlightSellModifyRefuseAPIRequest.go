package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightSellModifyRefuseAPIRequest 销售改签拒绝 API请求
// alitrip.agent.flight.sell.modify.refuse
//
// 销售改签拒绝
type AlitripAgentFlightSellModifyRefuseAPIRequest struct {
	model.Params
	// 申请单号
	_applyId string
	// 拒绝原因
	_refuseReason string
	// 国际国内标识
	_domesticIntl int64
}

// NewAlitripAgentFlightSellModifyRefuseRequest 初始化AlitripAgentFlightSellModifyRefuseAPIRequest对象
func NewAlitripAgentFlightSellModifyRefuseRequest() *AlitripAgentFlightSellModifyRefuseAPIRequest {
	return &AlitripAgentFlightSellModifyRefuseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellModifyRefuseAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.modify.refuse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellModifyRefuseAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetApplyId is ApplyId Setter
// 申请单号
func (r *AlitripAgentFlightSellModifyRefuseAPIRequest) SetApplyId(_applyId string) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r AlitripAgentFlightSellModifyRefuseAPIRequest) GetApplyId() string {
	return r._applyId
}

// SetRefuseReason is RefuseReason Setter
// 拒绝原因
func (r *AlitripAgentFlightSellModifyRefuseAPIRequest) SetRefuseReason(_refuseReason string) error {
	r._refuseReason = _refuseReason
	r.Set("refuse_reason", _refuseReason)
	return nil
}

// GetRefuseReason RefuseReason Getter
func (r AlitripAgentFlightSellModifyRefuseAPIRequest) GetRefuseReason() string {
	return r._refuseReason
}

// SetDomesticIntl is DomesticIntl Setter
// 国际国内标识
func (r *AlitripAgentFlightSellModifyRefuseAPIRequest) SetDomesticIntl(_domesticIntl int64) error {
	r._domesticIntl = _domesticIntl
	r.Set("domestic_intl", _domesticIntl)
	return nil
}

// GetDomesticIntl DomesticIntl Getter
func (r AlitripAgentFlightSellModifyRefuseAPIRequest) GetDomesticIntl() int64 {
	return r._domesticIntl
}
