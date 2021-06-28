package tmallhk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
上传小票图片 APIResponse
tmall.traceplatform.ticket.picture.upload

uploadPicture
*/
type TmallTraceplatformTicketPictureUploadAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_traceplatform_ticket_picture_upload_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回值
    
    Result   *DataResult `json:"result,omitempty" xml:"