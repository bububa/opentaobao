package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
回传订单退款审核结果 API请求
alibaba.einvoice.order.refund.update

ISV回传订单退款审核结果
*/
type AlibabaEinvoiceOrderRefundUpdateRequest struct {
    model.Params
    // 退款审核结果DTO
    _orderRefundResultDto   *InvoiceOrderRefundResultDTO
}

// 初始化AlibabaEinvoiceOrderRefundUpdateRequest对象
func NewAlibabaEinvoiceOrderRefundUpdateRequest() *AlibabaEinvoiceOrderRefundUpdateRequest{
    return &AlibabaEinvoiceOrderRefundUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceOrderRefundUpdateRequest) GetApiMethodName() string {
    return "alibaba.einvoice.order.refund.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceOrderRefundUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderRefundResultDto Setter
// 退款审核结果DTO
func (r *AlibabaEinvoiceOrderRefundUpdateRequest) SetOrderRefundResultDto(_orderRefundResultDto *InvoiceOrderRefundResultDTO) error {
    r._orderRefundResultDto = _orderRefundResultDto
    r.Set("order_refund_result_dto", _orderRefundResultDto)
    return nil
}

// OrderRefundResultDto Getter
func (r AlibabaEinvoiceOrderRefundUpdateRequest) GetOrderRefundResultDto() *InvoiceOrderRefundResultDTO {
    return r._orderRefundResultDto
}
