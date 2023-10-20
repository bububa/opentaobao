package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest 翱象商家自有渠道 逆向单申请 API请求
// alibaba.tcls.aelophy.merchant.channel.refund.apply
//
// 翱象小程序 用户逆向单申请
type AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest struct {
	model.Params
	// 请求对象
	_refundApplyInfo *RefundApplyInfo
}

// NewAlibabaTclsAelophyMerchantChannelRefundApplyRequest 初始化AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest对象
func NewAlibabaTclsAelophyMerchantChannelRefundApplyRequest() *AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest {
	return &AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest) Reset() {
	r._refundApplyInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.channel.refund.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundApplyInfo is RefundApplyInfo Setter
// 请求对象
func (r *AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest) SetRefundApplyInfo(_refundApplyInfo *RefundApplyInfo) error {
	r._refundApplyInfo = _refundApplyInfo
	r.Set("refund_apply_info", _refundApplyInfo)
	return nil
}

// GetRefundApplyInfo RefundApplyInfo Getter
func (r AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest) GetRefundApplyInfo() *RefundApplyInfo {
	return r._refundApplyInfo
}

var poolAlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTclsAelophyMerchantChannelRefundApplyRequest()
	},
}

// GetAlibabaTclsAelophyMerchantChannelRefundApplyRequest 从 sync.Pool 获取 AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest
func GetAlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest() *AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest {
	return poolAlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest.Get().(*AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest)
}

// ReleaseAlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest 将 AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest 放入 sync.Pool
func ReleaseAlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest(v *AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest) {
	v.Reset()
	poolAlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest.Put(v)
}
