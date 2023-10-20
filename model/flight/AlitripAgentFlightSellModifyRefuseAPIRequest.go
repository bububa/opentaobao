package flight

import (
	"net/url"
	"sync"

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
	// 国际国内标识:1:国内,2:国际
	_domesticIntl int64
}

// NewAlitripAgentFlightSellModifyRefuseRequest 初始化AlitripAgentFlightSellModifyRefuseAPIRequest对象
func NewAlitripAgentFlightSellModifyRefuseRequest() *AlitripAgentFlightSellModifyRefuseAPIRequest {
	return &AlitripAgentFlightSellModifyRefuseAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripAgentFlightSellModifyRefuseAPIRequest) Reset() {
	r._applyId = ""
	r._refuseReason = ""
	r._domesticIntl = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellModifyRefuseAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.modify.refuse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellModifyRefuseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripAgentFlightSellModifyRefuseAPIRequest) GetRawParams() model.Params {
	return r.Params
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
// 国际国内标识:1:国内,2:国际
func (r *AlitripAgentFlightSellModifyRefuseAPIRequest) SetDomesticIntl(_domesticIntl int64) error {
	r._domesticIntl = _domesticIntl
	r.Set("domestic_intl", _domesticIntl)
	return nil
}

// GetDomesticIntl DomesticIntl Getter
func (r AlitripAgentFlightSellModifyRefuseAPIRequest) GetDomesticIntl() int64 {
	return r._domesticIntl
}

var poolAlitripAgentFlightSellModifyRefuseAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripAgentFlightSellModifyRefuseRequest()
	},
}

// GetAlitripAgentFlightSellModifyRefuseRequest 从 sync.Pool 获取 AlitripAgentFlightSellModifyRefuseAPIRequest
func GetAlitripAgentFlightSellModifyRefuseAPIRequest() *AlitripAgentFlightSellModifyRefuseAPIRequest {
	return poolAlitripAgentFlightSellModifyRefuseAPIRequest.Get().(*AlitripAgentFlightSellModifyRefuseAPIRequest)
}

// ReleaseAlitripAgentFlightSellModifyRefuseAPIRequest 将 AlitripAgentFlightSellModifyRefuseAPIRequest 放入 sync.Pool
func ReleaseAlitripAgentFlightSellModifyRefuseAPIRequest(v *AlitripAgentFlightSellModifyRefuseAPIRequest) {
	v.Reset()
	poolAlitripAgentFlightSellModifyRefuseAPIRequest.Put(v)
}
