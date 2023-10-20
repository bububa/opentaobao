package damai

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenWithdrawticketAPIRequest 大麦换验平台-第三方对外开放-票单接口withdrawTicket API请求
// alibaba.damai.mev.open.withdrawticket
//
// 开放接口退票
type AlibabaDamaiMevOpenWithdrawticketAPIRequest struct {
	model.Params
	// 入参ticketIdOpenParam
	_ticketIdOpenParam *TicketIdOpenParam
}

// NewAlibabaDamaiMevOpenWithdrawticketRequest 初始化AlibabaDamaiMevOpenWithdrawticketAPIRequest对象
func NewAlibabaDamaiMevOpenWithdrawticketRequest() *AlibabaDamaiMevOpenWithdrawticketAPIRequest {
	return &AlibabaDamaiMevOpenWithdrawticketAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMevOpenWithdrawticketAPIRequest) Reset() {
	r._ticketIdOpenParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenWithdrawticketAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.withdrawticket"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenWithdrawticketAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMevOpenWithdrawticketAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTicketIdOpenParam is TicketIdOpenParam Setter
// 入参ticketIdOpenParam
func (r *AlibabaDamaiMevOpenWithdrawticketAPIRequest) SetTicketIdOpenParam(_ticketIdOpenParam *TicketIdOpenParam) error {
	r._ticketIdOpenParam = _ticketIdOpenParam
	r.Set("ticket_id_open_param", _ticketIdOpenParam)
	return nil
}

// GetTicketIdOpenParam TicketIdOpenParam Getter
func (r AlibabaDamaiMevOpenWithdrawticketAPIRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
	return r._ticketIdOpenParam
}

var poolAlibabaDamaiMevOpenWithdrawticketAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMevOpenWithdrawticketRequest()
	},
}

// GetAlibabaDamaiMevOpenWithdrawticketRequest 从 sync.Pool 获取 AlibabaDamaiMevOpenWithdrawticketAPIRequest
func GetAlibabaDamaiMevOpenWithdrawticketAPIRequest() *AlibabaDamaiMevOpenWithdrawticketAPIRequest {
	return poolAlibabaDamaiMevOpenWithdrawticketAPIRequest.Get().(*AlibabaDamaiMevOpenWithdrawticketAPIRequest)
}

// ReleaseAlibabaDamaiMevOpenWithdrawticketAPIRequest 将 AlibabaDamaiMevOpenWithdrawticketAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMevOpenWithdrawticketAPIRequest(v *AlibabaDamaiMevOpenWithdrawticketAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMevOpenWithdrawticketAPIRequest.Put(v)
}
