package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenUnlockticketAPIRequest 大麦换验平台-第三方对外开放-票单接口unlockTicket API请求
// alibaba.damai.mev.open.unlockticket
//
// 开放接口 解锁票单
type AlibabaDamaiMevOpenUnlockticketAPIRequest struct {
	model.Params
	// 入参ticketIdOpenParam
	_ticketIdOpenParam *TicketIdOpenParam
}

// NewAlibabaDamaiMevOpenUnlockticketRequest 初始化AlibabaDamaiMevOpenUnlockticketAPIRequest对象
func NewAlibabaDamaiMevOpenUnlockticketRequest() *AlibabaDamaiMevOpenUnlockticketAPIRequest {
	return &AlibabaDamaiMevOpenUnlockticketAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenUnlockticketAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.unlockticket"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenUnlockticketAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMevOpenUnlockticketAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTicketIdOpenParam is TicketIdOpenParam Setter
// 入参ticketIdOpenParam
func (r *AlibabaDamaiMevOpenUnlockticketAPIRequest) SetTicketIdOpenParam(_ticketIdOpenParam *TicketIdOpenParam) error {
	r._ticketIdOpenParam = _ticketIdOpenParam
	r.Set("ticket_id_open_param", _ticketIdOpenParam)
	return nil
}

// GetTicketIdOpenParam TicketIdOpenParam Getter
func (r AlibabaDamaiMevOpenUnlockticketAPIRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
	return r._ticketIdOpenParam
}
