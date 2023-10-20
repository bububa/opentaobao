package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentflightsellrefunddetailAPIRequest 销售退票单详情 API请求
// alitrip.agent.flight.sell.refund.detail
//
// 销售退票单详情
type AlitripagentflightsellrefunddetailAPIRequest struct {
	model.Params
	// 申请单号
	_applyId string
	// 国际国内标识(1 国内,2 国际)
	_domesticIntl int64
}

// NewAlitripagentflightsellrefunddetailRequest 初始化AlitripagentflightsellrefunddetailAPIRequest对象
func NewAlitripagentflightsellrefunddetailRequest() *AlitripagentflightsellrefunddetailAPIRequest {
	return &AlitripagentflightsellrefunddetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripagentflightsellrefunddetailAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.refund.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripagentflightsellrefunddetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripagentflightsellrefunddetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyId is ApplyId Setter
// 申请单号
func (r *AlitripagentflightsellrefunddetailAPIRequest) SetApplyId(_applyId string) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r AlitripagentflightsellrefunddetailAPIRequest) GetApplyId() string {
	return r._applyId
}

// SetDomesticIntl is DomesticIntl Setter
// 国际国内标识(1 国内,2 国际)
func (r *AlitripagentflightsellrefunddetailAPIRequest) SetDomesticIntl(_domesticIntl int64) error {
	r._domesticIntl = _domesticIntl
	r.Set("domestic_intl", _domesticIntl)
	return nil
}

// GetDomesticIntl DomesticIntl Getter
func (r AlitripagentflightsellrefunddetailAPIRequest) GetDomesticIntl() int64 {
	return r._domesticIntl
}
