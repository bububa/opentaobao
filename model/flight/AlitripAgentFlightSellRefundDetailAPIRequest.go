package flight

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightSellRefundDetailAPIRequest 销售退票单详情 API请求
// alitrip.agent.flight.sell.refund.detail
//
// 销售退票单详情
type AlitripAgentFlightSellRefundDetailAPIRequest struct {
	model.Params
	// 申请单号
	_applyId string
	// 国际国内标识(1 国内,2 国际)
	_domesticIntl int64
}

// NewAlitripAgentFlightSellRefundDetailRequest 初始化AlitripAgentFlightSellRefundDetailAPIRequest对象
func NewAlitripAgentFlightSellRefundDetailRequest() *AlitripAgentFlightSellRefundDetailAPIRequest {
	return &AlitripAgentFlightSellRefundDetailAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripAgentFlightSellRefundDetailAPIRequest) Reset() {
	r._applyId = ""
	r._domesticIntl = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellRefundDetailAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.refund.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellRefundDetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripAgentFlightSellRefundDetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyId is ApplyId Setter
// 申请单号
func (r *AlitripAgentFlightSellRefundDetailAPIRequest) SetApplyId(_applyId string) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r AlitripAgentFlightSellRefundDetailAPIRequest) GetApplyId() string {
	return r._applyId
}

// SetDomesticIntl is DomesticIntl Setter
// 国际国内标识(1 国内,2 国际)
func (r *AlitripAgentFlightSellRefundDetailAPIRequest) SetDomesticIntl(_domesticIntl int64) error {
	r._domesticIntl = _domesticIntl
	r.Set("domestic_intl", _domesticIntl)
	return nil
}

// GetDomesticIntl DomesticIntl Getter
func (r AlitripAgentFlightSellRefundDetailAPIRequest) GetDomesticIntl() int64 {
	return r._domesticIntl
}

var poolAlitripAgentFlightSellRefundDetailAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripAgentFlightSellRefundDetailRequest()
	},
}

// GetAlitripAgentFlightSellRefundDetailRequest 从 sync.Pool 获取 AlitripAgentFlightSellRefundDetailAPIRequest
func GetAlitripAgentFlightSellRefundDetailAPIRequest() *AlitripAgentFlightSellRefundDetailAPIRequest {
	return poolAlitripAgentFlightSellRefundDetailAPIRequest.Get().(*AlitripAgentFlightSellRefundDetailAPIRequest)
}

// ReleaseAlitripAgentFlightSellRefundDetailAPIRequest 将 AlitripAgentFlightSellRefundDetailAPIRequest 放入 sync.Pool
func ReleaseAlitripAgentFlightSellRefundDetailAPIRequest(v *AlitripAgentFlightSellRefundDetailAPIRequest) {
	v.Reset()
	poolAlitripAgentFlightSellRefundDetailAPIRequest.Put(v)
}
