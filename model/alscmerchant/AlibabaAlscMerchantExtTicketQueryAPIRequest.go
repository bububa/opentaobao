package alscmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscMerchantExtTicketQueryAPIRequest 口碑凭证模版查询服务 API请求
// alibaba.alsc.merchant.ext.ticket.query
//
// isv需要在c端展示凭证信息,需要查询凭证模版
type AlibabaAlscMerchantExtTicketQueryAPIRequest struct {
	model.Params
	// 凭证模版ID
	_ticketTemplateId string
	// 请求requestID
	_ticketRequestId string
}

// NewAlibabaAlscMerchantExtTicketQueryRequest 初始化AlibabaAlscMerchantExtTicketQueryAPIRequest对象
func NewAlibabaAlscMerchantExtTicketQueryRequest() *AlibabaAlscMerchantExtTicketQueryAPIRequest {
	return &AlibabaAlscMerchantExtTicketQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscMerchantExtTicketQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.merchant.ext.ticket.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscMerchantExtTicketQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscMerchantExtTicketQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTicketTemplateId is TicketTemplateId Setter
// 凭证模版ID
func (r *AlibabaAlscMerchantExtTicketQueryAPIRequest) SetTicketTemplateId(_ticketTemplateId string) error {
	r._ticketTemplateId = _ticketTemplateId
	r.Set("ticket_template_id", _ticketTemplateId)
	return nil
}

// GetTicketTemplateId TicketTemplateId Getter
func (r AlibabaAlscMerchantExtTicketQueryAPIRequest) GetTicketTemplateId() string {
	return r._ticketTemplateId
}

// SetTicketRequestId is TicketRequestId Setter
// 请求requestID
func (r *AlibabaAlscMerchantExtTicketQueryAPIRequest) SetTicketRequestId(_ticketRequestId string) error {
	r._ticketRequestId = _ticketRequestId
	r.Set("ticket_request_id", _ticketRequestId)
	return nil
}

// GetTicketRequestId TicketRequestId Getter
func (r AlibabaAlscMerchantExtTicketQueryAPIRequest) GetTicketRequestId() string {
	return r._ticketRequestId
}
