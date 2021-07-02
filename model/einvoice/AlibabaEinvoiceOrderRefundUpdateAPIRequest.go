package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceOrderRefundUpdateAPIRequest 回传订单退款审核结果 API请求
// alibaba.einvoice.order.refund.update
//
// ISV回传订单退款审核结果
type AlibabaEinvoiceOrderRefundUpdateAPIRequest struct {
	model.Params
	// 退款审核结果DTO
	_orderRefundResultDto *InvoiceOrderRefundResultDto
}

// NewAlibabaEinvoiceOrderRefundUpdateRequest 初始化AlibabaEinvoiceOrderRefundUpdateAPIRequest对象
func NewAlibabaEinvoiceOrderRefundUpdateRequest() *AlibabaEinvoiceOrderRefundUpdateAPIRequest {
	return &AlibabaEinvoiceOrderRefundUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceOrderRefundUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.order.refund.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceOrderRefundUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderRefundResultDto is OrderRefundResultDto Setter
// 退款审核结果DTO
func (r *AlibabaEinvoiceOrderRefundUpdateAPIRequest) SetOrderRefundResultDto(_orderRefundResultDto *InvoiceOrderRefundResultDto) error {
	r._orderRefundResultDto = _orderRefundResultDto
	r.Set("order_refund_result_dto", _orderRefundResultDto)
	return nil
}

// GetOrderRefundResultDto OrderRefundResultDto Getter
func (r AlibabaEinvoiceOrderRefundUpdateAPIRequest) GetOrderRefundResultDto() *InvoiceOrderRefundResultDto {
	return r._orderRefundResultDto
}
