package tmallhk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmalltraceplatformticketpictureuploadAPIResponse 上传小票图片 API返回值
// tmall.traceplatform.ticket.picture.upload
//
// uploadPicture
type TmalltraceplatformticketpictureuploadAPIResponse struct {
	model.CommonResponse
	TmalltraceplatformticketpictureuploadAPIResponseModel
}

// TmalltraceplatformticketpictureuploadAPIResponseModel is 上传小票图片 成功返回结果
type TmalltraceplatformticketpictureuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_traceplatform_ticket_picture_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}
