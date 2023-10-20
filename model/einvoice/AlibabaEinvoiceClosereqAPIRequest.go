package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceclosereqAPIRequest 关闭开票失败请求（失败列表可重试） API请求
// alibaba.einvoice.closereq
//
// 关闭失败开票请求，避免造成重复开票
type AlibabaeinvoiceclosereqAPIRequest struct {
	model.Params
	// 流水号
	_serialNo string
	// 税号
	_payeeRegisterNo string
}

// NewAlibabaeinvoiceclosereqRequest 初始化AlibabaeinvoiceclosereqAPIRequest对象
func NewAlibabaeinvoiceclosereqRequest() *AlibabaeinvoiceclosereqAPIRequest {
	return &AlibabaeinvoiceclosereqAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoiceclosereqAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.closereq"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoiceclosereqAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoiceclosereqAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSerialNo is SerialNo Setter
// 流水号
func (r *AlibabaeinvoiceclosereqAPIRequest) SetSerialNo(_serialNo string) error {
	r._serialNo = _serialNo
	r.Set("serial_no", _serialNo)
	return nil
}

// GetSerialNo SerialNo Getter
func (r AlibabaeinvoiceclosereqAPIRequest) GetSerialNo() string {
	return r._serialNo
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 税号
func (r *AlibabaeinvoiceclosereqAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaeinvoiceclosereqAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}
