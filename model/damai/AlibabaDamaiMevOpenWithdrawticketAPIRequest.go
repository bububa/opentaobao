package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenWithdrawticketAPIRequest
大麦换验平台-第三方对外开放-票单接口withdrawTicket API请求
alibaba.damai.mev.open.withdrawticket

开放接口退票 */
type AlibabaDamaiMevOpenWithdrawticketAPIRequest struct {
	model.Params
	// 入参ticketIdOpenParam
	_ticketIdOpenParam *TicketIdOpenParam
}

// NewAlibabaDamaiMevOpenWithdrawticketRequest 初始化AlibabaDamaiMevOpenWithdrawticketAPIRequest对象
func NewAlibabaDamaiMevOpenWithdrawticketRequest() *AlibabaDamaiMevOpenWithdrawticketAPIRequest {
	return &AlibabaDamaiMevOpenWithdrawticketAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenWithdrawticketAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.withdrawticket"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenWithdrawticketAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TicketIdOpenParam Setter
// 入参ticketIdOpenParam
func (r *AlibabaDamaiMevOpenWithdrawticketAPIRequest) SetTicketIdOpenParam(_ticketIdOpenParam *TicketIdOpenParam) error {
	r._ticketIdOpenParam = _ticketIdOpenParam
	r.Set("ticket_id_open_param", _ticketIdOpenParam)
	return nil
}

// Get TicketIdOpenParam Getter
func (r AlibabaDamaiMevOpenWithdrawticketAPIRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
	return r._ticketIdOpenParam
}
