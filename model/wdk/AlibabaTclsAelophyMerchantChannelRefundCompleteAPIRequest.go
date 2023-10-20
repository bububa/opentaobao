package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest 翱象商家自有渠道 逆向单完成 API请求
// alibaba.tcls.aelophy.merchant.channel.refund.complete
//
// 翱象小程序 退款完成
type AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest struct {
	model.Params
	// 请求对象
	_refundCompleteInfo *RefundCompleteInfo
}

// NewAlibabaTclsAelophyMerchantChannelRefundCompleteRequest 初始化AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest对象
func NewAlibabaTclsAelophyMerchantChannelRefundCompleteRequest() *AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest {
	return &AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest) Reset() {
	r._refundCompleteInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.channel.refund.complete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundCompleteInfo is RefundCompleteInfo Setter
// 请求对象
func (r *AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest) SetRefundCompleteInfo(_refundCompleteInfo *RefundCompleteInfo) error {
	r._refundCompleteInfo = _refundCompleteInfo
	r.Set("refund_complete_info", _refundCompleteInfo)
	return nil
}

// GetRefundCompleteInfo RefundCompleteInfo Getter
func (r AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest) GetRefundCompleteInfo() *RefundCompleteInfo {
	return r._refundCompleteInfo
}

var poolAlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTclsAelophyMerchantChannelRefundCompleteRequest()
	},
}

// GetAlibabaTclsAelophyMerchantChannelRefundCompleteRequest 从 sync.Pool 获取 AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest
func GetAlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest() *AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest {
	return poolAlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest.Get().(*AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest)
}

// ReleaseAlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest 将 AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest 放入 sync.Pool
func ReleaseAlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest(v *AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest) {
	v.Reset()
	poolAlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest.Put(v)
}
