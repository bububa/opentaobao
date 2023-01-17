package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceIncomeOcrReturnAPIRequest 服务商回传发票ocr的结果 API请求
// alibaba.einvoice.income.ocr.return
//
// 服务商回传发票ocr的结果，分两种场景：扫描驱动服务商主动回传；阿里主动发起的ocr回传
type AlibabaEinvoiceIncomeOcrReturnAPIRequest struct {
	model.Params
	// 校验码，ocr结果为普票，success=true并且invoiceKind=1时必填
	_checksum string
	// 错误码，success=false是必填
	_errorCode string
	// 错误消息，success=false是必填
	_errorMessage string
	// 发票ocr影像编号，type=1时必填
	_imageId string
	// 发票代码，success=true时必填
	_invoiceCode string
	// 开票日期，格式为yyyy-MM-dd，success=true时必填
	_invoiceDate string
	// 发票号码，success=true时必填
	_invoiceNo string
	// 开票请求标识，扫描驱动回传type=1时填批次号
	_reqIndex string
	// 不含税金额，ocr结果为专票，success=true并且invoiceKind=2时必填
	_sumPrice string
	// 发票ocr影像文件，type=1时必填
	_imageData *model.File
	// 发票种类，1=普票，2=专票，success=true时必填
	_invoiceKind int64
	// 请求类型，0=阿里主动发起的cor，1=扫描驱动服务商主动回传ocr结果
	_type int64
	// ocr结果，true=成功，false=失败
	_success bool
}

// NewAlibabaEinvoiceIncomeOcrReturnRequest 初始化AlibabaEinvoiceIncomeOcrReturnAPIRequest对象
func NewAlibabaEinvoiceIncomeOcrReturnRequest() *AlibabaEinvoiceIncomeOcrReturnAPIRequest {
	return &AlibabaEinvoiceIncomeOcrReturnAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceIncomeOcrReturnAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.income.ocr.return"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceIncomeOcrReturnAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceIncomeOcrReturnAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChecksum is Checksum Setter
// 校验码，ocr结果为普票，success=true并且invoiceKind=1时必填
func (r *AlibabaEinvoiceIncomeOcrReturnAPIRequest) SetChecksum(_checksum string) error {
	r._checksum = _checksum
	r.Set("checksum", _checksum)
	return nil
}

// GetChecksum Checksum Getter
func (r AlibabaEinvoiceIncomeOcrReturnAPIRequest) GetChecksum() string {
	return r._checksum
}

// SetErrorCode is ErrorCode Setter
// 错误码，success=false是必填
func (r *AlibabaEinvoiceIncomeOcrReturnAPIRequest) SetErrorCode(_errorCode string) error {
	r._errorCode = _errorCode
	r.Set("error_code", _errorCode)
	return nil
}

// GetErrorCode ErrorCode Getter
func (r AlibabaEinvoiceIncomeOcrReturnAPIRequest) GetErrorCode() string {
	return r._errorCode
}

// SetErrorMessage is ErrorMessage Setter
// 错误消息，success=false是必填
func (r *AlibabaEinvoiceIncomeOcrReturnAPIRequest) SetErrorMessage(_errorMessage string) error {
	r._errorMessage = _errorMessage
	r.Set("error_message", _errorMessage)
	return nil
}

// GetErrorMessage ErrorMessage Getter
func (r AlibabaEinvoiceIncomeOcrReturnAPIRequest) GetErrorMessage() string {
	return r._errorMessage
}

// SetImageId is ImageId Setter
// 发票ocr影像编号，type=1时必填
func (r *AlibabaEinvoiceIncomeOcrReturnAPIRequest) SetImageId(_imageId string) error {
	r._imageId = _imageId
	r.Set("image_id", _imageId)
	return nil
}

// GetImageId ImageId Getter
func (r AlibabaEinvoiceIncomeOcrReturnAPIRequest) GetImageId() string {
	return r._imageId
}

// SetInvoiceCode is InvoiceCode Setter
// 发票代码，success=true时必填
func (r *AlibabaEinvoiceIncomeOcrReturnAPIRequest) SetInvoiceCode(_invoiceCode string) error {
	r._invoiceCode = _invoiceCode
	r.Set("invoice_code", _invoiceCode)
	return nil
}

// GetInvoiceCode InvoiceCode Getter
func (r AlibabaEinvoiceIncomeOcrReturnAPIRequest) GetInvoiceCode() string {
	return r._invoiceCode
}

// SetInvoiceDate is InvoiceDate Setter
// 开票日期，格式为yyyy-MM-dd，success=true时必填
func (r *AlibabaEinvoiceIncomeOcrReturnAPIRequest) SetInvoiceDate(_invoiceDate string) error {
	r._invoiceDate = _invoiceDate
	r.Set("invoice_date", _invoiceDate)
	return nil
}

// GetInvoiceDate InvoiceDate Getter
func (r AlibabaEinvoiceIncomeOcrReturnAPIRequest) GetInvoiceDate() string {
	return r._invoiceDate
}

// SetInvoiceNo is InvoiceNo Setter
// 发票号码，success=true时必填
func (r *AlibabaEinvoiceIncomeOcrReturnAPIRequest) SetInvoiceNo(_invoiceNo string) error {
	r._invoiceNo = _invoiceNo
	r.Set("invoice_no", _invoiceNo)
	return nil
}

// GetInvoiceNo InvoiceNo Getter
func (r AlibabaEinvoiceIncomeOcrReturnAPIRequest) GetInvoiceNo() string {
	return r._invoiceNo
}

// SetReqIndex is ReqIndex Setter
// 开票请求标识，扫描驱动回传type=1时填批次号
func (r *AlibabaEinvoiceIncomeOcrReturnAPIRequest) SetReqIndex(_reqIndex string) error {
	r._reqIndex = _reqIndex
	r.Set("req_index", _reqIndex)
	return nil
}

// GetReqIndex ReqIndex Getter
func (r AlibabaEinvoiceIncomeOcrReturnAPIRequest) GetReqIndex() string {
	return r._reqIndex
}

// SetSumPrice is SumPrice Setter
// 不含税金额，ocr结果为专票，success=true并且invoiceKind=2时必填
func (r *AlibabaEinvoiceIncomeOcrReturnAPIRequest) SetSumPrice(_sumPrice string) error {
	r._sumPrice = _sumPrice
	r.Set("sum_price", _sumPrice)
	return nil
}

// GetSumPrice SumPrice Getter
func (r AlibabaEinvoiceIncomeOcrReturnAPIRequest) GetSumPrice() string {
	return r._sumPrice
}

// SetImageData is ImageData Setter
// 发票ocr影像文件，type=1时必填
func (r *AlibabaEinvoiceIncomeOcrReturnAPIRequest) SetImageData(_imageData *model.File) error {
	r._imageData = _imageData
	r.Set("image_data", _imageData)
	return nil
}

// GetImageData ImageData Getter
func (r AlibabaEinvoiceIncomeOcrReturnAPIRequest) GetImageData() *model.File {
	return r._imageData
}

// SetInvoiceKind is InvoiceKind Setter
// 发票种类，1=普票，2=专票，success=true时必填
func (r *AlibabaEinvoiceIncomeOcrReturnAPIRequest) SetInvoiceKind(_invoiceKind int64) error {
	r._invoiceKind = _invoiceKind
	r.Set("invoice_kind", _invoiceKind)
	return nil
}

// GetInvoiceKind InvoiceKind Getter
func (r AlibabaEinvoiceIncomeOcrReturnAPIRequest) GetInvoiceKind() int64 {
	return r._invoiceKind
}

// SetType is Type Setter
// 请求类型，0=阿里主动发起的cor，1=扫描驱动服务商主动回传ocr结果
func (r *AlibabaEinvoiceIncomeOcrReturnAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaEinvoiceIncomeOcrReturnAPIRequest) GetType() int64 {
	return r._type
}

// SetSuccess is Success Setter
// ocr结果，true=成功，false=失败
func (r *AlibabaEinvoiceIncomeOcrReturnAPIRequest) SetSuccess(_success bool) error {
	r._success = _success
	r.Set("success", _success)
	return nil
}

// GetSuccess Success Getter
func (r AlibabaEinvoiceIncomeOcrReturnAPIRequest) GetSuccess() bool {
	return r._success
}
