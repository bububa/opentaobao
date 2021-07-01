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

// New
