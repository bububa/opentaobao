package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
库存盘点通知接口 APIResponse
taobao.qimen.inventory.report

WMS调用奇门的接口,将库存盘点情况回传ERP
*/
type TaobaoQimenInventoryReportAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_inventory_report_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *Response `json:"response,omitempty" xml:"