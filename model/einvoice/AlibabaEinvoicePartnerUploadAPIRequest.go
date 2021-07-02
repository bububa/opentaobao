package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoicePartnerUploadAPIRequest 服务商发票上传接口（非授权） API请求
// alibaba.einvoice.partner.upload
//
// 服务商发票上传接口（非授权）
type AlibabaEinvoicePartnerUploadAPIRequest struct {
	model.Params
	// 原蓝票发票号码
	_normalInvoiceNo string
	// 原蓝票发票代码
	_normalInvoiceCode string
	// 销方税号
	_payeeRegisterNo string
	// 发票数据，upload_type=0且invoiceKind=0电子发票时必填
	_invoiceFileData *model.File
	// 发票号码，upload_type=0时必填
	_invoiceNo string
	// 发票代码，upload_type=0时必填
	_invoiceCode string
	// 开票日期，upload_type=0时必填
	_invoiceDate string
	// 密码区
	_cipherText string
	// 机器编号
	_deviceNo string
	// 校验码
	_antiFakeCode string
	// 发票类型，upload_type=0且invoiceKind=0电子发票时必填，暂时只支持pdf
	_fileDataType string
	// 原蓝票的reqIndex
	_reqIndex string
	// 发票种类，0=电子发票，1=纸质普票，2=纸质专票
	_invoiceKind int64
	// 上传的类型，0=冲红上传，1=作废上传
	_uploadType int64
}

// NewAlibabaEinvoicePartnerUploadRequest 初始化AlibabaEinvoicePartnerUploadAPIRequest对象
func NewAlibabaEinvoicePartnerUploadRequest() *AlibabaEinvoicePartnerUploadAPIRequest {
	return &AlibabaEinvoicePartnerUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoicePartnerUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.partner.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoicePartnerUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetNormalInvoiceNo is NormalInvoiceNo Setter
// 原蓝票发票号码
func (r *AlibabaEinvoicePartnerUploadAPIRequest) SetNormalInvoiceNo(_normalInvoiceNo string) error {
	r._normalInvoiceNo = _normalInvoiceNo
	r.Set("normal_invoice_no", _normalInvoiceNo)
	return nil
}

// GetNormalInvoiceNo NormalInvoiceNo Getter
func (r AlibabaEinvoicePartnerUploadAPIRequest) GetNormalInvoiceNo() string {
	return r._normalInvoiceNo
}

// SetNormalInvoiceCode is NormalInvoiceCode Setter
// 原蓝票发票代码
func (r *AlibabaEinvoicePartnerUploadAPIRequest) SetNormalInvoiceCode(_normalInvoiceCode string) error {
	r._normalInvoiceCode = _normalInvoiceCode
	r.Set("normal_invoice_code", _normalInvoiceCode)
	return nil
}

// GetNormalInvoiceCode NormalInvoiceCode Getter
func (r AlibabaEinvoicePartnerUploadAPIRequest) GetNormalInvoiceCode() string {
	return r._normalInvoiceCode
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 销方税号
func (r *AlibabaEinvoicePartnerUploadAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaEinvoicePartnerUploadAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetInvoiceFileData is InvoiceFileData Setter
// 发票数据，upload_type=0且invoiceKind=0电子发票时必填
func (r *AlibabaEinvoicePartnerUploadAPIRequest) SetInvoiceFileData(_invoiceFileData *model.File) error {
	r._invoiceFileData = _invoiceFileData
	r.Set("invoice_file_data", _invoiceFileData)
	return nil
}

// GetInvoiceFileData InvoiceFileData Getter
func (r AlibabaEinvoicePartnerUploadAPIRequest) GetInvoiceFileData() *model.File {
	return r._invoiceFileData
}

// SetInvoiceNo is InvoiceNo Setter
// 发票号码，upload_type=0时必填
func (r *AlibabaEinvoicePartnerUploadAPIRequest) SetInvoiceNo(_invoiceNo string) error {
	r._invoiceNo = _invoiceNo
	r.Set("invoice_no", _invoiceNo)
	return nil
}

// GetInvoiceNo InvoiceNo Getter
func (r AlibabaEinvoicePartnerUploadAPIRequest) GetInvoiceNo() string {
	return r._invoiceNo
}

// SetInvoiceCode is InvoiceCode Setter
// 发票代码，upload_type=0时必填
func (r *AlibabaEinvoicePartnerUploadAPIRequest) SetInvoiceCode(_invoiceCode string) error {
	r._invoiceCode = _invoiceCode
	r.Set("invoice_code", _invoiceCode)
	return nil
}

// GetInvoiceCode InvoiceCode Getter
func (r AlibabaEinvoicePartnerUploadAPIRequest) GetInvoiceCode() string {
	return r._invoiceCode
}

// SetInvoiceDate is InvoiceDate Setter
// 开票日期，upload_type=0时必填
func (r *AlibabaEinvoicePartnerUploadAPIRequest) SetInvoiceDate(_invoiceDate string) error {
	r._invoiceDate = _invoiceDate
	r.Set("invoice_date", _invoiceDate)
	return nil
}

// GetInvoiceDate InvoiceDate Getter
func (r AlibabaEinvoicePartnerUploadAPIRequest) GetInvoiceDate() string {
	return r._invoiceDate
}

// SetCipherText is CipherText Setter
// 密码区
func (r *AlibabaEinvoicePartnerUploadAPIRequest) SetCipherText(_cipherText string) error {
	r._cipherText = _cipherText
	r.Set("cipher_text", _cipherText)
	return nil
}

// GetCipherText CipherText Getter
func (r AlibabaEinvoicePartnerUploadAPIRequest) GetCipherText() string {
	return r._cipherText
}

// SetDeviceNo is DeviceNo Setter
// 机器编号
func (r *AlibabaEinvoicePartnerUploadAPIRequest) SetDeviceNo(_deviceNo string) error {
	r._deviceNo = _deviceNo
	r.Set("device_no", _deviceNo)
	return nil
}

// GetDeviceNo DeviceNo Getter
func (r AlibabaEinvoicePartnerUploadAPIRequest) GetDeviceNo() string {
	return r._deviceNo
}

// SetAntiFakeCode is AntiFakeCode Setter
// 校验码
func (r *AlibabaEinvoicePartnerUploadAPIRequest) SetAntiFakeCode(_antiFakeCode string) error {
	r._antiFakeCode = _antiFakeCode
	r.Set("anti_fake_code", _antiFakeCode)
	return nil
}

// GetAntiFakeCode AntiFakeCode Getter
func (r AlibabaEinvoicePartnerUploadAPIRequest) GetAntiFakeCode() string {
	return r._antiFakeCode
}

// SetFileDataType is FileDataType Setter
// 发票类型，upload_type=0且invoiceKind=0电子发票时必填，暂时只支持pdf
func (r *AlibabaEinvoicePartnerUploadAPIRequest) SetFileDataType(_fileDataType string) error {
	r._fileDataType = _fileDataType
	r.Set("file_data_type", _fileDataType)
	return nil
}

// GetFileDataType FileDataType Getter
func (r AlibabaEinvoicePartnerUploadAPIRequest) GetFileDataType() string {
	return r._fileDataType
}

// SetReqIndex is ReqIndex Setter
// 原蓝票的reqIndex
func (r *AlibabaEinvoicePartnerUploadAPIRequest) SetReqIndex(_reqIndex string) error {
	r._reqIndex = _reqIndex
	r.Set("req_index", _reqIndex)
	return nil
}

// GetReqIndex ReqIndex Getter
func (r AlibabaEinvoicePartnerUploadAPIRequest) GetReqIndex() string {
	return r._reqIndex
}

// SetInvoiceKind is InvoiceKind Setter
// 发票种类，0=电子发票，1=纸质普票，2=纸质专票
func (r *AlibabaEinvoicePartnerUploadAPIRequest) SetInvoiceKind(_invoiceKind int64) error {
	r._invoiceKind = _invoiceKind
	r.Set("invoice_kind", _invoiceKind)
	return nil
}

// GetInvoiceKind InvoiceKind Getter
func (r AlibabaEinvoicePartnerUploadAPIRequest) GetInvoiceKind() int64 {
	return r._invoiceKind
}

// SetUploadType is UploadType Setter
// 上传的类型，0=冲红上传，1=作废上传
func (r *AlibabaEinvoicePartnerUploadAPIRequest) SetUploadType(_uploadType int64) error {
	r._uploadType = _uploadType
	r.Set("upload_type", _uploadType)
	return nil
}

// GetUploadType UploadType Getter
func (r AlibabaEinvoicePartnerUploadAPIRequest) GetUploadType() int64 {
	return r._uploadType
}
