package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceSerialnoBatchGenerateAPIRequest 开票流水号批量生成接口 API请求
// alibaba.einvoice.serialno.batch.generate
//
// 批量获取开票流水号接口。此接口1次返回1000条开票流水号，每个应用每天限流1000次调用。
// 优先使用alibaba.einvoice.serial.generate。
type AlibabaEinvoiceSerialnoBatchGenerateAPIRequest struct {
	model.Params
}

// NewAlibabaEinvoiceSerialnoBatchGenerateRequest 初始化AlibabaEinvoiceSerialnoBatchGenerateAPIRequest对象
func NewAlibabaEinvoiceSerialnoBatchGenerateRequest() *AlibabaEinvoiceSerialnoBatchGenerateAPIRequest {
	return &AlibabaEinvoiceSerialnoBatchGenerateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceSerialnoBatchGenerateAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.serialno.batch.generate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceSerialnoBatchGenerateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceSerialnoBatchGenerateAPIRequest) GetRawParams() model.Params {
	return r.Params
}
