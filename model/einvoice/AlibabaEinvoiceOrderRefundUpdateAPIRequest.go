package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceorderrefundupdateAPIRequest 回传订单退款审核结果 API请求
// alibaba.einvoice.order.refund.update
//
// ISV回传订单退款审核结果
type AlibabaeinvoiceorderrefundupdateAPIRequest struct {
	model.Params
	// 退款审核结果DTO
	_orderRefundResultDto *InvoiceOrderRefundResultDto
}

// NewAlibabaeinvoiceorderrefundupdateRequest 初始化AlibabaeinvoiceorderrefundupdateAPIRequest对象
func NewAlibabaeinvoiceorderrefundupdateRequest() *AlibabaeinvoiceorderrefundupdateAPIRequest {
	return &AlibabaeinvoiceorderrefundupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoiceorderrefundupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.order.refund.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoiceorderrefundupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoiceorderrefundupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderRefundResultDto is OrderRefundResultDto Setter
// 退款审核结果DTO
func (r *AlibabaeinvoiceorderrefundupdateAPIRequest) SetOrderRefundResultDto(_orderRefundResultDto *InvoiceOrderRefundResultDto) error {
	r._orderRefundResultDto = _orderRefundResultDto
	r.Set("order_refund_result_dto", _orderRefundResultDto)
	return nil
}

// GetOrderRefundResultDto OrderRefundResultDto Getter
func (r AlibabaeinvoiceorderrefundupdateAPIRequest) GetOrderRefundResultDto() *InvoiceOrderRefundResultDto {
	return r._orderRefundResultDto
}
