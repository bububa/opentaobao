package bus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusTicketSetAPIRequest 出票接口 API请求
// taobao.bus.ticket.set
//
// 提供给汽车票商家出票使用
type TaobaoBusTicketSetAPIRequest struct {
	model.Params
	// 系统自动生成
	_ticketParams *B2BBookOrderRq
}

// NewTaobaoBusTicketSetRequest 初始化TaobaoBusTicketSetAPIRequest对象
func NewTaobaoBusTicketSetRequest() *TaobaoBusTicketSetAPIRequest {
	return &TaobaoBusTicketSetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBusTicketSetAPIRequest) Reset() {
	r._ticketParams = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusTicketSetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.ticket.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusTicketSetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusTicketSetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTicketParams is TicketParams Setter
// 系统自动生成
func (r *TaobaoBusTicketSetAPIRequest) SetTicketParams(_ticketParams *B2BBookOrderRq) error {
	r._ticketParams = _ticketParams
	r.Set("ticket_params", _ticketParams)
	return nil
}

// GetTicketParams TicketParams Getter
func (r TaobaoBusTicketSetAPIRequest) GetTicketParams() *B2BBookOrderRq {
	return r._ticketParams
}

var poolTaobaoBusTicketSetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBusTicketSetRequest()
	},
}

// GetTaobaoBusTicketSetRequest 从 sync.Pool 获取 TaobaoBusTicketSetAPIRequest
func GetTaobaoBusTicketSetAPIRequest() *TaobaoBusTicketSetAPIRequest {
	return poolTaobaoBusTicketSetAPIRequest.Get().(*TaobaoBusTicketSetAPIRequest)
}

// ReleaseTaobaoBusTicketSetAPIRequest 将 TaobaoBusTicketSetAPIRequest 放入 sync.Pool
func ReleaseTaobaoBusTicketSetAPIRequest(v *TaobaoBusTicketSetAPIRequest) {
	v.Reset()
	poolTaobaoBusTicketSetAPIRequest.Put(v)
}
