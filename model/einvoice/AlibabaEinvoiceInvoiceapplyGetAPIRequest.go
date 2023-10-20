package einvoice

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEinvoiceInvoiceapplyGetAPIRequest) Reset() {
	r._applyId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceInvoiceapplyGetAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.invoiceapply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceInvoiceapplyGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceInvoiceapplyGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaEinvoiceInvoiceapplyGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEinvoiceInvoiceapplyGetRequest()
	},
}

// GetAlibabaEinvoiceInvoiceapplyGetRequest 从 sync.Pool 获取 AlibabaEinvoiceInvoiceapplyGetAPIRequest
func GetAlibabaEinvoiceInvoiceapplyGetAPIRequest() *AlibabaEinvoiceInvoiceapplyGetAPIRequest {
	return poolAlibabaEinvoiceInvoiceapplyGetAPIRequest.Get().(*AlibabaEinvoiceInvoiceapplyGetAPIRequest)
}

// ReleaseAlibabaEinvoiceInvoiceapplyGetAPIRequest 将 AlibabaEinvoiceInvoiceapplyGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaEinvoiceInvoiceapplyGetAPIRequest(v *AlibabaEinvoiceInvoiceapplyGetAPIRequest) {
	v.Reset()
	poolAlibabaEinvoiceInvoiceapplyGetAPIRequest.Put(v)
}
