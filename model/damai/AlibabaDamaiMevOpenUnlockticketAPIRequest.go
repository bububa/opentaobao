package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenUnlockticketAPIRequest
大麦换验平台-第三方对外开放-票单接口unlockTicket API请求
alibaba.damai.mev.open.unlockticket

开放接口 解锁票单 */
type AlibabaDamaiMevOpenUnlockticketAPIRequest struct {
	model.Params
	// 入参ticketIdOpenParam
	_ticketIdOpenParam *TicketIdOpenParam
}

// New
