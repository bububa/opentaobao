package einvoice

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceInvoiceapplyUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.invoiceapply.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceInvoiceapplyUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
