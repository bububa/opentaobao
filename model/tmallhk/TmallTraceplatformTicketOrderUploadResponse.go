package tmallhk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
上传小票数据 APIResponse
tmall.traceplatform.ticket.order.upload

upsertOrderBySeller
*/
type TmallTraceplatformTicketOrderUploadAPIResponse struct {
    model.CommonResponse
    Response *TmallTraceplatformTicketOrderUploadResponse `json:"tmall_traceplatform_ticket_order_upload_response,omitempty"`
}

type TmallTraceplatformTicketOrderUploadResponse struct {

    // 返回值
    Result   *DataResult `json:"result,omitempty"`

}
