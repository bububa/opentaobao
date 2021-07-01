package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenLockticketAPIRequest
大麦换验平台-第三方对外开放-票单接口lockTicket API请求
alibaba.damai.mev.open.lockticket

开放接口 冻结票单 */
type AlibabaDamaiMevOpenLockticketAPIRequest struct {
	model.Params
	// 入参ticketIdOpenParam
	_ticketIdOpenParam *TicketIdOpenParam
}

// New
