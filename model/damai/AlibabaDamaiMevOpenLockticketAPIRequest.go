package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopenlockticketAPIRequest 大麦换验平台-第三方对外开放-票单接口lockTicket API请求
// alibaba.damai.mev.open.lockticket
//
// 开放接口 冻结票单
type AlibabadamaimevopenlockticketAPIRequest struct {
	model.Params
	// 入参ticketIdOpenParam
	_ticketIdOpenParam *TicketIdOpenParam
}

// NewAlibabadamaimevopenlockticketRequest 初始化AlibabadamaimevopenlockticketAPIRequest对象
func NewAlibabadamaimevopenlockticketRequest() *AlibabadamaimevopenlockticketAPIRequest {
	return &AlibabadamaimevopenlockticketAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimevopenlockticketAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.lockticket"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimevopenlockticketAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimevopenlockticketAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTicketIdOpenParam is TicketIdOpenParam Setter
// 入参ticketIdOpenParam
func (r *AlibabadamaimevopenlockticketAPIRequest) SetTicketIdOpenParam(_ticketIdOpenParam *TicketIdOpenParam) error {
	r._ticketIdOpenParam = _ticketIdOpenParam
	r.Set("ticket_id_open_param", _ticketIdOpenParam)
	return nil
}

// GetTicketIdOpenParam TicketIdOpenParam Getter
func (r AlibabadamaimevopenlockticketAPIRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
	return r._ticketIdOpenParam
}
