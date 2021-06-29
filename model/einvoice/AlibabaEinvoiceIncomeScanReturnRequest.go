package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
进项扫描状态回传 API请求
alibaba.einvoice.income.scan.return

回传进项扫描每个阶段的状态，比如ocr开始，ocr结束，查验开始，查验结束等
*/
type AlibabaEinvoiceIncomeScanReturnRequest struct {
    model.Params
    // 扫描的批次号
    batchNo   string
    // 扫描状态，0=开始ocr，1=ocr结束，2=开始查验，3=查验结束
    status   int64
    // 该批次对应的发票数量，扫描结束和查验结束status=1，3时必填
    invoiceCount   int64
    // 驱动是否成功，true=成功，false=失败
    success   bool
    // 错误码，success=false时填入
    errorCode   string
    // 错误信息，success=false时必填
    errorMessage   string
}

// 初始化AlibabaEinvoiceIncomeScanReturnRequest对象
func NewAlibabaEinvoiceIncomeScanReturnRequest() *AlibabaEinvoiceIncomeScanReturnRequest{
    return &AlibabaEinvoiceIncomeScanReturnRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceIncomeScanReturnRequest) GetApiMethodName() string {
    return "alibaba.einvoice.income.scan.return"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceIncomeScanReturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BatchNo Setter
// 扫描的批次号
func (r *AlibabaEinvoiceIncomeScanReturnRequest) SetBatchNo(batchNo string) error {
    r.batchNo = batchNo
    r.Set("batch_no", batchNo)
    return nil
}

// BatchNo Getter
func (r AlibabaEinvoiceIncomeScanReturnRequest) GetBatchNo() string {
    return r.batchNo
}
// Status Setter
// 扫描状态，0=开始ocr，1=ocr结束，2=开始查验，3=查验结束
func (r *AlibabaEinvoiceIncomeScanReturnRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r AlibabaEinvoiceIncomeScanReturnRequest) GetStatus() int64 {
    return r.status
}
// InvoiceCount Setter
// 该批次对应的发票数量，扫描结束和查验结束status=1，3时必填
func (r *AlibabaEinvoiceIncomeScanReturnRequest) SetInvoiceCount(invoiceCount int64) error {
    r.invoiceCount = invoiceCount
    r.Set("invoice_count", invoiceCount)
    return nil
}

// InvoiceCount Getter
func (r AlibabaEinvoiceIncomeScanReturnRequest) GetInvoiceCount() int64 {
    return r.invoiceCount
}
// Success Setter
// 驱动是否成功，true=成功，false=失败
func (r *AlibabaEinvoiceIncomeScanReturnRequest) SetSuccess(success bool) error {
    r.success = success
    r.Set("success", success)
    return nil
}

// Success Getter
func (r AlibabaEinvoiceIncomeScanReturnRequest) GetSuccess() bool {
    return r.success
}
// ErrorCode Setter
// 错误码，success=false时填入
func (r *AlibabaEinvoiceIncomeScanReturnRequest) SetErrorCode(errorCode string) error {
    r.errorCode = errorCode
    r.Set("error_code", errorCode)
    return nil
}

// ErrorCode Getter
func (r AlibabaEinvoiceIncomeScanReturnRequest) GetErrorCode() string {
    return r.errorCode
}
// ErrorMessage Setter
// 错误信息，success=false时必填
func (r *AlibabaEinvoiceIncomeScanReturnRequest) SetErrorMessage(errorMessage string) error {
    r.errorMessage = errorMessage
    r.Set("error_message", errorMessage)
    return nil
}

// ErrorMessage Getter
func (r AlibabaEinvoiceIncomeScanReturnRequest) GetErrorMessage() string {
    return r.errorMessage
}
