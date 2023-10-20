package einvoice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceRedCreatereqAPIRequest 发票冲红接口 API请求
// alibaba.einvoice.red.createreq
//
// 发票冲红接口，通过蓝票流水号或者发票号码+发票代码进行冲红
type AlibabaEinvoiceRedCreatereqAPIRequest struct {
	model.Params
	// 销售方税号
	_payeeRegisterNo string
	// 蓝票流水号，优先级高于发票代码+发票号码
	_blueSerialNo string
	// 红票流水号
	_redSerialNo string
	// 蓝票发票代码
	_invoiceCode string
	// 蓝票发票号码
	_invoiceNo string
}

// NewAlibabaEinvoiceRedCreatereqRequest 初始化AlibabaEinvoiceRedCreatereqAPIRequest对象
func NewAlibabaEinvoiceRedCreatereqRequest() *AlibabaEinvoiceRedCreatereqAPIRequest {
	return &AlibabaEinvoiceRedCreatereqAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEinvoiceRedCreatereqAPIRequest) Reset() {
	r._payeeRegisterNo = ""
	r._blueSerialNo = ""
	r._redSerialNo = ""
	r._invoiceCode = ""
	r._invoiceNo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceRedCreatereqAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.red.createreq"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceRedCreatereqAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceRedCreatereqAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 销售方税号
func (r *AlibabaEinvoiceRedCreatereqAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaEinvoiceRedCreatereqAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetBlueSerialNo is BlueSerialNo Setter
// 蓝票流水号，优先级高于发票代码+发票号码
func (r *AlibabaEinvoiceRedCreatereqAPIRequest) SetBlueSerialNo(_blueSerialNo string) error {
	r._blueSerialNo = _blueSerialNo
	r.Set("blue_serial_no", _blueSerialNo)
	return nil
}

// GetBlueSerialNo BlueSerialNo Getter
func (r AlibabaEinvoiceRedCreatereqAPIRequest) GetBlueSerialNo() string {
	return r._blueSerialNo
}

// SetRedSerialNo is RedSerialNo Setter
// 红票流水号
func (r *AlibabaEinvoiceRedCreatereqAPIRequest) SetRedSerialNo(_redSerialNo string) error {
	r._redSerialNo = _redSerialNo
	r.Set("red_serial_no", _redSerialNo)
	return nil
}

// GetRedSerialNo RedSerialNo Getter
func (r AlibabaEinvoiceRedCreatereqAPIRequest) GetRedSerialNo() string {
	return r._redSerialNo
}

// SetInvoiceCode is InvoiceCode Setter
// 蓝票发票代码
func (r *AlibabaEinvoiceRedCreatereqAPIRequest) SetInvoiceCode(_invoiceCode string) error {
	r._invoiceCode = _invoiceCode
	r.Set("invoice_code", _invoiceCode)
	return nil
}

// GetInvoiceCode InvoiceCode Getter
func (r AlibabaEinvoiceRedCreatereqAPIRequest) GetInvoiceCode() string {
	return r._invoiceCode
}

// SetInvoiceNo is InvoiceNo Setter
// 蓝票发票号码
func (r *AlibabaEinvoiceRedCreatereqAPIRequest) SetInvoiceNo(_invoiceNo string) error {
	r._invoiceNo = _invoiceNo
	r.Set("invoice_no", _invoiceNo)
	return nil
}

// GetInvoiceNo InvoiceNo Getter
func (r AlibabaEinvoiceRedCreatereqAPIRequest) GetInvoiceNo() string {
	return r._invoiceNo
}

var poolAlibabaEinvoiceRedCreatereqAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEinvoiceRedCreatereqRequest()
	},
}

// GetAlibabaEinvoiceRedCreatereqRequest 从 sync.Pool 获取 AlibabaEinvoiceRedCreatereqAPIRequest
func GetAlibabaEinvoiceRedCreatereqAPIRequest() *AlibabaEinvoiceRedCreatereqAPIRequest {
	return poolAlibabaEinvoiceRedCreatereqAPIRequest.Get().(*AlibabaEinvoiceRedCreatereqAPIRequest)
}

// ReleaseAlibabaEinvoiceRedCreatereqAPIRequest 将 AlibabaEinvoiceRedCreatereqAPIRequest 放入 sync.Pool
func ReleaseAlibabaEinvoiceRedCreatereqAPIRequest(v *AlibabaEinvoiceRedCreatereqAPIRequest) {
	v.Reset()
	poolAlibabaEinvoiceRedCreatereqAPIRequest.Put(v)
}
