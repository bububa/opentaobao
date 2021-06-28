package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
订单异常通知接口 APIResponse
taobao.qimen.orderexception.report

WMS调用奇门的接口,当WMS接收到ERP的发货指令时，由于种种原因（5.1.5说明了各种异常场景）可能无法完成发货。WMS通过调用此接口，通知ERP具体异常情况
*/
type TaobaoQimenOrderexceptionReportAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenOrderexceptionReportResponse `json:"qimen_orderexception_report_response,omitempty"` 
    TaobaoQimenOrderexceptionReportResponse
}

/* model for simplify = false
type TaobaoQimenOrderexceptionReportResponse struct {

    // 
    
    Response  *struct {
        Response  *Response `json:"response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenOrderexceptionReportResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
