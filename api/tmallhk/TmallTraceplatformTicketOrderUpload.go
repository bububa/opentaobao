package tmallhk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// TmallTraceplatformTicketOrderUpload 上传小票数据
// tmall.traceplatform.ticket.order.upload
//
// upsertOrderBySeller
func TmallTraceplatformTicketOrderUpload(clt *core.SDKClient, req *tmallhk.TmallTraceplatformTicketOrderUploadAPIRequest, resp *tmallhk.TmallTraceplatformTicketOrderUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
