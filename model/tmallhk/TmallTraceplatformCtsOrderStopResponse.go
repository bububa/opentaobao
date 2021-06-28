package tmallhk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
CTS截断订单 APIResponse
tmall.traceplatform.cts.order.stop

截断CTS订单
*/
type TmallTraceplatformCtsOrderStopAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_traceplatform_cts_order_stop_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *DataResult `json:"result,omitempty" xml:"