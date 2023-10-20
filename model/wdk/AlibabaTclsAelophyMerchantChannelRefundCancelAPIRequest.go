package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest 翱象商家自有渠道 逆向单申请取消 API请求
// alibaba.tcls.aelophy.merchant.channel.refund.cancel
//
// 翱象小程序 用户逆向申请取消
type AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest struct {
	model.Params
	// 逆向申请取消
	_refundCancelInfo *RefundCancelInfo
}

// NewAlibabaTclsAelophyMerchantChannelRefundCancelRequest 初始化AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest对象
func NewAlibabaTclsAelophyMerchantChannelRefundCancelRequest() *AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest {
	return &AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest) Reset() {
	r._refundCancelInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.channel.refund.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundCancelInfo is RefundCancelInfo Setter
// 逆向申请取消
func (r *AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest) SetRefundCancelInfo(_refundCancelInfo *RefundCancelInfo) error {
	r._refundCancelInfo = _refundCancelInfo
	r.Set("refund_cancel_info", _refundCancelInfo)
	return nil
}

// GetRefundCancelInfo RefundCancelInfo Getter
func (r AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest) GetRefundCancelInfo() *RefundCancelInfo {
	return r._refundCancelInfo
}

var poolAlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTclsAelophyMerchantChannelRefundCancelRequest()
	},
}

// GetAlibabaTclsAelophyMerchantChannelRefundCancelRequest 从 sync.Pool 获取 AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest
func GetAlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest() *AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest {
	return poolAlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest.Get().(*AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest)
}

// ReleaseAlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest 将 AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest 放入 sync.Pool
func ReleaseAlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest(v *AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest) {
	v.Reset()
	poolAlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest.Put(v)
}
