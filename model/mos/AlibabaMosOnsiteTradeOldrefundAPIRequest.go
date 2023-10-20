package mos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosOnsiteTradeOldrefundAPIRequest 线下新退款接口（专为老退款接口调用） API请求
// alibaba.mos.onsite.trade.oldrefund
//
// 线下新退款接口（专为老退款接口调用）。新接口支付支付宝、手淘、天猫，老接口退款时，需要调用该接口退新单，并适配老接口响应参数返回
type AlibabaMosOnsiteTradeOldrefundAPIRequest struct {
	model.Params
	// 交易退款请求
	_onsiteRefundRequest *OnsiteRefundRequest
}

// NewAlibabaMosOnsiteTradeOldrefundRequest 初始化AlibabaMosOnsiteTradeOldrefundAPIRequest对象
func NewAlibabaMosOnsiteTradeOldrefundRequest() *AlibabaMosOnsiteTradeOldrefundAPIRequest {
	return &AlibabaMosOnsiteTradeOldrefundAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMosOnsiteTradeOldrefundAPIRequest) Reset() {
	r._onsiteRefundRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosOnsiteTradeOldrefundAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.onsite.trade.oldrefund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosOnsiteTradeOldrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosOnsiteTradeOldrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOnsiteRefundRequest is OnsiteRefundRequest Setter
// 交易退款请求
func (r *AlibabaMosOnsiteTradeOldrefundAPIRequest) SetOnsiteRefundRequest(_onsiteRefundRequest *OnsiteRefundRequest) error {
	r._onsiteRefundRequest = _onsiteRefundRequest
	r.Set("onsite_refund_request", _onsiteRefundRequest)
	return nil
}

// GetOnsiteRefundRequest OnsiteRefundRequest Getter
func (r AlibabaMosOnsiteTradeOldrefundAPIRequest) GetOnsiteRefundRequest() *OnsiteRefundRequest {
	return r._onsiteRefundRequest
}

var poolAlibabaMosOnsiteTradeOldrefundAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMosOnsiteTradeOldrefundRequest()
	},
}

// GetAlibabaMosOnsiteTradeOldrefundRequest 从 sync.Pool 获取 AlibabaMosOnsiteTradeOldrefundAPIRequest
func GetAlibabaMosOnsiteTradeOldrefundAPIRequest() *AlibabaMosOnsiteTradeOldrefundAPIRequest {
	return poolAlibabaMosOnsiteTradeOldrefundAPIRequest.Get().(*AlibabaMosOnsiteTradeOldrefundAPIRequest)
}

// ReleaseAlibabaMosOnsiteTradeOldrefundAPIRequest 将 AlibabaMosOnsiteTradeOldrefundAPIRequest 放入 sync.Pool
func ReleaseAlibabaMosOnsiteTradeOldrefundAPIRequest(v *AlibabaMosOnsiteTradeOldrefundAPIRequest) {
	v.Reset()
	poolAlibabaMosOnsiteTradeOldrefundAPIRequest.Put(v)
}
