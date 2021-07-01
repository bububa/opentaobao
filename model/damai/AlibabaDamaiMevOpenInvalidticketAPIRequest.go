package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenInvalidticketAPIRequest
大麦换验平台-第三方对外开放-票单接口invalidTicket API请求
alibaba.damai.mev.open.invalidticket

开放接口 使票无效 */
type AlibabaDamaiMevOpenInvalidticketAPIRequest struct {
	model.Params
	// 入参ticketIdOpenParam
	_ticketIdOpenParam *TicketIdOpenParam
}

// New
