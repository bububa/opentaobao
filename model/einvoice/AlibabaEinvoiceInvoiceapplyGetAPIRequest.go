package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceInvoiceapplyGetAPIRequest 获取商家的开票申请 API请求
// alibaba.einvoice.invoiceapply.get
//
// 开票服务商接收到商家发起的开票申请消息后，调用此接口拉取商家详细的开票申请内容
type AlibabaEinvoiceInvoiceapplyGetAPIRequest struct {
	model.Params
	// 开票申请id
	_applyId string
}

// NewAlibabaEinvoiceInvoiceapplyGetRequest 初始化AlibabaEinvoiceInvoiceapplyGetAPIRequest对象
func NewAlibabaEinvoiceInvoiceapplyGetRequest() *AlibabaEinvoiceInvoiceapplyGetAPIRequest {
	return &AlibabaEinvoiceInvoiceapplyGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceInvoiceapplyGetAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.invoiceapply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceInvoiceapplyGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetApplyId is ApplyId Setter
// 开票申请id
func (r *AlibabaEinvoiceInvoiceapplyGetAPIRequest) SetApplyId(_applyId string) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r AlibabaEinvoiceInvoiceapplyGetAPIRequest) GetApplyId() string {
	return r._applyId
}
