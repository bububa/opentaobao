package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophymerchantchannelrefundcompleteAPIRequest 翱象商家自有渠道 逆向单完成 API请求
// alibaba.tcls.aelophy.merchant.channel.refund.complete
//
// 翱象小程序 退款完成
type AlibabatclsaelophymerchantchannelrefundcompleteAPIRequest struct {
	model.Params
	// 请求对象
	_refundCompleteInfo *RefundCompleteInfo
}

// NewAlibabatclsaelophymerchantchannelrefundcompleteRequest 初始化AlibabatclsaelophymerchantchannelrefundcompleteAPIRequest对象
func NewAlibabatclsaelophymerchantchannelrefundcompleteRequest() *AlibabatclsaelophymerchantchannelrefundcompleteAPIRequest {
	return &AlibabatclsaelophymerchantchannelrefundcompleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsaelophymerchantchannelrefundcompleteAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.channel.refund.complete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsaelophymerchantchannelrefundcompleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsaelophymerchantchannelrefundcompleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundCompleteInfo is RefundCompleteInfo Setter
// 请求对象
func (r *AlibabatclsaelophymerchantchannelrefundcompleteAPIRequest) SetRefundCompleteInfo(_refundCompleteInfo *RefundCompleteInfo) error {
	r._refundCompleteInfo = _refundCompleteInfo
	r.Set("refund_complete_info", _refundCompleteInfo)
	return nil
}

// GetRefundCompleteInfo RefundCompleteInfo Getter
func (r AlibabatclsaelophymerchantchannelrefundcompleteAPIRequest) GetRefundCompleteInfo() *RefundCompleteInfo {
	return r._refundCompleteInfo
}
