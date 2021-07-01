package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenWithdrawticketAPIRequest
大麦换验平台-第三方对外开放-票单接口withdrawTicket API请求
alibaba.damai.mev.open.withdrawticket

开放接口退票 */
type AlibabaDamaiMevOpenWithdrawticketAPIRequest struct {
	model.Params
	// 入参ticketIdOpenParam
	_ticketIdOpenParam *TicketIdOpenParam
}

// New
