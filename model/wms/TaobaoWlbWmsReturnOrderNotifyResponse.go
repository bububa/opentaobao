package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
销售退货通知 APIResponse
taobao.wlb.wms.return.order.notify

销售退货通知
*/
type TaobaoWlbWmsReturnOrderNotifyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_wms_return_order_notify_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 订单创建时间
    
    CreateTime   string `json:"create_time,omitempty" xml:"