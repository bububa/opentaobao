package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophymerchantchannelrefundapplyAPIRequest 翱象商家自有渠道 逆向单申请 API请求
// alibaba.tcls.aelophy.merchant.channel.refund.apply
//
// 翱象小程序 用户逆向单申请
type AlibabatclsaelophymerchantchannelrefundapplyAPIRequest struct {
	model.Params
	// 请求对象
	_refundApplyInfo *RefundApplyInfo
}

// NewAlibabatclsaelophymerchantchannelrefundapplyRequest 初始化AlibabatclsaelophymerchantchannelrefundapplyAPIRequest对象
func NewAlibabatclsaelophymerchantchannelrefundapplyRequest() *AlibabatclsaelophymerchantchannelrefundapplyAPIRequest {
	return &AlibabatclsaelophymerchantchannelrefundapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsaelophymerchantchannelrefundapplyAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.channel.refund.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsaelophymerchantchannelrefundapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsaelophymerchantchannelrefundapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundApplyInfo is RefundApplyInfo Setter
// 请求对象
func (r *AlibabatclsaelophymerchantchannelrefundapplyAPIRequest) SetRefundApplyInfo(_refundApplyInfo *RefundApplyInfo) error {
	r._refundApplyInfo = _refundApplyInfo
	r.Set("refund_apply_info", _refundApplyInfo)
	return nil
}

// GetRefundApplyInfo RefundApplyInfo Getter
func (r AlibabatclsaelophymerchantchannelrefundapplyAPIRequest) GetRefundApplyInfo() *RefundApplyInfo {
	return r._refundApplyInfo
}
