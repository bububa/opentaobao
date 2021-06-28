package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
库存异动通知接口 APIResponse
taobao.qimen.stockchange.report

WMS调用奇门的接口,将库存异动信息信息回传给ERP
*/
type TaobaoQimenStockchangeReportAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenStockchangeReportResponse `json:"qimen_stockchange_report_response,omitempty"` 
    TaobaoQimenStockchangeReportResponse
}

/* model for simplify = false
type TaobaoQimenStockchangeReportResponse struct {

    // 
    
    Response  *struct {
        Response  *Response `json:"response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenStockchangeReportResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
