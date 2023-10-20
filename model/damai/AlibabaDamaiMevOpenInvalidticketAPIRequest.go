package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopeninvalidticketAPIRequest 大麦换验平台-第三方对外开放-票单接口invalidTicket API请求
// alibaba.damai.mev.open.invalidticket
//
// 开放接口 使票无效
type AlibabadamaimevopeninvalidticketAPIRequest struct {
	model.Params
	// 入参ticketIdOpenParam
	_ticketIdOpenParam *TicketIdOpenParam
}

// NewAlibabadamaimevopeninvalidticketRequest 初始化AlibabadamaimevopeninvalidticketAPIRequest对象
func NewAlibabadamaimevopeninvalidticketRequest() *AlibabadamaimevopeninvalidticketAPIRequest {
	return &AlibabadamaimevopeninvalidticketAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimevopeninvalidticketAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.invalidticket"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimevopeninvalidticketAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimevopeninvalidticketAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTicketIdOpenParam is TicketIdOpenParam Setter
// 入参ticketIdOpenParam
func (r *AlibabadamaimevopeninvalidticketAPIRequest) SetTicketIdOpenParam(_ticketIdOpenParam *TicketIdOpenParam) error {
	r._ticketIdOpenParam = _ticketIdOpenParam
	r.Set("ticket_id_open_param", _ticketIdOpenParam)
	return nil
}

// GetTicketIdOpenParam TicketIdOpenParam Getter
func (r AlibabadamaimevopeninvalidticketAPIRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
	return r._ticketIdOpenParam
}
