package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjOcOnlineTicketnoGetAPIRequest
线上小票号获取 API请求
alibaba.mj.oc.online.ticketno.get

线上小票号获取 */
type AlibabaMjOcOnlineTicketnoGetAPIRequest struct {
	model.Params
	// 外部门店号
	_outStoreNo string
}

// New
