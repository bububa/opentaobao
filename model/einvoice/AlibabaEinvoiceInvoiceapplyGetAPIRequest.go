package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceinvoiceapplygetAPIRequest 获取商家的开票申请 API请求
// alibaba.einvoice.invoiceapply.get
//
// 开票服务商接收到商家发起的开票申请消息后，调用此接口拉取商家详细的开票申请内容
type AlibabaeinvoiceinvoiceapplygetAPIRequest struct {
	model.Params
	// 开票申请id
	_applyId string
}

// NewAlibabaeinvoiceinvoiceapplygetRequest 初始化AlibabaeinvoiceinvoiceapplygetAPIRequest对象
func NewAlibabaeinvoiceinvoiceapplygetRequest() *AlibabaeinvoiceinvoiceapplygetAPIRequest {
	return &AlibabaeinvoiceinvoiceapplygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoiceinvoiceapplygetAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.invoiceapply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoiceinvoiceapplygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoiceinvoiceapplygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyId is ApplyId Setter
// 开票申请id
func (r *AlibabaeinvoiceinvoiceapplygetAPIRequest) SetApplyId(_applyId string) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r AlibabaeinvoiceinvoiceapplygetAPIRequest) GetApplyId() string {
	return r._applyId
}
