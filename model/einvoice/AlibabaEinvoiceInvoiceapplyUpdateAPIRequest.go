package einvoice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceInvoiceapplyUpdateAPIRequest 商家开票申请单状态回传 API请求
// alibaba.einvoice.invoiceapply.update
//
// 开票服务商更新商家开票申请单状态
type AlibabaEinvoiceInvoiceapplyUpdateAPIRequest struct {
	model.Params
	// 申请单id
	_applyId string
	// 扩展信息,目前用于回传文本及物流消息
	_exInfo string
	// 申请单状态：1：待确认，2：开票中，3：拒绝开票，4：发票已发出，0：完成开票
	_status int64
}

// NewAlibabaEinvoiceInvoiceapplyUpdateRequest 初始化AlibabaEinvoiceInvoiceapplyUpdateAPIRequest对象
func NewAlibabaEinvoiceInvoiceapplyUpdateRequest() *AlibabaEinvoiceInvoiceapplyUpdateAPIRequest {
	return &AlibabaEinvoiceInvoiceapplyUpdateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEinvoiceInvoiceapplyUpdateAPIRequest) Reset() {
	r._applyId = ""
	r._exInfo = ""
	r._status = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceInvoiceapplyUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.invoiceapply.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceInvoiceapplyUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceInvoiceapplyUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyId is ApplyId Setter
// 申请单id
func (r *AlibabaEinvoiceInvoiceapplyUpdateAPIRequest) SetApplyId(_applyId string) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r AlibabaEinvoiceInvoiceapplyUpdateAPIRequest) GetApplyId() string {
	return r._applyId
}

// SetExInfo is ExInfo Setter
// 扩展信息,目前用于回传文本及物流消息
func (r *AlibabaEinvoiceInvoiceapplyUpdateAPIRequest) SetExInfo(_exInfo string) error {
	r._exInfo = _exInfo
	r.Set("ex_info", _exInfo)
	return nil
}

// GetExInfo ExInfo Getter
func (r AlibabaEinvoiceInvoiceapplyUpdateAPIRequest) GetExInfo() string {
	return r._exInfo
}

// SetStatus is Status Setter
// 申请单状态：1：待确认，2：开票中，3：拒绝开票，4：发票已发出，0：完成开票
func (r *AlibabaEinvoiceInvoiceapplyUpdateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaEinvoiceInvoiceapplyUpdateAPIRequest) GetStatus() int64 {
	return r._status
}

var poolAlibabaEinvoiceInvoiceapplyUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEinvoiceInvoiceapplyUpdateRequest()
	},
}

// GetAlibabaEinvoiceInvoiceapplyUpdateRequest 从 sync.Pool 获取 AlibabaEinvoiceInvoiceapplyUpdateAPIRequest
func GetAlibabaEinvoiceInvoiceapplyUpdateAPIRequest() *AlibabaEinvoiceInvoiceapplyUpdateAPIRequest {
	return poolAlibabaEinvoiceInvoiceapplyUpdateAPIRequest.Get().(*AlibabaEinvoiceInvoiceapplyUpdateAPIRequest)
}

// ReleaseAlibabaEinvoiceInvoiceapplyUpdateAPIRequest 将 AlibabaEinvoiceInvoiceapplyUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaEinvoiceInvoiceapplyUpdateAPIRequest(v *AlibabaEinvoiceInvoiceapplyUpdateAPIRequest) {
	v.Reset()
	poolAlibabaEinvoiceInvoiceapplyUpdateAPIRequest.Put(v)
}
