package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenChangeticketAPIRequest 大麦换验平台-第三方对外开放-票单接口changeTicket API请求
// alibaba.damai.mev.open.changeticket
//
// 开放接口 换票
type AlibabaDamaiMevOpenChangeticketAPIRequest struct {
	model.Params
	// 入参ticketIdOpenParam
	_ticketIdOpenParam *TicketIdOpenParam
}

// NewAlibabaDamaiMevOpenChangeticketRequest 初始化AlibabaDamaiMevOpenChangeticketAPIRequest对象
func NewAlibabaDamaiMevOpenChangeticketRequest() *AlibabaDamaiMevOpenChangeticketAPIRequest {
	return &AlibabaDamaiMevOpenChangeticketAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenChangeticketAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.changeticket"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenChangeticketAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTicketIdOpenParam is TicketIdOpenParam Setter
// 入参ticketIdOpenParam
func (r *AlibabaDamaiMevOpenChangeticketAPIRequest) SetTicketIdOpenParam(_ticketIdOpenParam *TicketIdOpenParam) error {
	r._ticketIdOpenParam = _ticketIdOpenParam
	r.Set("ticket_id_open_param", _ticketIdOpenParam)
	return nil
}

// GetTicketIdOpenParam TicketIdOpenParam Getter
func (r AlibabaDamaiMevOpenChangeticketAPIRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
	return r._ticketIdOpenParam
}
