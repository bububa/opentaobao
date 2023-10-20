package damai

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMevOpenInvalidticketAPIRequest) Reset() {
	r._ticketIdOpenParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenInvalidticketAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.invalidticket"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenInvalidticketAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMevOpenInvalidticketAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaDamaiMevOpenInvalidticketAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMevOpenInvalidticketRequest()
	},
}

// GetAlibabaDamaiMevOpenInvalidticketRequest 从 sync.Pool 获取 AlibabaDamaiMevOpenInvalidticketAPIRequest
func GetAlibabaDamaiMevOpenInvalidticketAPIRequest() *AlibabaDamaiMevOpenInvalidticketAPIRequest {
	return poolAlibabaDamaiMevOpenInvalidticketAPIRequest.Get().(*AlibabaDamaiMevOpenInvalidticketAPIRequest)
}

// ReleaseAlibabaDamaiMevOpenInvalidticketAPIRequest 将 AlibabaDamaiMevOpenInvalidticketAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMevOpenInvalidticketAPIRequest(v *AlibabaDamaiMevOpenInvalidticketAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMevOpenInvalidticketAPIRequest.Put(v)
}
