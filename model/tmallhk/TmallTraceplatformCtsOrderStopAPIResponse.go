package tmallhk

import (
	"encoding/xml"

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

// TmallTraceplatformCtsOrderStopAPIResponseModel is CTS截断订单 成功返回结果
type TmallTraceplatformCtsOrderStopAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_traceplatform_cts_order_stop_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}
