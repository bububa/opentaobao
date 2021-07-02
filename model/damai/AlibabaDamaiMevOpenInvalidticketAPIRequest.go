package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenInvalidticketAPIRequest 大麦换验平台-第三方对外开放-票单接口invalidTicket API请求
// alibaba.damai.mev.open.invalidticket
//
// 开放接口 使票无效
type AlibabaDamaiMevOpenInvalidticketAPIRequest struct {
	model.Params
	// 入参ticketIdOpenParam
	_ticketIdOpenParam *TicketIdOpenParam
}

// NewAlibabaDamaiMevOpenInvalidticketRequest 初始化AlibabaDamaiMevOpenInvalidticketAPIRequest对象
func NewAlibabaDamaiMevOpenInvalidticketRequest() *AlibabaDamaiMevOpenInvalidticketAPIRequest {
	return &AlibabaDamaiMevOpenInvalidticketAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenInvalidticketAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.invalidticket"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenInvalidticketAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTicketIdOpenParam is TicketIdOpenParam Setter
// 入参ticketIdOpenParam
func (r *AlibabaDamaiMevOpenInvalidticketAPIRequest) SetTicketIdOpenParam(_ticketIdOpenParam *TicketIdOpenParam) error {
	r._ticketIdOpenParam = _ticketIdOpenParam
	r.Set("ticket_id_open_param", _ticketIdOpenParam)
	return nil
}

// GetTicketIdOpenParam TicketIdOpenParam Getter
func (r AlibabaDamaiMevOpenInvalidticketAPIRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
	return r._ticketIdOpenParam
}
