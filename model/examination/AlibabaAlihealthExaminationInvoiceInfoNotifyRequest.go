package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
体检机构同步发票信息给阿里健康 API请求
alibaba.alihealth.examination.invoice.info.notify

体检机构向阿里健康同步发票信息
*/
type AlibabaAlihealthExaminationInvoiceInfoNotifyRequest struct {
    model.Params
    // 开票状态；（have_submit已提交、invoice_done已开票）
    invoiceStatus   string
    // 发票访问地址；（invoice_status在已开票状态下必填）
    invoiceUrl   string
    // 阿里健康预约凭证
    reserveNumber   string
}

// 初始化AlibabaAlihealthExaminationInvoiceInfoNotifyRequest对象
func NewAlibabaAlihealthExaminationInvoiceInfoNotifyRequest() *AlibabaAlihealthExaminationInvoiceInfoNotifyRequest{
    return &AlibabaAlihealthExaminationInvoiceInfoNotifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationInvoiceInfoNotifyRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.invoice.info.notify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationInvoiceInfoNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InvoiceStatus Setter
// 开票状态；（have_submit已提交、invoice_done已开票）
func (r *AlibabaAlihealthExaminationInvoiceInfoNotifyRequest) SetInvoiceStatus(invoiceStatus string) error {
    r.invoiceStatus = invoiceStatus
    r.Set("invoice_status", invoiceStatus)
    return nil
}

// InvoiceStatus Getter
func (r AlibabaAlihealthExaminationInvoiceInfoNotifyRequest) GetInvoiceStatus() string {
    return r.invoiceStatus
}
// InvoiceUrl Setter
// 发票访问地址；（invoice_status在已开票状态下必填）
func (r *AlibabaAlihealthExaminationInvoiceInfoNotifyRequest) SetInvoiceUrl(invoiceUrl string) error {
    r.invoiceUrl = invoiceUrl
    r.Set("invoice_url", invoiceUrl)
    return nil
}

// InvoiceUrl Getter
func (r AlibabaAlihealthExaminationInvoiceInfoNotifyRequest) GetInvoiceUrl() string {
    return r.invoiceUrl
}
// ReserveNumber Setter
// 阿里健康预约凭证
func (r *AlibabaAlihealthExaminationInvoiceInfoNotifyRequest) SetReserveNumber(reserveNumber string) error {
    r.reserveNumber = reserveNumber
    r.Set("reserve_number", reserveNumber)
    return nil
}

// ReserveNumber Getter
func (r AlibabaAlihealthExaminationInvoiceInfoNotifyRequest) GetReserveNumber() string {
    return r.reserveNumber
}
