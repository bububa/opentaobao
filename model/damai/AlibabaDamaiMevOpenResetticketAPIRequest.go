package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopenresetticketAPIRequest 大麦换验平台-第三方对外开放-票单接口resetTicket API请求
// alibaba.damai.mev.open.resetticket
//
// 开放接口重打票
type AlibabadamaimevopenresetticketAPIRequest struct {
	model.Params
	// 入参ticketIdOpenParam
	_ticketIdOpenParam *TicketIdOpenParam
}

// NewAlibabadamaimevopenresetticketRequest 初始化AlibabadamaimevopenresetticketAPIRequest对象
func NewAlibabadamaimevopenresetticketRequest() *AlibabadamaimevopenresetticketAPIRequest {
	return &AlibabadamaimevopenresetticketAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimevopenresetticketAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.resetticket"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimevopenresetticketAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimevopenresetticketAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTicketIdOpenParam is TicketIdOpenParam Setter
// 入参ticketIdOpenParam
func (r *AlibabadamaimevopenresetticketAPIRequest) SetTicketIdOpenParam(_ticketIdOpenParam *TicketIdOpenParam) error {
	r._ticketIdOpenParam = _ticketIdOpenParam
	r.Set("ticket_id_open_param", _ticketIdOpenParam)
	return nil
}

// GetTicketIdOpenParam TicketIdOpenParam Getter
func (r AlibabadamaimevopenresetticketAPIRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
	return r._ticketIdOpenParam
}
