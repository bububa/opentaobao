package tmallhk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTraceplatformTicketOrderUploadAPIResponse 上传小票数据 API返回值
// tmall.traceplatform.ticket.order.upload
//
// upsertOrderBySeller
type TmallTraceplatformTicketOrderUploadAPIResponse struct {
	model.CommonResponse
	TmallTraceplatformTicketOrderUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TmallTraceplatformTicketOrderUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallTraceplatformTicketOrderUploadAPIResponseModel).Reset()
}

// TmallTraceplatformTicketOrderUploadAPIResponseModel is 上传小票数据 成功返回结果
type TmallTraceplatformTicketOrderUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_traceplatform_ticket_order_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallTraceplatformTicketOrderUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallTraceplatformTicketOrderUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallTraceplatformTicketOrderUploadAPIResponse)
	},
}

// GetTmallTraceplatformTicketOrderUploadAPIResponse 从 sync.Pool 获取 TmallTraceplatformTicketOrderUploadAPIResponse
func GetTmallTraceplatformTicketOrderUploadAPIResponse() *TmallTraceplatformTicketOrderUploadAPIResponse {
	return poolTmallTraceplatformTicketOrderUploadAPIResponse.Get().(*TmallTraceplatformTicketOrderUploadAPIResponse)
}

// ReleaseTmallTraceplatformTicketOrderUploadAPIResponse 将 TmallTraceplatformTicketOrderUploadAPIResponse 保存到 sync.Pool
func ReleaseTmallTraceplatformTicketOrderUploadAPIResponse(v *TmallTraceplatformTicketOrderUploadAPIResponse) {
	v.Reset()
	poolTmallTraceplatformTicketOrderUploadAPIResponse.Put(v)
}
