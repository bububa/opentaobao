package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发货单SN通知接口 APIResponse
taobao.qimen.sn.report

WMS调用奇门的接口,在仓库出库单后, 把SN信息回传给ERP
*/
type TaobaoQimenSnReportAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenSnReportResponse `json:"qimen_sn_report_response,omitempty"` 
    TaobaoQimenSnReportResponse
}

/* model for simplify = false
type TaobaoQimenSnReportResponse struct {

    // 
    
    Response  *struct {
        Response  *Response `json:"response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenSnReportResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
