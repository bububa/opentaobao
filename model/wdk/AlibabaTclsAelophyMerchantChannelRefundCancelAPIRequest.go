package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophymerchantchannelrefundcancelAPIRequest 翱象商家自有渠道 逆向单申请取消 API请求
// alibaba.tcls.aelophy.merchant.channel.refund.cancel
//
// 翱象小程序 用户逆向申请取消
type AlibabatclsaelophymerchantchannelrefundcancelAPIRequest struct {
	model.Params
	// 逆向申请取消
	_refundCancelInfo *RefundCancelInfo
}

// NewAlibabatclsaelophymerchantchannelrefundcancelRequest 初始化AlibabatclsaelophymerchantchannelrefundcancelAPIRequest对象
func NewAlibabatclsaelophymerchantchannelrefundcancelRequest() *AlibabatclsaelophymerchantchannelrefundcancelAPIRequest {
	return &AlibabatclsaelophymerchantchannelrefundcancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsaelophymerchantchannelrefundcancelAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.channel.refund.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsaelophymerchantchannelrefundcancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsaelophymerchantchannelrefundcancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundCancelInfo is RefundCancelInfo Setter
// 逆向申请取消
func (r *AlibabatclsaelophymerchantchannelrefundcancelAPIRequest) SetRefundCancelInfo(_refundCancelInfo *RefundCancelInfo) error {
	r._refundCancelInfo = _refundCancelInfo
	r.Set("refund_cancel_info", _refundCancelInfo)
	return nil
}

// GetRefundCancelInfo RefundCancelInfo Getter
func (r AlibabatclsaelophymerchantchannelrefundcancelAPIRequest) GetRefundCancelInfo() *RefundCancelInfo {
	return r._refundCancelInfo
}
