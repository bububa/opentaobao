package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceprodresultfileurlgetAPIRequest 发票中台-发票文件下载地址查询 API请求
// alibaba.einvoice.prod.result.fileurl.get
//
// 发票文件下载地址查询，外部ISV通过该接口可以查对应发票文件
type AlibabaeinvoiceprodresultfileurlgetAPIRequest struct {
	model.Params
	// 业务平台商户ID/卖家用户ID
	_platformUserId string
	// 发票号码
	_invoiceNo string
	// 发票代码
	_invoiceCode string
	// 发票文件类型，小写，pdf/ofd/jpg
	_fileType string
	// 业务平台code, 由发票中台分配
	_platformCode string
}

// NewAlibabaeinvoiceprodresultfileurlgetRequest 初始化AlibabaeinvoiceprodresultfileurlgetAPIRequest对象
func NewAlibabaeinvoiceprodresultfileurlgetRequest() *AlibabaeinvoiceprodresultfileurlgetAPIRequest {
	return &AlibabaeinvoiceprodresultfileurlgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoiceprodresultfileurlgetAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.prod.result.fileurl.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoiceprodresultfileurlgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoiceprodresultfileurlgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPlatformUserId is PlatformUserId Setter
// 业务平台商户ID/卖家用户ID
func (r *AlibabaeinvoiceprodresultfileurlgetAPIRequest) SetPlatformUserId(_platformUserId string) error {
	r._platformUserId = _platformUserId
	r.Set("platform_user_id", _platformUserId)
	return nil
}

// GetPlatformUserId PlatformUserId Getter
func (r AlibabaeinvoiceprodresultfileurlgetAPIRequest) GetPlatformUserId() string {
	return r._platformUserId
}

// SetInvoiceNo is InvoiceNo Setter
// 发票号码
func (r *AlibabaeinvoiceprodresultfileurlgetAPIRequest) SetInvoiceNo(_invoiceNo string) error {
	r._invoiceNo = _invoiceNo
	r.Set("invoice_no", _invoiceNo)
	return nil
}

// GetInvoiceNo InvoiceNo Getter
func (r AlibabaeinvoiceprodresultfileurlgetAPIRequest) GetInvoiceNo() string {
	return r._invoiceNo
}

// SetInvoiceCode is InvoiceCode Setter
// 发票代码
func (r *AlibabaeinvoiceprodresultfileurlgetAPIRequest) SetInvoiceCode(_invoiceCode string) error {
	r._invoiceCode = _invoiceCode
	r.Set("invoice_code", _invoiceCode)
	return nil
}

// GetInvoiceCode InvoiceCode Getter
func (r AlibabaeinvoiceprodresultfileurlgetAPIRequest) GetInvoiceCode() string {
	return r._invoiceCode
}

// SetFileType is FileType Setter
// 发票文件类型，小写，pdf/ofd/jpg
func (r *AlibabaeinvoiceprodresultfileurlgetAPIRequest) SetFileType(_fileType string) error {
	r._fileType = _fileType
	r.Set("file_type", _fileType)
	return nil
}

// GetFileType FileType Getter
func (r AlibabaeinvoiceprodresultfileurlgetAPIRequest) GetFileType() string {
	return r._fileType
}

// SetPlatformCode is PlatformCode Setter
// 业务平台code, 由发票中台分配
func (r *AlibabaeinvoiceprodresultfileurlgetAPIRequest) SetPlatformCode(_platformCode string) error {
	r._platformCode = _platformCode
	r.Set("platform_code", _platformCode)
	return nil
}

// GetPlatformCode PlatformCode Getter
func (r AlibabaeinvoiceprodresultfileurlgetAPIRequest) GetPlatformCode() string {
	return r._platformCode
}
