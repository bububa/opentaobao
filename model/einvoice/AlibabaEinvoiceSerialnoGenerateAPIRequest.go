package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceserialnogenerateAPIRequest 获取统一开票流水号 API请求
// alibaba.einvoice.serialno.generate
//
// erp调用开票请求时需要一个开票流水号，此接口就提供了统一的开票流水号，避免了不同系统的冲突
type AlibabaeinvoiceserialnogenerateAPIRequest struct {
	model.Params
}

// NewAlibabaeinvoiceserialnogenerateRequest 初始化AlibabaeinvoiceserialnogenerateAPIRequest对象
func NewAlibabaeinvoiceserialnogenerateRequest() *AlibabaeinvoiceserialnogenerateAPIRequest {
	return &AlibabaeinvoiceserialnogenerateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoiceserialnogenerateAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.serialno.generate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoiceserialnogenerateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoiceserialnogenerateAPIRequest) GetRawParams() model.Params {
	return r.Params
}
