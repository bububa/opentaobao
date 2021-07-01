package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenChangeticketAPIRequest
大麦换验平台-第三方对外开放-票单接口changeTicket API请求
alibaba.damai.mev.open.changeticket

开放接口 换票 */
type AlibabaDamaiMevOpenChangeticketAPIRequest struct {
	model.Params
	// 入参ticketIdOpenParam
	_ticketIdOpenParam *TicketIdOpenParam
}

// New
