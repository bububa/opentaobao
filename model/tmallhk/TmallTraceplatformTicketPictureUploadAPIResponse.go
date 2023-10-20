package tmallhk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTraceplatformTicketPictureUploadAPIResponse 上传小票图片 API返回值
// tmall.traceplatform.ticket.picture.upload
//
// uploadPicture
type TmallTraceplatformTicketPictureUploadAPIResponse struct {
	model.CommonResponse
	TmallTraceplatformTicketPictureUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TmallTraceplatformTicketPictureUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallTraceplatformTicketPictureUploadAPIResponseModel).Reset()
}

// TmallTraceplatformTicketPictureUploadAPIResponseModel is 上传小票图片 成功返回结果
type TmallTraceplatformTicketPictureUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_traceplatform_ticket_picture_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallTraceplatformTicketPictureUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallTraceplatformTicketPictureUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallTraceplatformTicketPictureUploadAPIResponse)
	},
}

// GetTmallTraceplatformTicketPictureUploadAPIResponse 从 sync.Pool 获取 TmallTraceplatformTicketPictureUploadAPIResponse
func GetTmallTraceplatformTicketPictureUploadAPIResponse() *TmallTraceplatformTicketPictureUploadAPIResponse {
	return poolTmallTraceplatformTicketPictureUploadAPIResponse.Get().(*TmallTraceplatformTicketPictureUploadAPIResponse)
}

// ReleaseTmallTraceplatformTicketPictureUploadAPIResponse 将 TmallTraceplatformTicketPictureUploadAPIResponse 保存到 sync.Pool
func ReleaseTmallTraceplatformTicketPictureUploadAPIResponse(v *TmallTraceplatformTicketPictureUploadAPIResponse) {
	v.Reset()
	poolTmallTraceplatformTicketPictureUploadAPIResponse.Put(v)
}
