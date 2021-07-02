package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkorderSharestockInsuranceRefundgetAPIRequest 共享库存投保业务售后逆向订单数据获取 API请求
// alibaba.wdkorder.sharestock.insurance.refundget
//
// 共享库存投保业务售后逆向订单数据获取
type AlibabaWdkorderSharestockInsuranceRefundgetAPIRequest struct {
	model.Params
	// 淘宝子订单ID
	_tbSubOrderId string
	// 退货单ID
	_refundId string
}

// NewAlibabaWdkorderSharestockInsuranceRefundgetRequest 初始化AlibabaWdkorderSharestockInsuranceRefundgetAPIRequest对象
func NewAlibabaWdkorderSharestockInsuranceRefundgetRequest() *AlibabaWdkorderSharestockInsuranceRefundgetAPIRequest {
	return &AlibabaWdkorderSharestockInsuranceRefundgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkorderSharestockInsuranceRefundgetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdkorder.sharestock.insurance.refundget"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkorderSharestockInsuranceRefundgetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTbSubOrderId is TbSubOrderId Setter
// 淘宝子订单ID
func (r *AlibabaWdkorderSharestockInsuranceRefundgetAPIRequest) SetTbSubOrderId(_tbSubOrderId string) error {
	r._tbSubOrderId = _tbSubOrderId
	r.Set("tb_sub_order_id", _tbSubOrderId)
	return nil
}

// GetTbSubOrderId TbSubOrderId Getter
func (r AlibabaWdkorderSharestockInsuranceRefundgetAPIRequest) GetTbSubOrderId() string {
	return r._tbSubOrderId
}

// SetRefundId is RefundId Setter
// 退货单ID
func (r *AlibabaWdkorderSharestockInsuranceRefundgetAPIRequest) SetRefundId(_refundId string) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r AlibabaWdkorderSharestockInsuranceRefundgetAPIRequest) GetRefundId() string {
	return r._refundId
}
