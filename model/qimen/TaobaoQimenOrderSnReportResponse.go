package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
订单SN通知接口 APIResponse
taobao.qimen.order.sn.report

WMS调用奇门的接口,在出库、发货、入库等场景下，ERP和WMS之间同步操作的SN列表
*/
type TaobaoQimenOrderSnReportAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenOrderSnReportResponse `json:"qimen_order_sn_report_response,omitempty"` 
    TaobaoQimenOrderSnReportResponse
}

/* model for simplify = false
type TaobaoQimenOrderSnReportResponse struct {

    // 
    
    Response  *struct {
        Response  *Response `json:"response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenOrderSnReportResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
