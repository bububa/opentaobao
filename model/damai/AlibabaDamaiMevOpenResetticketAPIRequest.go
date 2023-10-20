package damai

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenResetticketAPIRequest 大麦换验平台-第三方对外开放-票单接口resetTicket API请求
// alibaba.damai.mev.open.resetticket
//
// 开放接口重打票
type AlibabaDamaiMevOpenResetticketAPIRequest struct {
	model.Params
	// 入参ticketIdOpenParam
	_ticketIdOpenParam *TicketIdOpenParam
}

// NewAlibabaDamaiMevOpenResetticketRequest 初始化AlibabaDamaiMevOpenResetticketAPIRequest对象
func NewAlibabaDamaiMevOpenResetticketRequest() *AlibabaDamaiMevOpenResetticketAPIRequest {
	return &AlibabaDamaiMevOpenResetticketAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMevOpenResetticketAPIRequest) Reset() {
	r._ticketIdOpenParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenResetticketAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.resetticket"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenResetticketAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMevOpenResetticketAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTicketIdOpenParam is TicketIdOpenParam Setter
// 入参ticketIdOpenParam
func (r *AlibabaDamaiMevOpenResetticketAPIRequest) SetTicketIdOpenParam(_ticketIdOpenParam *TicketIdOpenParam) error {
	r._ticketIdOpenParam = _ticketIdOpenParam
	r.Set("ticket_id_open_param", _ticketIdOpenParam)
	return nil
}

// GetTicketIdOpenParam TicketIdOpenParam Getter
func (r AlibabaDamaiMevOpenResetticketAPIRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
	return r._ticketIdOpenParam
}

var poolAlibabaDamaiMevOpenResetticketAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMevOpenResetticketRequest()
	},
}

// GetAlibabaDamaiMevOpenResetticketRequest 从 sync.Pool 获取 AlibabaDamaiMevOpenResetticketAPIRequest
func GetAlibabaDamaiMevOpenResetticketAPIRequest() *AlibabaDamaiMevOpenResetticketAPIRequest {
	return poolAlibabaDamaiMevOpenResetticketAPIRequest.Get().(*AlibabaDamaiMevOpenResetticketAPIRequest)
}

// ReleaseAlibabaDamaiMevOpenResetticketAPIRequest 将 AlibabaDamaiMevOpenResetticketAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMevOpenResetticketAPIRequest(v *AlibabaDamaiMevOpenResetticketAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMevOpenResetticketAPIRequest.Put(v)
}
