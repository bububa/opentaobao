package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
调拨单通知 APIResponse
taobao.qimen.transferorder.report

调拨单通知
*/
type TaobaoQimenTransferorderReportAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenTransferorderReportResponse `json:"qimen_transferorder_report_response,omitempty"` 
    TaobaoQimenTransferorderReportResponse
}

/* model for simplify = false
type TaobaoQimenTransferorderReportResponse struct {

    // 
    
    Response  *struct {
        TaobaoQimenTransferorderReportStruct  *TaobaoQimenTransferorderReportStruct `json:"taobao_qimen_transferorder_report_struct,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenTransferorderReportResponse struct {

    // 
    Response   *TaobaoQimenTransferorderReportStruct `json:"response,omitempty"`

}
