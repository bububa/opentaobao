package xhotel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取直连酒店（客栈）签约上线进度信息 APIResponse
taobao.xhotel.order.hotelsign.query

获取直连酒店（客栈）签约上线进度信息
*/
type TaobaoXhotelOrderHotelsignQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"xhotel_order_hotelsign_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // outUuid
    
    OutUuid   string `json:"out_uuid,omitempty" xml:"