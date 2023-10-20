package tmallhk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmalltraceplatformticketorderuploadAPIResponse 上传小票数据 API返回值
// tmall.traceplatform.ticket.order.upload
//
// upsertOrderBySeller
type TmalltraceplatformticketorderuploadAPIResponse struct {
	model.CommonResponse
	TmalltraceplatformticketorderuploadAPIResponseModel
}

// TmalltraceplatformticketorderuploadAPIResponseModel is 上传小票数据 成功返回结果
type TmalltraceplatformticketorderuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_traceplatform_ticket_order_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}
