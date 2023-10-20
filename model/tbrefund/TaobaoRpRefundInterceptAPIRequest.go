package tbrefund

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRpRefundInterceptAPIRequest 卖家发起拦截 API请求
// taobao.rp.refund.intercept
//
// 卖家发起拦截
type TaobaoRpRefundInterceptAPIRequest struct {
	model.Params
	// 退款编号
	_refundId int64
	// 退款版本号
	_refundVersion int64
}

// NewTaobaoRpRefundInterceptRequest 初始化TaobaoRpRefundInterceptAPIRequest对象
func NewTaobaoRpRefundInterceptRequest() *TaobaoRpRefundInterceptAPIRequest {
	return &TaobaoRpRefundInterceptAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRpRefundInterceptAPIRequest) Reset() {
	r._refundId = 0
	r._refundVersion = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRpRefundInterceptAPIRequest) GetApiMethodName() string {
	return "taobao.rp.refund.intercept"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRpRefundInterceptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRpRefundInterceptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundId is RefundId Setter
// 退款编号
func (r *TaobaoRpRefundInterceptAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaoRpRefundInterceptAPIRequest) GetRefundId() int64 {
	return r._refundId
}

// SetRefundVersion is RefundVersion Setter
// 退款版本号
func (r *TaobaoRpRefundInterceptAPIRequest) SetRefundVersion(_refundVersion int64) error {
	r._refundVersion = _refundVersion
	r.Set("refund_version", _refundVersion)
	return nil
}

// GetRefundVersion RefundVersion Getter
func (r TaobaoRpRefundInterceptAPIRequest) GetRefundVersion() int64 {
	return r._refundVersion
}

var poolTaobaoRpRefundInterceptAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRpRefundInterceptRequest()
	},
}

// GetTaobaoRpRefundInterceptRequest 从 sync.Pool 获取 TaobaoRpRefundInterceptAPIRequest
func GetTaobaoRpRefundInterceptAPIRequest() *TaobaoRpRefundInterceptAPIRequest {
	return poolTaobaoRpRefundInterceptAPIRequest.Get().(*TaobaoRpRefundInterceptAPIRequest)
}

// ReleaseTaobaoRpRefundInterceptAPIRequest 将 TaobaoRpRefundInterceptAPIRequest 放入 sync.Pool
func ReleaseTaobaoRpRefundInterceptAPIRequest(v *TaobaoRpRefundInterceptAPIRequest) {
	v.Reset()
	poolTaobaoRpRefundInterceptAPIRequest.Put(v)
}
