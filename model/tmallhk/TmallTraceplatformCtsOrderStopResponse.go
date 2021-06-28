package tmallhk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
CTS截断订单 APIResponse
tmall.traceplatform.cts.order.stop

截断CTS订单
*/
type TmallTraceplatformCtsOrderStopAPIResponse struct {
    model.CommonResponse
    // Response *TmallTraceplatformCtsOrderStopResponse `json:"tmall_traceplatform_cts_order_stop_response,omitempty"` 
    TmallTraceplatformCtsOrderStopResponse
}

/* model for simplify = false
type TmallTraceplatformCtsOrderStopResponse struct {

    // result
    
    Result  *struct {
        DataResult  *DataResult `json:"data_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallTraceplatformCtsOrderStopResponse struct {

    // result
    Result   *DataResult `json:"result,omitempty"`

}
