package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusTicketSetAPIRequest
出票接口 API请求
taobao.bus.ticket.set

提供给汽车票商家出票使用 */
type TaobaoBusTicketSetAPIRequest struct {
	model.Params
	// 系统自动生成
	_ticketParams *B2BBookOrderRq
}

// NewTaobaoBusTicketSetRequest 初始化TaobaoBusTicketSetAPIRequest对象
func NewTaobaoBusTicketSetRequest() *TaobaoBusTicketSetAPIRequest {
	return &TaobaoBusTicketSetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusTicketSetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.ticket.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusTicketSetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TicketParams Setter
// 系统自动生成
func (r *TaobaoBusTicketSetAPIRequest) SetTicketParams(_ticketParams *B2BBookOrderRq) error {
	r._ticketParams = _ticketParams
	r.Set("ticket_params", _ticketParams)
	return nil
}

// Get TicketParams Getter
func (r TaobaoBusTicketSetAPIRequest) GetTicketParams() *B2BBookOrderRq {
	return r._ticketParams
}
