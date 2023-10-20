package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopenunlockticketAPIRequest 大麦换验平台-第三方对外开放-票单接口unlockTicket API请求
// alibaba.damai.mev.open.unlockticket
//
// 开放接口 解锁票单
type AlibabadamaimevopenunlockticketAPIRequest struct {
	model.Params
	// 入参ticketIdOpenParam
	_ticketIdOpenParam *TicketIdOpenParam
}

// NewAlibabadamaimevopenunlockticketRequest 初始化AlibabadamaimevopenunlockticketAPIRequest对象
func NewAlibabadamaimevopenunlockticketRequest() *AlibabadamaimevopenunlockticketAPIRequest {
	return &AlibabadamaimevopenunlockticketAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimevopenunlockticketAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.unlockticket"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimevopenunlockticketAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimevopenunlockticketAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTicketIdOpenParam is TicketIdOpenParam Setter
// 入参ticketIdOpenParam
func (r *AlibabadamaimevopenunlockticketAPIRequest) SetTicketIdOpenParam(_ticketIdOpenParam *TicketIdOpenParam) error {
	r._ticketIdOpenParam = _ticketIdOpenParam
	r.Set("ticket_id_open_param", _ticketIdOpenParam)
	return nil
}

// GetTicketIdOpenParam TicketIdOpenParam Getter
func (r AlibabadamaimevopenunlockticketAPIRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
	return r._ticketIdOpenParam
}
