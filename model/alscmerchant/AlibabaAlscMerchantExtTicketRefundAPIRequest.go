package alscmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscMerchantExtTicketRefundAPIRequest 口碑凭证售后退 API请求
// alibaba.alsc.merchant.ext.ticket.refund
//
// 口碑凭证售后退
type AlibabaAlscMerchantExtTicketRefundAPIRequest struct {
	model.Params
	// 券核销流水号，针对该次核销发起售后退操作
	_transId string
	// 外部请求号，支持英文字母和数字，由开发者自行定义（不允许重复）
	_ticketRequestId string
	// 凭证码，包括内部凭证码和外部凭证码，内部凭证码为12位，纯数字，且唯一不重复
	_ticketCode string
}

// NewAlibabaAlscMerchantExtTicketRefundRequest 初始化AlibabaAlscMerchantExtTicketRefundAPIRequest对象
func NewAlibabaAlscMerchantExtTicketRefundRequest() *AlibabaAlscMerchantExtTicketRefundAPIRequest {
	return &AlibabaAlscMerchantExtTicketRefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscMerchantExtTicketRefundAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.merchant.ext.ticket.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscMerchantExtTicketRefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscMerchantExtTicketRefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTransId is TransId Setter
// 券核销流水号，针对该次核销发起售后退操作
func (r *AlibabaAlscMerchantExtTicketRefundAPIRequest) SetTransId(_transId string) error {
	r._transId = _transId
	r.Set("trans_id", _transId)
	return nil
}

// GetTransId TransId Getter
func (r AlibabaAlscMerchantExtTicketRefundAPIRequest) GetTransId() string {
	return r._transId
}

// SetTicketRequestId is TicketRequestId Setter
// 外部请求号，支持英文字母和数字，由开发者自行定义（不允许重复）
func (r *AlibabaAlscMerchantExtTicketRefundAPIRequest) SetTicketRequestId(_ticketRequestId string) error {
	r._ticketRequestId = _ticketRequestId
	r.Set("ticket_request_id", _ticketRequestId)
	return nil
}

// GetTicketRequestId TicketRequestId Getter
func (r AlibabaAlscMerchantExtTicketRefundAPIRequest) GetTicketRequestId() string {
	return r._ticketRequestId
}

// SetTicketCode is TicketCode Setter
// 凭证码，包括内部凭证码和外部凭证码，内部凭证码为12位，纯数字，且唯一不重复
func (r *AlibabaAlscMerchantExtTicketRefundAPIRequest) SetTicketCode(_ticketCode string) error {
	r._ticketCode = _ticketCode
	r.Set("ticket_code", _ticketCode)
	return nil
}

// GetTicketCode TicketCode Getter
func (r AlibabaAlscMerchantExtTicketRefundAPIRequest) GetTicketCode() string {
	return r._ticketCode
}
