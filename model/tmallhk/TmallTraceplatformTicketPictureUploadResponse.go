package tmallhk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
上传小票图片 APIResponse
tmall.traceplatform.ticket.picture.upload

uploadPicture
*/
type TmallTraceplatformTicketPictureUploadAPIResponse struct {
    model.CommonResponse
    Response *TmallTraceplatformTicketPictureUploadResponse `json:"tmall_traceplatform_ticket_picture_upload_response,omitempty"`
}

type TmallTraceplatformTicketPictureUploadResponse struct {

    // 返回值
    Result   *DataResult `json:"result,omitempty"`

}
