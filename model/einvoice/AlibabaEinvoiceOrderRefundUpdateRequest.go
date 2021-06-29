package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
回传订单退款审核结果 APIRequest
alibaba.einvoice.order.refund.update

ISV回传订单退款审核结果
*/
type AlibabaEinvoiceOrderRefundUpdateRequest struct {
    model.Params

    // 退款审核结果DTO
    orderRefundResultDto   *InvoiceOrderRefundResultDto 

}

func NewAlibabaEinvoiceOrderRefundUpdateRequest() *AlibabaEinvoiceOrderRefundUpdateRequest{
    return &AlibabaEinvoiceOrderRefundUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceOrderRefundUpdateRequest) GetApiMethodName() string {
    return "alibaba.einvoice.order.refund.update"
}

func (r AlibabaEinvoiceOrderRefundUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceOrderRefundUpdateRequest) SetOrderRefundResultDto(orderRefundResultDto *InvoiceOrderRefundResultDto) error {
    r.orderRefundResultDto = orderRefundResultDto
    r.Set("order_refund_result_dto", orderRefundResultDto)
    return nil
}

func (r AlibabaEinvoiceOrderRefundUpdateRequest) GetOrderRefundResultDto() *InvoiceOrderRefundResultDto {
    return r.orderRefundResultDto
}

