package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenResetticketAPIRequest
大麦换验平台-第三方对外开放-票单接口resetTicket API请求
alibaba.damai.mev.open.resetticket

开放接口重打票 */
type AlibabaDamaiMevOpenResetticketAPIRequest struct {
	model.Params
	// 入参ticketIdOpenParam
	_ticketIdOpenParam *TicketIdOpenParam
}

// New
