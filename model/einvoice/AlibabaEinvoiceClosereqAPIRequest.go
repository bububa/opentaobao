package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceClosereqAPIRequest 关闭开票失败请求（失败列表可重试） API请求
// alibaba.einvoice.closereq
//
// 关闭失败开票请求，避免造成重复开票
type AlibabaEinvoiceClosereqAPIRequest struct {
	model.Params
	// 流水号
	_serialNo string
	// 税号
	_payeeRegisterNo string
}

// NewAlibabaEinvoiceClosereqRequest 初始化AlibabaEinvoiceClosereqAPIRequest对象
func NewAlibabaEinvoiceClosereqRequest() *AlibabaEinvoiceClosereqAPIRequest {
	return &AlibabaEinvoiceClosereqAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceClosereqAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.closereq"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceClosereqAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SerialNo Setter
// 流水号
func (r *AlibabaEinvoiceClosereqAPIRequest) SetSerialNo(_serialNo string) error {
	r._serialNo = _serialNo
	r.Set("serial_no", _serialNo)
	return nil
}

// Get SerialNo Getter
func (r AlibabaEinvoiceClosereqAPIRequest) GetSerialNo() string {
	return r._serialNo
}

// Set is PayeeRegisterNo Setter
// 税号
func (r *AlibabaEinvoiceClosereqAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// Get PayeeRegisterNo Getter
func (r AlibabaEinvoiceClosereqAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}
