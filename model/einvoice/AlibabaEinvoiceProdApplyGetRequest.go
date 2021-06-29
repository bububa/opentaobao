package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询发票申请 APIRequest
alibaba.einvoice.prod.apply.get

查询申请的详细信息，包含申请所关联的发票摘要信息+板式文件+预览图；
场景使用：1、业务前台收到申请状态变更消息后，调用此接口查询申请详情；2、主动补偿查询：当指定了自动开票，且发票申请长时间未收到状态变更通知时，可能存在丢消息的情况，此时可主动查询该申请，然后更新本地工单状态。
*/
type AlibabaEinvoiceProdApplyGetRequest struct {
    model.Params

    // 查询申请请求
    invoiceApplyQueryDto   *InvoiceApplyDtlQueryDto 

}

func NewAlibabaEinvoiceProdApplyGetRequest() *AlibabaEinvoiceProdApplyGetRequest{
    return &AlibabaEinvoiceProdApplyGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceProdApplyGetRequest) GetApiMethodName() string {
    return "alibaba.einvoice.prod.apply.get"
}

func (r AlibabaEinvoiceProdApplyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceProdApplyGetRequest) SetInvoiceApplyQueryDto(invoiceApplyQueryDto *InvoiceApplyDtlQueryDto) error {
    r.invoiceApplyQueryDto = invoiceApplyQueryDto
    r.Set("invoice_apply_query_dto", invoiceApplyQueryDto)
    return nil
}

func (r AlibabaEinvoiceProdApplyGetRequest) GetInvoiceApplyQueryDto() *InvoiceApplyDtlQueryDto {
    return r.invoiceApplyQueryDto
}

