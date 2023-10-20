package damai

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenLockticketAPIRequest 大麦换验平台-第三方对外开放-票单接口lockTicket API请求
// alibaba.damai.mev.open.lockticket
//
// 开放接口 冻结票单
type AlibabaDamaiMevOpenLockticketAPIRequest struct {
	model.Params
	// 入参ticketIdOpenParam
	_ticketIdOpenParam *TicketIdOpenParam
}

// NewAlibabaDamaiMevOpenLockticketRequest 初始化AlibabaDamaiMevOpenLockticketAPIRequest对象
func NewAlibabaDamaiMevOpenLockticketRequest() *AlibabaDamaiMevOpenLockticketAPIRequest {
	return &AlibabaDamaiMevOpenLockticketAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMevOpenLockticketAPIRequest) Reset() {
	r._ticketIdOpenParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenLockticketAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.lockticket"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenLockticketAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMevOpenLockticketAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTicketIdOpenParam is TicketIdOpenParam Setter
// 入参ticketIdOpenParam
func (r *AlibabaDamaiMevOpenLockticketAPIRequest) SetTicketIdOpenParam(_ticketIdOpenParam *TicketIdOpenParam) error {
	r._ticketIdOpenParam = _ticketIdOpenParam
	r.Set("ticket_id_open_param", _ticketIdOpenParam)
	return nil
}

// GetTicketIdOpenParam TicketIdOpenParam Getter
func (r AlibabaDamaiMevOpenLockticketAPIRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
	return r._ticketIdOpenParam
}

var poolAlibabaDamaiMevOpenLockticketAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMevOpenLockticketRequest()
	},
}

// GetAlibabaDamaiMevOpenLockticketRequest 从 sync.Pool 获取 AlibabaDamaiMevOpenLockticketAPIRequest
func GetAlibabaDamaiMevOpenLockticketAPIRequest() *AlibabaDamaiMevOpenLockticketAPIRequest {
	return poolAlibabaDamaiMevOpenLockticketAPIRequest.Get().(*AlibabaDamaiMevOpenLockticketAPIRequest)
}

// ReleaseAlibabaDamaiMevOpenLockticketAPIRequest 将 AlibabaDamaiMevOpenLockticketAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMevOpenLockticketAPIRequest(v *AlibabaDamaiMevOpenLockticketAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMevOpenLockticketAPIRequest.Put(v)
}
