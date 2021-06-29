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
    _batchNo   string
    // 扫描状态，0=开始ocr，1=ocr结束，2=开始查验，3=查验结束
    _status   int64
    // 该批次对应的发票数量，扫描结束和查验结束status=1，3时必填
    _invoiceCount   int64
    // 驱动是否成功，true=成功，false=失败
    _success   bool
    // 错误码，success=false时填入
    _errorCode   string
    // 错误信息，success=false时必填
    _errorMessage   string
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
func (r *AlibabaEinvoiceIncomeScanReturnRequest) SetBatchNo(_batchNo string) error {
    r._batchNo = _batchNo
    r.Set("batch_no", _batchNo)
    return nil
}

// BatchNo Getter
func (r AlibabaEinvoiceIncomeScanReturnRequest) GetBatchNo() string {
    return r._batchNo
}
// Status Setter
// 扫描状态，0=开始ocr，1=ocr结束，2=开始查验，3=查验结束
func (r *AlibabaEinvoiceIncomeScanReturnRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaEinvoiceIncomeScanReturnRequest) GetStatus() int64 {
    return r._status
}
// InvoiceCount Setter
// 该批次对应的发票数量，扫描结束和查验结束status=1，3时必填
func (r *AlibabaEinvoiceIncomeScanReturnRequest) SetInvoiceCount(_invoiceCount int64) error {
    r._invoiceCount = _invoiceCount
    r.Set("invoice_count", _invoiceCount)
    return nil
}

// InvoiceCount Getter
func (r AlibabaEinvoiceIncomeScanReturnRequest) GetInvoiceCount() int64 {
    return r._invoiceCount
}
// Success Setter
// 驱动是否成功，true=成功，false=失败
func (r *AlibabaEinvoiceIncomeScanReturnRequest) SetSuccess(_success bool) error {
    r._success = _success
    r.Set("success", _success)
    return nil
}

// Success Getter
func (r AlibabaEinvoiceIncomeScanReturnRequest) GetSuccess() bool {
    return r._success
}
// ErrorCode Setter
// 错误码，success=false时填入
func (r *AlibabaEinvoiceIncomeScanReturnRequest) SetErrorCode(_errorCode string) error {
    r._errorCode = _errorCode
    r.Set("error_code", _errorCode)
    return nil
}

// ErrorCode Getter
func (r AlibabaEinvoiceIncomeScanReturnRequest) GetErrorCode() string {
    return r._errorCode
}
// ErrorMessage Setter
// 错误信息，success=false时必填
func (r *AlibabaEinvoiceIncomeScanReturnRequest) SetErrorMessage(_errorMessage string) error {
    r._errorMessage = _errorMessage
    r.Set("error_message", _errorMessage)
    return nil
}

// ErrorMessage Getter
func (r AlibabaEinvoiceIncomeScanReturnRequest) GetErrorMessage() string {
    return r._errorMessage
}
