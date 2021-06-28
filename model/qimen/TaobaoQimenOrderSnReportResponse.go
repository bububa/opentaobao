package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单SN通知接口 APIResponse
taobao.qimen.order.sn.report

WMS调用奇门的接口,在出库、发货、入库等场景下，ERP和WMS之间同步操作的SN列表
*/
type TaobaoQimenOrderSnReportAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_order_sn_report_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *Response `json:"response,omitempty" xml:"