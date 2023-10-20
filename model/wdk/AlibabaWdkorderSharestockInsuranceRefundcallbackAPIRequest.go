package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkorderSharestockInsuranceRefundcallbackAPIRequest 共享库存逆向订单理赔单回传 API请求
// alibaba.wdkorder.sharestock.insurance.refundcallback
//
// 共享库存逆向订单理赔单回传
type AlibabaWdkorderSharestockInsuranceRefundcallbackAPIRequest struct {
	model.Params
	// 退款单ID
	_refundId string
	// 理赔单ID
	_claimId string
	// 淘宝交易子单ID
	_tbSubOrderId int64
}

// NewAlibabaWdkorderSharestockInsuranceRefundcallbackRequest 初始化AlibabaWdkorderSharestockInsuranceRefundcallbackAPIRequest对象
func NewAlibabaWdkorderSharestockInsuranceRefundcallbackRequest() *AlibabaWdkorderSharestockInsuranceRefundcallbackAPIRequest {
	return &AlibabaWdkorderSharestockInsuranceRefundcallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkorderSharestockInsuranceRefundcallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.wdkorder.sharestock.insurance.refundcallback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkorderSharestockInsuranceRefundcallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkorderSharestockInsuranceRefundcallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundId is RefundId Setter
// 退款单ID
func (r *AlibabaWdkorderSharestockInsuranceRefundcallbackAPIRequest) SetRefundId(_refundId string) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r AlibabaWdkorderSharestockInsuranceRefundcallbackAPIRequest) GetRefundId() string {
	return r._refundId
}

// SetClaimId is ClaimId Setter
// 理赔单ID
func (r *AlibabaWdkorderSharestockInsuranceRefundcallbackAPIRequest) SetClaimId(_claimId string) error {
	r._claimId = _claimId
	r.Set("claim_id", _claimId)
	return nil
}

// GetClaimId ClaimId Getter
func (r AlibabaWdkorderSharestockInsuranceRefundcallbackAPIRequest) GetClaimId() string {
	return r._claimId
}

// SetTbSubOrderId is TbSubOrderId Setter
// 淘宝交易子单ID
func (r *AlibabaWdkorderSharestockInsuranceRefundcallbackAPIRequest) SetTbSubOrderId(_tbSubOrderId int64) error {
	r._tbSubOrderId = _tbSubOrderId
	r.Set("tb_sub_order_id", _tbSubOrderId)
	return nil
}

// GetTbSubOrderId TbSubOrderId Getter
func (r AlibabaWdkorderSharestockInsuranceRefundcallbackAPIRequest) GetTbSubOrderId() int64 {
	return r._tbSubOrderId
}
