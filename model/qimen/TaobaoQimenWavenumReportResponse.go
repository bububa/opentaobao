package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发货单波次通知接口 APIResponse
taobao.qimen.wavenum.report

WMS调用奇门的接口,该接口用来给ERP或者OMS回传波次号及对应的发货单号，以支持商家货票同行、波次内包裹的处理等需求
*/
type TaobaoQimenWavenumReportAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenWavenumReportResponse `json:"qimen_wavenum_report_response,omitempty"` 
    TaobaoQimenWavenumReportResponse
}

/* model for simplify = false
type TaobaoQimenWavenumReportResponse struct {

    // 
    
    Response  *struct {
        Response  *Response `json:"response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenWavenumReportResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
