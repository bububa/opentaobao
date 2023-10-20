package flight

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightSellModifyDetailAPIRequest 销售改签详情 API请求
// alitrip.agent.flight.sell.modify.detail
//
// 销售改签详情
type AlitripAgentFlightSellModifyDetailAPIRequest struct {
	model.Params
	// 申请单号
	_applyId string
	// 国内国际标识:1:国内,2:国际
	_domesticIntl int64
}

// NewAlitripAgentFlightSellModifyDetailRequest 初始化AlitripAgentFlightSellModifyDetailAPIRequest对象
func NewAlitripAgentFlightSellModifyDetailRequest() *AlitripAgentFlightSellModifyDetailAPIRequest {
	return &AlitripAgentFlightSellModifyDetailAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripAgentFlightSellModifyDetailAPIRequest) Reset() {
	r._applyId = ""
	r._domesticIntl = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellModifyDetailAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.modify.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellModifyDetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripAgentFlightSellModifyDetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyId is ApplyId Setter
// 申请单号
func (r *AlitripAgentFlightSellModifyDetailAPIRequest) SetApplyId(_applyId string) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r AlitripAgentFlightSellModifyDetailAPIRequest) GetApplyId() string {
	return r._applyId
}

// SetDomesticIntl is DomesticIntl Setter
// 国内国际标识:1:国内,2:国际
func (r *AlitripAgentFlightSellModifyDetailAPIRequest) SetDomesticIntl(_domesticIntl int64) error {
	r._domesticIntl = _domesticIntl
	r.Set("domestic_intl", _domesticIntl)
	return nil
}

// GetDomesticIntl DomesticIntl Getter
func (r AlitripAgentFlightSellModifyDetailAPIRequest) GetDomesticIntl() int64 {
	return r._domesticIntl
}

var poolAlitripAgentFlightSellModifyDetailAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripAgentFlightSellModifyDetailRequest()
	},
}

// GetAlitripAgentFlightSellModifyDetailRequest 从 sync.Pool 获取 AlitripAgentFlightSellModifyDetailAPIRequest
func GetAlitripAgentFlightSellModifyDetailAPIRequest() *AlitripAgentFlightSellModifyDetailAPIRequest {
	return poolAlitripAgentFlightSellModifyDetailAPIRequest.Get().(*AlitripAgentFlightSellModifyDetailAPIRequest)
}

// ReleaseAlitripAgentFlightSellModifyDetailAPIRequest 将 AlitripAgentFlightSellModifyDetailAPIRequest 放入 sync.Pool
func ReleaseAlitripAgentFlightSellModifyDetailAPIRequest(v *AlitripAgentFlightSellModifyDetailAPIRequest) {
	v.Reset()
	poolAlitripAgentFlightSellModifyDetailAPIRequest.Put(v)
}
