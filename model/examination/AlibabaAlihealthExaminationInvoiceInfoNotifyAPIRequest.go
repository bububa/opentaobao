package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationinvoiceinfonotifyAPIRequest 体检机构同步发票信息给阿里健康 API请求
// alibaba.alihealth.examination.invoice.info.notify
//
// 体检机构向阿里健康同步发票信息
type AlibabaalihealthexaminationinvoiceinfonotifyAPIRequest struct {
	model.Params
	// 阿里健康预约凭证
	_reserveNumber string
	// 开票状态；（have_submit已提交、invoice_done已开票）
	_invoiceStatus string
	// 发票访问地址；（invoice_status在已开票状态下必填）
	_invoiceUrl string
}

// NewAlibabaalihealthexaminationinvoiceinfonotifyRequest 初始化AlibabaalihealthexaminationinvoiceinfonotifyAPIRequest对象
func NewAlibabaalihealthexaminationinvoiceinfonotifyRequest() *AlibabaalihealthexaminationinvoiceinfonotifyAPIRequest {
	return &AlibabaalihealthexaminationinvoiceinfonotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthexaminationinvoiceinfonotifyAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.invoice.info.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthexaminationinvoiceinfonotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthexaminationinvoiceinfonotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReserveNumber is ReserveNumber Setter
// 阿里健康预约凭证
func (r *AlibabaalihealthexaminationinvoiceinfonotifyAPIRequest) SetReserveNumber(_reserveNumber string) error {
	r._reserveNumber = _reserveNumber
	r.Set("reserve_number", _reserveNumber)
	return nil
}

// GetReserveNumber ReserveNumber Getter
func (r AlibabaalihealthexaminationinvoiceinfonotifyAPIRequest) GetReserveNumber() string {
	return r._reserveNumber
}

// SetInvoiceStatus is InvoiceStatus Setter
// 开票状态；（have_submit已提交、invoice_done已开票）
func (r *AlibabaalihealthexaminationinvoiceinfonotifyAPIRequest) SetInvoiceStatus(_invoiceStatus string) error {
	r._invoiceStatus = _invoiceStatus
	r.Set("invoice_status", _invoiceStatus)
	return nil
}

// GetInvoiceStatus InvoiceStatus Getter
func (r AlibabaalihealthexaminationinvoiceinfonotifyAPIRequest) GetInvoiceStatus() string {
	return r._invoiceStatus
}

// SetInvoiceUrl is InvoiceUrl Setter
// 发票访问地址；（invoice_status在已开票状态下必填）
func (r *AlibabaalihealthexaminationinvoiceinfonotifyAPIRequest) SetInvoiceUrl(_invoiceUrl string) error {
	r._invoiceUrl = _invoiceUrl
	r.Set("invoice_url", _invoiceUrl)
	return nil
}

// GetInvoiceUrl InvoiceUrl Getter
func (r AlibabaalihealthexaminationinvoiceinfonotifyAPIRequest) GetInvoiceUrl() string {
	return r._invoiceUrl
}
