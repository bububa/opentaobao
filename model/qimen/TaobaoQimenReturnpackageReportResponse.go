package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
退货包裹状态通知接口 APIResponse
taobao.qimen.returnpackage.report

退货包裹状态通知接口
*/
type TaobaoQimenReturnpackageReportAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenReturnpackageReportResponse `json:"qimen_returnpackage_report_response,omitempty"` 
    TaobaoQimenReturnpackageReportResponse
}

/* model for simplify = false
type TaobaoQimenReturnpackageReportResponse struct {

    // 
    
    Response  *struct {
        Response  *Response `json:"response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenReturnpackageReportResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
