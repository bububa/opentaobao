package damai

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenChangeticketAPIRequest 大麦换验平台-第三方对外开放-票单接口changeTicket API请求
// alibaba.damai.mev.open.changeticket
//
// 开放接口 换票
type AlibabaDamaiMevOpenChangeticketAPIRequest struct {
	model.Params
	// 入参ticketIdOpenParam
	_ticketIdOpenParam *TicketIdOpenParam
}

// NewAlibabaDamaiMevOpenChangeticketRequest 初始化AlibabaDamaiMevOpenChangeticketAPIRequest对象
func NewAlibabaDamaiMevOpenChangeticketRequest() *AlibabaDamaiMevOpenChangeticketAPIRequest {
	return &AlibabaDamaiMevOpenChangeticketAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMevOpenChangeticketAPIRequest) Reset() {
	r._ticketIdOpenParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenChangeticketAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.changeticket"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenChangeticketAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMevOpenChangeticketAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTicketIdOpenParam is TicketIdOpenParam Setter
// 入参ticketIdOpenParam
func (r *AlibabaDamaiMevOpenChangeticketAPIRequest) SetTicketIdOpenParam(_ticketIdOpenParam *TicketIdOpenParam) error {
	r._ticketIdOpenParam = _ticketIdOpenParam
	r.Set("ticket_id_open_param", _ticketIdOpenParam)
	return nil
}

// GetTicketIdOpenParam TicketIdOpenParam Getter
func (r AlibabaDamaiMevOpenChangeticketAPIRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
	return r._ticketIdOpenParam
}

var poolAlibabaDamaiMevOpenChangeticketAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMevOpenChangeticketRequest()
	},
}

// GetAlibabaDamaiMevOpenChangeticketRequest 从 sync.Pool 获取 AlibabaDamaiMevOpenChangeticketAPIRequest
func GetAlibabaDamaiMevOpenChangeticketAPIRequest() *AlibabaDamaiMevOpenChangeticketAPIRequest {
	return poolAlibabaDamaiMevOpenChangeticketAPIRequest.Get().(*AlibabaDamaiMevOpenChangeticketAPIRequest)
}

// ReleaseAlibabaDamaiMevOpenChangeticketAPIRequest 将 AlibabaDamaiMevOpenChangeticketAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMevOpenChangeticketAPIRequest(v *AlibabaDamaiMevOpenChangeticketAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMevOpenChangeticketAPIRequest.Put(v)
}
