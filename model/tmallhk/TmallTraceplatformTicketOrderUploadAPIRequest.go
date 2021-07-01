package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallTraceplatformTicketOrderUploadAPIRequest
上传小票数据 API请求
tmall.traceplatform.ticket.order.upload

upsertOrderBySeller */
type TmallTraceplatformTicketOrderUploadAPIRequest struct {
	model.Params
	// 上传小票参数
	_ticketOrder *TicketOrderUpdator
}

// New
