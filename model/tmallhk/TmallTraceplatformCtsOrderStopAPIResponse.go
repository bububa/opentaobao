package tmallhk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTraceplatformCtsOrderStopAPIResponse CTS截断订单 API返回值
// tmall.traceplatform.cts.order.stop
//
// 截断CTS订单
type TmallTraceplatformCtsOrderStopAPIResponse struct {
	model.CommonResponse
	TmallTraceplatformCtsOrderStopAPIResponseModel
}

// Reset 清空结构体
func (m *TmallTraceplatformCtsOrderStopAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallTraceplatformCtsOrderStopAPIResponseModel).Reset()
}

// TmallTraceplatformCtsOrderStopAPIResponseModel is CTS截断订单 成功返回结果
type TmallTraceplatformCtsOrderStopAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_traceplatform_cts_order_stop_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallTraceplatformCtsOrderStopAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallTraceplatformCtsOrderStopAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallTraceplatformCtsOrderStopAPIResponse)
	},
}

// GetTmallTraceplatformCtsOrderStopAPIResponse 从 sync.Pool 获取 TmallTraceplatformCtsOrderStopAPIResponse
func GetTmallTraceplatformCtsOrderStopAPIResponse() *TmallTraceplatformCtsOrderStopAPIResponse {
	return poolTmallTraceplatformCtsOrderStopAPIResponse.Get().(*TmallTraceplatformCtsOrderStopAPIResponse)
}

// ReleaseTmallTraceplatformCtsOrderStopAPIResponse 将 TmallTraceplatformCtsOrderStopAPIResponse 保存到 sync.Pool
func ReleaseTmallTraceplatformCtsOrderStopAPIResponse(v *TmallTraceplatformCtsOrderStopAPIResponse) {
	v.Reset()
	poolTmallTraceplatformCtsOrderStopAPIResponse.Put(v)
}
