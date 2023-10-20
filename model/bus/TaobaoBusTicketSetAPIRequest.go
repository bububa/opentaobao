package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobusticketsetAPIRequest 出票接口 API请求
// taobao.bus.ticket.set
//
// 提供给汽车票商家出票使用
type TaobaobusticketsetAPIRequest struct {
	model.Params
	// 系统自动生成
	_ticketParams *B2bbookOrderRq
}

// NewTaobaobusticketsetRequest 初始化TaobaobusticketsetAPIRequest对象
func NewTaobaobusticketsetRequest() *TaobaobusticketsetAPIRequest {
	return &TaobaobusticketsetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobusticketsetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.ticket.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobusticketsetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobusticketsetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTicketParams is TicketParams Setter
// 系统自动生成
func (r *TaobaobusticketsetAPIRequest) SetTicketParams(_ticketParams *B2bbookOrderRq) error {
	r._ticketParams = _ticketParams
	r.Set("ticket_params", _ticketParams)
	return nil
}

// GetTicketParams TicketParams Getter
func (r TaobaobusticketsetAPIRequest) GetTicketParams() *B2bbookOrderRq {
	return r._ticketParams
}
