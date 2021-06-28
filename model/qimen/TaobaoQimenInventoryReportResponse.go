package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
库存盘点通知接口 APIResponse
taobao.qimen.inventory.report

WMS调用奇门的接口,将库存盘点情况回传ERP
*/
type TaobaoQimenInventoryReportAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenInventoryReportResponse `json:"qimen_inventory_report_response,omitempty"` 
    TaobaoQimenInventoryReportResponse
}

/* model for simplify = false
type TaobaoQimenInventoryReportResponse struct {

    // 
    
    Response  *struct {
        Response  *Response `json:"response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenInventoryReportResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
