package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest
翱象商家自有渠道 逆向单申请 API请求
alibaba.tcls.aelophy.merchant.channel.refund.apply

翱象小程序 用户逆向单申请 */
type AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest struct {
	model.Params
	// 请求对象
	_refundApplyInfo *RefundApplyInfo
}

// NewAlibabaTclsAelophyMerchantChannelRefundApplyRequest 初始化AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest对象
func NewAlibabaTclsAelophyMerchantChannelRefundApplyRequest() *AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest {
	return &AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.channel.refund.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefundApplyInfo Setter
// 请求对象
func (r *AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest) SetRefundApplyInfo(_refundApplyInfo *RefundApplyInfo) error {
	r._refundApplyInfo = _refundApplyInfo
	r.Set("refund_apply_info", _refundApplyInfo)
	return nil
}

// Get RefundApplyInfo Getter
func (r AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest) GetRefundApplyInfo() *RefundApplyInfo {
	return r._refundApplyInfo
}
