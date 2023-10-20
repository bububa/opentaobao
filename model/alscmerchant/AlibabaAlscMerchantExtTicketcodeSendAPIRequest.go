package alscmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscMerchantExtTicketcodeSendAPIRequest 异步发码 API请求
// alibaba.alsc.merchant.ext.ticketcode.send
//
// 外部券异步发码
type AlibabaAlscMerchantExtTicketcodeSendAPIRequest struct {
	model.Params
	// 外部券异步发码
	_sendRequest *ExternalTicketSendRequest
}

// NewAlibabaAlscMerchantExtTicketcodeSendRequest 初始化AlibabaAlscMerchantExtTicketcodeSendAPIRequest对象
func NewAlibabaAlscMerchantExtTicketcodeSendRequest() *AlibabaAlscMerchantExtTicketcodeSendAPIRequest {
	return &AlibabaAlscMerchantExtTicketcodeSendAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscMerchantExtTicketcodeSendAPIRequest) Reset() {
	r._sendRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscMerchantExtTicketcodeSendAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.merchant.ext.ticketcode.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscMerchantExtTicketcodeSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscMerchantExtTicketcodeSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSendRequest is SendRequest Setter
// 外部券异步发码
func (r *AlibabaAlscMerchantExtTicketcodeSendAPIRequest) SetSendRequest(_sendRequest *ExternalTicketSendRequest) error {
	r._sendRequest = _sendRequest
	r.Set("send_request", _sendRequest)
	return nil
}

// GetSendRequest SendRequest Getter
func (r AlibabaAlscMerchantExtTicketcodeSendAPIRequest) GetSendRequest() *ExternalTicketSendRequest {
	return r._sendRequest
}

var poolAlibabaAlscMerchantExtTicketcodeSendAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscMerchantExtTicketcodeSendRequest()
	},
}

// GetAlibabaAlscMerchantExtTicketcodeSendRequest 从 sync.Pool 获取 AlibabaAlscMerchantExtTicketcodeSendAPIRequest
func GetAlibabaAlscMerchantExtTicketcodeSendAPIRequest() *AlibabaAlscMerchantExtTicketcodeSendAPIRequest {
	return poolAlibabaAlscMerchantExtTicketcodeSendAPIRequest.Get().(*AlibabaAlscMerchantExtTicketcodeSendAPIRequest)
}

// ReleaseAlibabaAlscMerchantExtTicketcodeSendAPIRequest 将 AlibabaAlscMerchantExtTicketcodeSendAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscMerchantExtTicketcodeSendAPIRequest(v *AlibabaAlscMerchantExtTicketcodeSendAPIRequest) {
	v.Reset()
	poolAlibabaAlscMerchantExtTicketcodeSendAPIRequest.Put(v)
}
