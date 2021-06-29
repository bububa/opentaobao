package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商回传发票ocr的结果 API请求
alibaba.einvoice.income.ocr.return

服务商回传发票ocr的结果，分两种场景：扫描驱动服务商主动回传；阿里主动发起的ocr回传
*/
type AlibabaEinvoiceIncomeOcrReturnRequest struct {
    model.Params
    // 校验码，ocr结果为普票，success=true并且invoiceKind=1时必填
    _checksum   string
    // 错误码，success=false是必填
    _errorCode   string
    // 错误消息，success=false是必填
    _errorMessage   string
    // 发票ocr影像文件，type=1时必填
    _imageData   *model.File
    // 发票ocr影像编号，type=1时必填
    _imageId   string
    // 发票代码，success=true时必填
    _invoiceCode   string
    // 开票日期，格式为yyyy-MM-dd，success=true时必填
    _invoiceDate   string
    // 发票种类，1=普票，2=专票，success=true时必填
    _invoiceKind   int64
    // 发票号码，success=true时必填
    _invoiceNo   string
    // 开票请求标识，扫描驱动回传type=1时填批次号
    _reqIndex   string
    // ocr结果，true=成功，false=失败
    _success   bool
    // 不含税金额，ocr结果为专票，success=true并且invoiceKind=2时必填
    _sumPrice   string
    // 请求类型，0=阿里主动发起的cor，1=扫描驱动服务商主动回传ocr结果
    _type   int64
}

// 初始化AlibabaEinvoiceIncomeOcrReturnRequest对象
func NewAlibabaEinvoiceIncomeOcrReturnRequest() *AlibabaEinvoiceIncomeOcrReturnRequest{
    return &AlibabaEinvoiceIncomeOcrReturnRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetApiMethodName() string {
    return "alibaba.einvoice.income.ocr.return"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Checksum Setter
// 校验码，ocr结果为普票，success=true并且invoiceKind=1时必填
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetChecksum(_checksum string) error {
    r._checksum = _checksum
    r.Set("checksum", _checksum)
    return nil
}

// Checksum Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetChecksum() string {
    return r._checksum
}
// ErrorCode Setter
// 错误码，success=false是必填
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetErrorCode(_errorCode string) error {
    r._errorCode = _errorCode
    r.Set("error_code", _errorCode)
    return nil
}

// ErrorCode Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetErrorCode() string {
    return r._errorCode
}
// ErrorMessage Setter
// 错误消息，success=false是必填
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetErrorMessage(_errorMessage string) error {
    r._errorMessage = _errorMessage
    r.Set("error_message", _errorMessage)
    return nil
}

// ErrorMessage Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetErrorMessage() string {
    return r._errorMessage
}
// ImageData Setter
// 发票ocr影像文件，type=1时必填
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetImageData(_imageData *model.File) error {
    r._imageData = _imageData
    r.Set("image_data", _imageData)
    return nil
}

// ImageData Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetImageData() *model.File {
    return r._imageData
}
// ImageId Setter
// 发票ocr影像编号，type=1时必填
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetImageId(_imageId string) error {
    r._imageId = _imageId
    r.Set("image_id", _imageId)
    return nil
}

// ImageId Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetImageId() string {
    return r._imageId
}
// InvoiceCode Setter
// 发票代码，success=true时必填
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetInvoiceCode(_invoiceCode string) error {
    r._invoiceCode = _invoiceCode
    r.Set("invoice_code", _invoiceCode)
    return nil
}

// InvoiceCode Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetInvoiceCode() string {
    return r._invoiceCode
}
// InvoiceDate Setter
// 开票日期，格式为yyyy-MM-dd，success=true时必填
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetInvoiceDate(_invoiceDate string) error {
    r._invoiceDate = _invoiceDate
    r.Set("invoice_date", _invoiceDate)
    return nil
}

// InvoiceDate Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetInvoiceDate() string {
    return r._invoiceDate
}
// InvoiceKind Setter
// 发票种类，1=普票，2=专票，success=true时必填
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetInvoiceKind(_invoiceKind int64) error {
    r._invoiceKind = _invoiceKind
    r.Set("invoice_kind", _invoiceKind)
    return nil
}

// InvoiceKind Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetInvoiceKind() int64 {
    return r._invoiceKind
}
// InvoiceNo Setter
// 发票号码，success=true时必填
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetInvoiceNo(_invoiceNo string) error {
    r._invoiceNo = _invoiceNo
    r.Set("invoice_no", _invoiceNo)
    return nil
}

// InvoiceNo Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetInvoiceNo() string {
    return r._invoiceNo
}
// ReqIndex Setter
// 开票请求标识，扫描驱动回传type=1时填批次号
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetReqIndex(_reqIndex string) error {
    r._reqIndex = _reqIndex
    r.Set("req_index", _reqIndex)
    return nil
}

// ReqIndex Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetReqIndex() string {
    return r._reqIndex
}
// Success Setter
// ocr结果，true=成功，false=失败
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetSuccess(_success bool) error {
    r._success = _success
    r.Set("success", _success)
    return nil
}

// Success Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetSuccess() bool {
    return r._success
}
// SumPrice Setter
// 不含税金额，ocr结果为专票，success=true并且invoiceKind=2时必填
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetSumPrice(_sumPrice string) error {
    r._sumPrice = _sumPrice
    r.Set("sum_price", _sumPrice)
    return nil
}

// SumPrice Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetSumPrice() string {
    return r._sumPrice
}
// Type Setter
// 请求类型，0=阿里主动发起的cor，1=扫描驱动服务商主动回传ocr结果
func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetType() int64 {
    return r._type
}
