package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceSerialnoGenerateAPIRequest 获取统一开票流水号 API请求
// alibaba.einvoice.serialno.generate
//
// erp调用开票请求时需要一个开票流水号，此接口就提供了统一的开票流水号，避免了不同系统的冲突
type AlibabaEinvoiceSerialnoGenerateAPIRequest struct {
	model.Params
}

// NewAlibabaEinvoiceSerialnoGenerateRequest 初始化AlibabaEinvoiceSerialnoGenerateAPIRequest对象
func NewAlibabaEinvoiceSerialnoGenerateRequest() *AlibabaEinvoiceSerialnoGenerateAPIRequest {
	return &AlibabaEinvoiceSerialnoGenerateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceSerialnoGenerateAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.serialno.generate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceSerialnoGenerateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceSerialnoGenerateAPIRequest) GetRawParams() model.Params {
	return r.Params
}
